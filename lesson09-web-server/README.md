# Lesson 09: Web Server Basics with net/http

## Learning Objectives
- Build HTTP servers using the net/http package
- Handle different HTTP methods and routes
- Work with request and response objects
- Implement middleware for cross-cutting concerns
- Serve static files
- Parse form data and query parameters
- Create proper HTTP responses with status codes

## Key Concepts

### Basic HTTP Server

**Simple server:**
```go
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**Server with configuration:**
```go
server := &http.Server{
    Addr:         ":8080",
    Handler:      mux,
    ReadTimeout:  10 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  60 * time.Second,
}

log.Fatal(server.ListenAndServe())
```

### ServeMux (Router)

**Creating and using a custom mux:**
```go
mux := http.NewServeMux()
mux.HandleFunc("/users", usersHandler)
mux.HandleFunc("/users/", userHandler)
mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
```

### Handler Functions

**Basic handler:**
```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, "<h1>Welcome!</h1>")
}
```

**Method-specific handling:**
```go
func usersHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        // Handle GET
    case http.MethodPost:
        // Handle POST
    default:
        w.Header().Set("Allow", "GET, POST")
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
```

### Request Handling

**URL parameters:**
```go
// Extract from URL path
path := strings.TrimPrefix(r.URL.Path, "/users/")
userID, err := strconv.Atoi(path)
```

**Query parameters:**
```go
name := r.URL.Query().Get("name")
age := r.URL.Query().Get("age")
```

**Form data:**
```go
err := r.ParseForm()
if err != nil {
    http.Error(w, "Failed to parse form", http.StatusBadRequest)
    return
}

name := r.Form.Get("name")
email := r.Form.Get("email")
```

**Headers:**
```go
userAgent := r.Header.Get("User-Agent")
contentType := r.Header.Get("Content-Type")
```

### Response Handling

**Setting headers:**
```go
w.Header().Set("Content-Type", "application/json")
w.Header().Set("Cache-Control", "no-cache")
```

**Status codes:**
```go
w.WriteHeader(http.StatusCreated)        // 201
w.WriteHeader(http.StatusNotFound)       // 404
w.WriteHeader(http.StatusInternalServerError) // 500
```

**Error responses:**
```go
http.Error(w, "User not found", http.StatusNotFound)
http.Error(w, "Invalid input", http.StatusBadRequest)
```

**Redirects:**
```go
http.Redirect(w, r, "/login", http.StatusFound)
http.Redirect(w, r, "/users", http.StatusSeeOther)
```

### Static File Serving

```go
// Serve files from ./static/ directory
fileServer := http.FileServer(http.Dir("./static/"))
mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
```

### Middleware

Middleware wraps handlers to add functionality:

```go
// Logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("Started %s %s", r.Method, r.URL.Path)
        
        next.ServeHTTP(w, r)
        
        log.Printf("Completed in %v", time.Since(start))
    })
}

// Usage
handler := loggingMiddleware(corsMiddleware(mux))
```

**Common middleware patterns:**
- Authentication
- Logging
- CORS headers
- Rate limiting
- Compression
- Request/response modification

### JSON Responses

**Manual JSON:**
```go
w.Header().Set("Content-Type", "application/json")
fmt.Fprintf(w, `{"id":%d,"name":"%s"}`, user.ID, user.Name)
```

**Using json package (preferred):**
```go
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(user)
```

## HTTP Status Codes

- **200 OK**: Request successful
- **201 Created**: Resource created
- **400 Bad Request**: Invalid request
- **401 Unauthorized**: Authentication required
- **403 Forbidden**: Access denied
- **404 Not Found**: Resource not found
- **405 Method Not Allowed**: HTTP method not supported
- **500 Internal Server Error**: Server error

## Running the Server

```bash
cd lesson09-web-server
go run main.go
```

Then visit:
- http://localhost:8080/ - Home page
- http://localhost:8080/hello - Simple greeting
- http://localhost:8080/users - User list (JSON)
- http://localhost:8080/form - User creation form
- http://localhost:8080/static/demo.html - Static file demo

## Testing with curl

```bash
# GET request
curl http://localhost:8080/users

# POST request with form data
curl -X POST -d "name=John&email=john@example.com" http://localhost:8080/users

# GET with query parameters
curl "http://localhost:8080/hello?name=Alice"

# Check headers
curl -I http://localhost:8080/health
```

## Best Practices

1. **Always set appropriate Content-Type headers**
2. **Use proper HTTP status codes**
3. **Validate input data**
4. **Handle errors gracefully**
5. **Use middleware for cross-cutting concerns**
6. **Set timeouts on your server**
7. **Log requests and responses**
8. **Follow RESTful conventions**

## Security Considerations

1. **Input validation and sanitization**
2. **HTTPS in production**
3. **Authentication and authorization**
4. **Rate limiting**
5. **CORS configuration**
6. **Request size limits**
7. **Timeout configuration**

## Try It Yourself

1. Add more CRUD operations (PUT, DELETE)
2. Implement user authentication
3. Add request validation middleware
4. Create a simple template system
5. Add file upload functionality
6. Implement API versioning
7. Add request rate limiting