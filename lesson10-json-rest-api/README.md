# Lesson 10: JSON Handling and REST API

## Learning Objectives
- Master JSON marshaling and unmarshaling in Go
- Build a complete RESTful API
- Handle different HTTP methods (GET, POST, PUT, DELETE)
- Implement proper API response structures
- Add request validation and error handling
- Use JSON struct tags effectively
- Apply REST API best practices

## Key Concepts

### JSON Marshaling and Unmarshaling

**Basic marshaling:**
```go
user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
jsonData, err := json.Marshal(user)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(jsonData))
```

**Pretty printing:**
```go
prettyJSON, err := json.MarshalIndent(user, "", "  ")
```

**Unmarshaling:**
```go
jsonString := `{"id":1,"name":"Alice","email":"alice@example.com"}`
var user User
err := json.Unmarshal([]byte(jsonString), &user)
```

**Streaming JSON:**
```go
// Encoding to writer
json.NewEncoder(w).Encode(user)

// Decoding from reader
var user User
json.NewDecoder(r.Body).Decode(&user)
```

### JSON Struct Tags

```go
type User struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    Email       string    `json:"email"`
    Age         int       `json:"age,omitempty"`      // Omit if zero value
    Password    string    `json:"-"`                   // Never include
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

**Common struct tag options:**
- `json:"field_name"` - Custom field name
- `json:",omitempty"` - Omit if zero value
- `json:"-"` - Never marshal/unmarshal
- `json:",string"` - Marshal as string

### RESTful API Design

**HTTP Methods and Endpoints:**
```
GET    /api/users       - Get all users
GET    /api/users/{id}  - Get specific user
POST   /api/users       - Create new user
PUT    /api/users/{id}  - Update user (full update)
PATCH  /api/users/{id}  - Update user (partial update)
DELETE /api/users/{id}  - Delete user
```

**HTTP Status Codes:**
- `200 OK` - Successful GET, PUT, PATCH
- `201 Created` - Successful POST
- `204 No Content` - Successful DELETE
- `400 Bad Request` - Invalid request data
- `404 Not Found` - Resource not found
- `422 Unprocessable Entity` - Validation errors
- `500 Internal Server Error` - Server error

### API Response Structures

**Success response:**
```go
type APIResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
}
```

**Error response:**
```go
type ErrorResponse struct {
    Error   string            `json:"error"`
    Details []ValidationError `json:"details,omitempty"`
}

type ValidationError struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}
```

### Request Handling Patterns

**Reading JSON request body:**
```go
func createUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    
    // Method 1: Read all then unmarshal
    body, err := io.ReadAll(r.Body)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Failed to read body")
        return
    }
    defer r.Body.Close()
    
    if err := json.Unmarshal(body, &req); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid JSON")
        return
    }
    
    // Method 2: Direct decoding
    // if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    //     respondWithError(w, http.StatusBadRequest, "Invalid JSON")
    //     return
    // }
    
    // Process request...
}
```

**Sending JSON responses:**
```go
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    
    if err := json.NewEncoder(w).Encode(data); err != nil {
        log.Printf("Error encoding JSON: %v", err)
    }
}
```

### Input Validation

```go
func validateCreateUserRequest(req CreateUserRequest) []ValidationError {
    var errors []ValidationError
    
    if strings.TrimSpace(req.Name) == "" {
        errors = append(errors, ValidationError{
            Field:   "name",
            Message: "Name is required",
        })
    }
    
    if !isValidEmail(req.Email) {
        errors = append(errors, ValidationError{
            Field:   "email",
            Message: "Invalid email format",
        })
    }
    
    return errors
}
```

### Partial Updates with Pointers

```go
type UpdateUserRequest struct {
    Name  *string `json:"name,omitempty"`
    Email *string `json:"email,omitempty"`
    Age   *int    `json:"age,omitempty"`
}

func updateUser(user *User, req UpdateUserRequest) {
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
}
```

### Error Handling Best Practices

**Consistent error responses:**
```go
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
    errorResp := ErrorResponse{Error: message}
    respondWithJSON(w, statusCode, errorResp)
}

func respondWithValidationErrors(w http.ResponseWriter, errors []ValidationError) {
    errorResp := ErrorResponse{
        Error:   "Validation failed",
        Details: errors,
    }
    respondWithJSON(w, http.StatusBadRequest, errorResp)
}
```

## Running the API

```bash
cd lesson10-json-rest-api
go run main.go
```

## Testing the API

**Using curl:**

```bash
# Get all users
curl http://localhost:8080/api/users

# Get specific user
curl http://localhost:8080/api/users/1

# Create user
curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com","age":30}' \
  http://localhost:8080/api/users

# Update user
curl -X PUT -H "Content-Type: application/json" \
  -d '{"name":"Alice Updated"}' \
  http://localhost:8080/api/users/1

# Delete user
curl -X DELETE http://localhost:8080/api/users/1

# Health check
curl http://localhost:8080/api/health
```

**Using the test script:**
```bash
chmod +x test_api.sh
./test_api.sh
```

## Working with Different Data Types

**Dates and times:**
```go
type Event struct {
    Name      string    `json:"name"`
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
}
```

**Nested structures:**
```go
type Address struct {
    Street  string `json:"street"`
    City    string `json:"city"`
    Country string `json:"country"`
}

type User struct {
    ID      int     `json:"id"`
    Name    string  `json:"name"`
    Address Address `json:"address"`
}
```

**Arrays and slices:**
```go
type User struct {
    ID       int      `json:"id"`
    Name     string   `json:"name"`
    Hobbies  []string `json:"hobbies"`
    Scores   []int    `json:"scores"`
}
```

## API Documentation

Good APIs should be self-documenting:

```go
func handleAPIDoc(w http.ResponseWriter, r *http.Request) {
    doc := map[string]interface{}{
        "name":        "User Management API",
        "version":     "1.0.0",
        "description": "RESTful API for managing users",
        "endpoints": map[string]string{
            "GET /api/users":       "Get all users",
            "POST /api/users":      "Create new user",
            "GET /api/users/{id}":  "Get user by ID",
            "PUT /api/users/{id}":  "Update user",
            "DELETE /api/users/{id}": "Delete user",
        },
    }
    respondWithJSON(w, http.StatusOK, doc)
}
```

## Best Practices

1. **Use consistent response formats**
2. **Implement proper error handling and validation**
3. **Use appropriate HTTP status codes**
4. **Include request/response examples in documentation**
5. **Implement pagination for large datasets**
6. **Use meaningful resource names and URLs**
7. **Version your APIs**
8. **Implement rate limiting and authentication**
9. **Log requests and responses for debugging**
10. **Use JSON struct tags appropriately**

## Try It Yourself

1. Add pagination to the GET /users endpoint
2. Implement user search and filtering
3. Add file upload functionality
4. Create nested resources (e.g., user posts)
5. Implement API authentication with JWT
6. Add request rate limiting
7. Create comprehensive API documentation
8. Add database integration instead of in-memory storage