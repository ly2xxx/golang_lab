// Lesson 10: JSON Handling and REST API
// This lesson covers JSON marshaling/unmarshaling and building RESTful APIs

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// User represents a user in our system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest represents the request payload for creating a user
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// UpdateUserRequest represents the request payload for updating a user
type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Age   *int    `json:"age,omitempty"`
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// ValidationError represents validation errors
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ErrorResponse represents error response with details
type ErrorResponse struct {
	Error   string            `json:"error"`
	Details []ValidationError `json:"details,omitempty"`
}

// In-memory database
var (
	users      = make(map[int]User)
	nextUserID = 1
)

func main() {
	fmt.Println("=== Lesson 10: JSON Handling and REST API ===")
	
	// Initialize with some sample data
	initializeData()
	
	// Demonstrate JSON operations
	demonstratJSON()
	
	// Create HTTP server
	mux := http.NewServeMux()
	registerAPIRoutes(mux)
	
	// Apply middleware
	handler := corsMiddleware(loggingMiddleware(mux))
	
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	
	fmt.Println("\nStarting REST API server on http://localhost:8080")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET    /api/users       - Get all users")
	fmt.Println("  GET    /api/users/{id}  - Get user by ID")
	fmt.Println("  POST   /api/users       - Create new user")
	fmt.Println("  PUT    /api/users/{id}  - Update user")
	fmt.Println("  DELETE /api/users/{id}  - Delete user")
	fmt.Println("  GET    /api/health      - API health check")
	fmt.Println("\nTest with curl:")
	fmt.Println(`  curl http://localhost:8080/api/users`)
	fmt.Println(`  curl -X POST -H "Content-Type: application/json" -d '{"name":"Alice","email":"alice@example.com","age":30}' http://localhost:8080/api/users`)
	fmt.Println("\nPress Ctrl+C to stop the server")
	
	log.Fatal(server.ListenAndServe())
}

func initializeData() {
	// Initialize with sample users
	users[1] = User{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Age:       25,
		CreatedAt: time.Now().Add(-24 * time.Hour),
		UpdatedAt: time.Now().Add(-24 * time.Hour),
	}
	
	users[2] = User{
		ID:        2,
		Name:      "Jane Smith",
		Email:     "jane@example.com",
		Age:       30,
		CreatedAt: time.Now().Add(-12 * time.Hour),
		UpdatedAt: time.Now().Add(-12 * time.Hour),
	}
	
	nextUserID = 3
}

func demonstratJSON() {
	fmt.Println("\n--- JSON Demonstration ---")
	
	// Create a user
	user := User{
		ID:        100,
		Name:      "Demo User",
		Email:     "demo@example.com",
		Age:       28,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	// Marshal to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}
	fmt.Printf("Marshaled JSON: %s\n", string(jsonData))
	
	// Marshal with indentation (pretty print)
	prettyJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}
	fmt.Printf("Pretty JSON:\n%s\n", string(prettyJSON))
	
	// Unmarshal from JSON
	jsonString := `{"id":200,"name":"Test User","email":"test@example.com","age":35,"created_at":"2024-01-01T10:00:00Z","updated_at":"2024-01-01T10:00:00Z"}`
	var unmarshaledUser User
	err = json.Unmarshal([]byte(jsonString), &unmarshaledUser)
	if err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		return
	}
	fmt.Printf("Unmarshaled user: %+v\n", unmarshaledUser)
	
	// Working with maps
	fmt.Println("\n--- JSON with Maps ---")
	data := map[string]interface{}{
		"name":    "Dynamic User",
		"age":     42,
		"active":  true,
		"scores":  []int{95, 87, 92},
		"address": map[string]string{"city": "New York", "country": "USA"},
	}
	
	mapJSON, _ := json.MarshalIndent(data, "", "  ")
	fmt.Printf("Map as JSON:\n%s\n", string(mapJSON))
	
	// Parse JSON into map
	var parsedData map[string]interface{}
	json.Unmarshal(mapJSON, &parsedData)
	fmt.Printf("Parsed back: %+v\n", parsedData)
	
	// Custom JSON tags demonstration
	fmt.Println("\n--- Custom JSON Tags ---")
	type Product struct {
		ID          int     `json:"id"`
		Name        string  `json:"product_name"`
		Price       float64 `json:"price"`
		InStock     bool    `json:"in_stock"`
		Description string  `json:"description,omitempty"`
		Internal    string  `json:"-"` // This field is ignored
	}
	
	product := Product{
		ID:       1,
		Name:     "Laptop",
		Price:    999.99,
		InStock:  true,
		Internal: "This won't be in JSON",
	}
	
	productJSON, _ := json.MarshalIndent(product, "", "  ")
	fmt.Printf("Product JSON:\n%s\n", string(productJSON))
}

