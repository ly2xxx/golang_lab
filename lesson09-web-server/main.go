// Lesson 09: Web Server Basics with net/http
// This lesson covers building HTTP servers, handling requests, and middleware

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// User struct for demonstration
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Simple in-memory "database"
var users = map[int]User{
	1: {ID: 1, Name: "Alice", Email: "alice@example.com"},
	2: {ID: 2, Name: "Bob", Email: "bob@example.com"},
	3: {ID: 3, Name: "Charlie", Email: "charlie@example.com"},
}
var nextUserID = 4

func main() {
	fmt.Println("=== Lesson 09: Web Server Basics ===")
	
	// Create a new ServeMux (router)
	mux := http.NewServeMux()
	
	// Register routes
	registerRoutes(mux)
	
	// Apply middleware
	handler := loggingMiddleware(corsMiddleware(mux))
	
	// Create server with configuration
	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	fmt.Println("Starting server on http://localhost:8080")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  /              - Home page")
	fmt.Println("  GET  /hello         - Simple greeting")
	fmt.Println("  GET  /hello/{name}  - Personalized greeting")
	fmt.Println("  GET  /users         - List all users")
	fmt.Println("  GET  /users/{id}    - Get specific user")
	fmt.Println("  POST /users         - Create new user (form data)")
	fmt.Println("  GET  /form          - User creation form")
	fmt.Println("  GET  /static/*      - Static files")
	fmt.Println("\nPress Ctrl+C to stop the server")
	
	// Start server
	log.Fatal(server.ListenAndServe())
}

func registerRoutes(mux *http.ServeMux) {
	// Static file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	
	// Basic routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/hello/", helloNameHandler)
	
	// User routes
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/", userHandler)
	
	// Form routes
	mux.HandleFunc("/form", formHandler)
	
	// Health check
	mux.HandleFunc("/health", healthHandler)
}

// Home page handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Go Web Server Tutorial</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .endpoint { background: #f4f4f4; padding: 10px; margin: 10px 0; border-radius: 5px; }
        a { color: #007bff; text-decoration: none; }
        a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <h1>Welcome to Go Web Server Tutorial!</h1>
    <p>This is a demonstration of various HTTP server features in Go.</p>
    
    <h2>Available Endpoints:</h2>
    <div class="endpoint">
        <strong>GET <a href="/hello">/hello</a></strong> - Simple greeting
    </div>
    <div class="endpoint">
        <strong>GET <a href="/hello/World">/hello/World</a></strong> - Personalized greeting
    </div>
    <div class="endpoint">
        <strong>GET <a href="/users">/users</a></strong> - List all users (JSON)
    </div>
    <div class="endpoint">
        <strong>GET <a href="/users/1">/users/1</a></strong> - Get specific user (JSON)
    </div>
    <div class="endpoint">
        <strong>GET <a href="/form">/form</a></strong> - User creation form
    </div>
    <div class="endpoint">
        <strong>GET <a href="/health">/health</a></strong> - Health check
    </div>
    
    <h2>Request Information:</h2>
    <p><strong>Method:</strong> %s</p>
    <p><strong>URL:</strong> %s</p>
    <p><strong>User Agent:</strong> %s</p>
    <p><strong>Remote Address:</strong> %s</p>
    <p><strong>Timestamp:</strong> %s</p>
</body>
</html>
`
	
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, html, r.Method, r.URL.String(), r.UserAgent(), r.RemoteAddr, time.Now().Format(time.RFC3339))
}

// Simple hello handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, %s! This is a Go web server.\n", name)
}

// Hello with name from URL path
func helloNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Extract name from URL path
	path := strings.TrimPrefix(r.URL.Path, "/hello/")
	if path == "" {
		http.Redirect(w, r, "/hello", http.StatusFound)
		return
	}
	
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, %s! Nice to meet you.\n", path)
}

// Users handler (handles both GET /users and POST /users)
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAllUsers(w, r)
	case http.MethodPost:
		createUser(w, r)
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Get all users
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Simple JSON response (in a real app, use json.Marshal)
	fmt.Fprint(w, "[")
	first := true
	for _, user := range users {
		if !first {
			fmt.Fprint(w, ",")
		}
		fmt.Fprintf(w, `{"id":%d,"name":"%s","email":"%s"}`, user.ID, user.Name, user.Email)
		first = false
	}
	fmt.Fprint(w, "]")
}

// Create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	
	if name == "" || email == "" {
		http.Error(w, "Name and email are required", http.StatusBadRequest)
		return
	}
	
	// Create new user
	user := User{
		ID:    nextUserID,
		Name:  name,
		Email: email,
	}
	users[nextUserID] = user
	nextUserID++
	
	// Return created user as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"id":%d,"name":"%s","email":"%s"}`, user.ID, user.Name, user.Email)
}

// Individual user handler
func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Extract user ID from URL
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	if path == "" {
		http.Redirect(w, r, "/users", http.StatusFound)
		return
	}
	
	userID, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	user, exists := users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"id":%d,"name":"%s","email":"%s"}`, user.ID, user.Name, user.Email)
}

// Form handler for creating users
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>Create User</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .form-group { margin-bottom: 15px; }
        label { display: block; margin-bottom: 5px; font-weight: bold; }
        input[type="text"], input[type="email"] {
            width: 100%;
            padding: 8px;
            border: 1px solid #ddd;
            border-radius: 4px;
            box-sizing: border-box;
        }
        button {
            background-color: #007bff;
            color: white;
            padding: 10px 20px;
            border: none;
            border-radius: 4px;
            cursor: pointer;
        }
        button:hover { background-color: #0056b3; }
        .back-link { margin-top: 20px; }
    </style>
</head>
<body>
    <h1>Create New User</h1>
    <form action="/users" method="POST">
        <div class="form-group">
            <label for="name">Name:</label>
            <input type="text" id="name" name="name" required>
        </div>
        <div class="form-group">
            <label for="email">Email:</label>
            <input type="email" id="email" name="email" required>
        </div>
        <button type="submit">Create User</button>
    </form>
    <div class="back-link">
        <a href="/">← Back to Home</a>
    </div>
</body>
</html>
`
	
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, tmpl)
}

// Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"healthy","timestamp":"%s","users_count":%d}`, 
		time.Now().Format(time.RFC3339), len(users))
}

// Middleware for logging requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Log request
		log.Printf("Started %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		
		// Call the next handler
		next.ServeHTTP(w, r)
		
		// Log completion
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// Middleware for CORS headers
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}