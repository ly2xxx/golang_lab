# Lesson 07: Error Handling

## Learning Objectives
- Understand Go's error handling philosophy
- Learn to create and use custom error types
- Master error wrapping and unwrapping
- Use panic and recover appropriately
- Apply error handling best practices

## Key Concepts

### Error Interface

In Go, errors are values that implement the `error` interface:

```go
type error interface {
    Error() string
}
```

### Basic Error Handling

**Function returning an error:**
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

**Handling the error:**
```go
result, err := divide(10, 0)
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}
fmt.Printf("Result: %.2f\n", result)
```

### Creating Errors

**Using `errors.New()`:**
```go
err := errors.New("something went wrong")
```

**Using `fmt.Errorf()` for formatted errors:**
```go
err := fmt.Errorf("failed to process user %d: %s", userID, reason)
```

**Predefined package-level errors:**
```go
var (
    ErrUserNotFound = errors.New("user not found")
    ErrInvalidInput = errors.New("invalid input")
)
```

### Custom Error Types

**Simple custom error:**
```go
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error in %s: %s", e.Field, e.Message)
}
```

**Error with additional methods:**
```go
type DatabaseError struct {
    Operation string
    Table     string
    Err       error
}

func (e DatabaseError) Error() string {
    return fmt.Sprintf("database error during %s on %s: %v", e.Operation, e.Table, e.Err)
}

// Unwrap method for error wrapping
func (e DatabaseError) Unwrap() error {
    return e.Err
}
```

### Error Wrapping

**Wrapping errors with context:**
```go
if err != nil {
    return fmt.Errorf("failed to save user: %w", err)
}
```

**Unwrapping errors:**
```go
originalErr := errors.Unwrap(err)
```

**Checking for specific errors:**
```go
// Check if error is a specific error
if errors.Is(err, ErrUserNotFound) {
    // Handle user not found
}

// Check if error is of specific type
var validationErr ValidationError
if errors.As(err, &validationErr) {
    fmt.Printf("Field: %s\n", validationErr.Field)
}
```

### Panic and Recover

**Panic:**
- Should be used for unrecoverable errors
- Stops normal execution
- Runs deferred functions

```go
if criticalCondition {
    panic("critical system failure")
}
```

**Recover:**
- Can only be called in deferred functions
- Catches panics and allows graceful handling

```go
func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    
    // Code that might panic
    riskyOperation()
}
```

### Error Handling Patterns

**Early return pattern:**
```go
func processData() error {
    data, err := fetchData()
    if err != nil {
        return fmt.Errorf("failed to fetch data: %w", err)
    }
    
    err = validateData(data)
    if err != nil {
        return fmt.Errorf("invalid data: %w", err)
    }
    
    return saveData(data)
}
```

**Error aggregation:**
```go
func validateAll(items []Item) []error {
    var errors []error
    for _, item := range items {
        if err := validate(item); err != nil {
            errors = append(errors, err)
        }
    }
    return errors
}
```

## Best Practices

1. **Always handle errors explicitly**
   ```go
   // DON'T ignore errors
   data, _ := fetchData()
   
   // DO handle errors
   data, err := fetchData()
   if err != nil {
       return err
   }
   ```

2. **Provide context in error messages**
   ```go
   return fmt.Errorf("failed to save user %d to database: %w", userID, err)
   ```

3. **Use custom error types for structured error handling**

4. **Don't panic for normal error conditions**
   - Use panic only for programming errors
   - Prefer returning errors

5. **Wrap errors with context**
   ```go
   if err != nil {
       return fmt.Errorf("operation failed: %w", err)
   }
   ```

6. **Create package-level error variables for common errors**
   ```go
   var ErrNotFound = errors.New("resource not found")
   ```

## Running the Code

```bash
cd lesson07-error-handling
go run main.go
```

## Try It Yourself
1. Create a custom error type for your domain
2. Implement error wrapping in a multi-step operation
3. Write a function that aggregates multiple validation errors
4. Create a safe wrapper function using panic/recover
5. Implement retry logic with error handling