func registerAPIRoutes(mux *http.ServeMux) {
	// User routes
	mux.HandleFunc("/api/users", handleUsers)
	mux.HandleFunc("/api/users/", handleUser)
	
	// Health check
	mux.HandleFunc("/api/health", handleHealth)
	
	// API documentation
	mux.HandleFunc("/api", handleAPIDoc)
}

// Handle multiple users (GET /api/users, POST /api/users)
func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAllUsers(w, r)
	case http.MethodPost:
		createUser(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// Handle single user (GET, PUT, DELETE /api/users/{id})
func handleUser(w http.ResponseWriter, r *http.Request) {
	userID, err := extractUserID(r.URL.Path)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		getUser(w, r, userID)
	case http.MethodPut:
		updateUser(w, r, userID)
	case http.MethodDelete:
		deleteUser(w, r, userID)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// GET /api/users
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	userList := make([]User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}
	
	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    userList,
		Message: fmt.Sprintf("Found %d users", len(userList)),
	})
}

// GET /api/users/{id}
func getUser(w http.ResponseWriter, r *http.Request, userID int) {
	user, exists := users[userID]
	if !exists {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	
	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    user,
	})
}

// POST /api/users
func createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	
	// Read and parse JSON body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to read request body")
		return
	}
	defer r.Body.Close()
	
	if err := json.Unmarshal(body, &req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}
	
	// Validate request
	if errors := validateCreateUserRequest(req); len(errors) > 0 {
		respondWithValidationErrors(w, errors)
		return
	}
	
	// Create user
	now := time.Now()
	user := User{
		ID:        nextUserID,
		Name:      req.Name,
		Email:     req.Email,
		Age:       req.Age,
		CreatedAt: now,
		UpdatedAt: now,
	}
	
	users[nextUserID] = user
	nextUserID++
	
	respondWithJSON(w, http.StatusCreated, APIResponse{
		Success: true,
		Data:    user,
		Message: "User created successfully",
	})
}

// PUT /api/users/{id}
func updateUser(w http.ResponseWriter, r *http.Request, userID int) {
	user, exists := users[userID]
	if !exists {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	
	var req UpdateUserRequest
	
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to read request body")
		return
	}
	defer r.Body.Close()
	
	if err := json.Unmarshal(body, &req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}
	
	// Update fields if provided
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Age != nil {
		user.Age = *req.Age
	}
	user.UpdatedAt = time.Now()
	
	users[userID] = user
	
	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    user,
		Message: "User updated successfully",
	})
}

// DELETE /api/users/{id}
func deleteUser(w http.ResponseWriter, r *http.Request, userID int) {
	_, exists := users[userID]
	if !exists {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	
	delete(users, userID)
	
	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "User deleted successfully",
	})
}

// GET /api/health
func handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	
	health := map[string]interface{}{
		"status":     "healthy",
		"timestamp":  time.Now().Format(time.RFC3339),
		"users_count": len(users),
		"version":    "1.0.0",
	}
	
	respondWithJSON(w, http.StatusOK, health)
}

// GET /api
func handleAPIDoc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	
	doc := map[string]interface{}{
		"name":        "User Management API",
		"version":     "1.0.0",
		"description": "RESTful API for managing users with JSON",
		"endpoints": map[string]interface{}{
			"GET /api/users":       "Get all users",
			"GET /api/users/{id}":  "Get user by ID",
			"POST /api/users":      "Create new user",
			"PUT /api/users/{id}":  "Update user",
			"DELETE /api/users/{id}": "Delete user",
			"GET /api/health":      "API health check",
		},
	}
	
	respondWithJSON(w, http.StatusOK, doc)
}

// Helper functions

func extractUserID(path string) (int, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 3 {
		return 0, fmt.Errorf("invalid path")
	}
	return strconv.Atoi(parts[2])
}

func validateCreateUserRequest(req CreateUserRequest) []ValidationError {
	var errors []ValidationError
	
	if strings.TrimSpace(req.Name) == "" {
		errors = append(errors, ValidationError{
			Field:   "name",
			Message: "Name is required",
		})
	}
	
	if strings.TrimSpace(req.Email) == "" {
		errors = append(errors, ValidationError{
			Field:   "email",
			Message: "Email is required",
		})
	} else if !strings.Contains(req.Email, "@") {
		errors = append(errors, ValidationError{
			Field:   "email",
			Message: "Invalid email format",
		})
	}
	
	if req.Age < 0 || req.Age > 150 {
		errors = append(errors, ValidationError{
			Field:   "age",
			Message: "Age must be between 0 and 150",
		})
	}
	
	return errors
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	errorResp := ErrorResponse{
		Error: message,
	}
	respondWithJSON(w, statusCode, errorResp)
}

func respondWithValidationErrors(w http.ResponseWriter, errors []ValidationError) {
	errorResp := ErrorResponse{
		Error:   "Validation failed",
		Details: errors,
	}
	respondWithJSON(w, http.StatusBadRequest, errorResp)
}

// Middleware

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		
		next.ServeHTTP(w, r)
		
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}