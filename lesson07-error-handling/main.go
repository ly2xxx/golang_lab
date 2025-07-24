// Lesson 07: Error Handling
// This lesson covers Go's error handling patterns and best practices

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Custom error types
type ValidationError struct {
	Field   string
	Message string
}

// Implement error interface
func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error in field '%s': %s", e.Field, e.Message)
}

// Custom error with additional context
type DatabaseError struct {
	Operation string
	Table     string
	Err       error
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("database error during %s on table %s: %v", e.Operation, e.Table, e.Err)
}

// Unwrap method for error wrapping
func (e DatabaseError) Unwrap() error {
	return e.Err
}

// User struct for demonstration
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	fmt.Println("=== Lesson 07: Error Handling ===")
	
	// Basic error handling
	fmt.Println("\n--- Basic Error Handling ---")
	demonstrateBasicErrors()
	
	// Creating custom errors
	fmt.Println("\n--- Custom Errors ---")
	demonstrateCustomErrors()
	
	// Error wrapping and unwrapping
	fmt.Println("\n--- Error Wrapping ---")
	demonstrateErrorWrapping()
	
	// Panic and recover
	fmt.Println("\n--- Panic and Recover ---")
	demonstratePanicRecover()
	
	// Best practices
	fmt.Println("\n--- Error Handling Best Practices ---")
	demonstrateBestPractices()
}

func demonstrateBasicErrors() {
	// Function that returns an error
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}
	
	// Division by zero error
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
	
	// Multiple return values with error
	user, err := findUser(1)
	if err != nil {
		fmt.Printf("Error finding user: %v\n", err)
	} else {
		fmt.Printf("Found user: %+v\n", user)
	}
	
	// File operations (common source of errors)
	content, err := readFile("nonexistent.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	} else {
		fmt.Printf("File content: %s\n", content)
	}
}

func demonstrateCustomErrors() {
	// Using custom error types
	user := User{ID: 1, Name: "", Email: "invalid-email", Age: -5}
	
	err := validateUser(user)
	if err != nil {
		fmt.Printf("Validation failed: %v\n", err)
		
		// Type assertion to get specific error type
		if validationErr, ok := err.(ValidationError); ok {
			fmt.Printf("Field with error: %s\n", validationErr.Field)
		}
	}
	
	// Multiple validation errors
	errors := validateUserComprehensive(user)
	if len(errors) > 0 {
		fmt.Println("Validation errors:")
		for _, err := range errors {
			fmt.Printf("  - %v\n", err)
		}
	}
	
	// Database error example
	err = saveUser(user)
	if err != nil {
		fmt.Printf("Save failed: %v\n", err)
	}
}

func demonstrateErrorWrapping() {
	// Error wrapping with fmt.Errorf
	err := processUserData(0)
	if err != nil {
		fmt.Printf("Process failed: %v\n", err)
		
		// Unwrap the error
		originalErr := errors.Unwrap(err)
		if originalErr != nil {
			fmt.Printf("Original error: %v\n", originalErr)
		}
		
		// Check if error is of specific type
		var dbErr DatabaseError
		if errors.As(err, &dbErr) {
			fmt.Printf("Database operation: %s\n", dbErr.Operation)
		}
		
		// Check if error is a specific error
		if errors.Is(err, ErrUserNotFound) {
			fmt.Println("User not found error detected")
		}
	}
}

func demonstratePanicRecover() {
	// Safe function that recovers from panic
	fmt.Println("Calling function that might panic...")
	
	result := safeOperation(func() interface{} {
		return riskyOperation(10, 0)
	})
	
	if result != nil {
		fmt.Printf("Safe operation result: %v\n", result)
	} else {
		fmt.Println("Operation failed safely")
	}
	
	// Demonstrate panic/recover in a goroutine
	fmt.Println("\nDemonstrating panic handling in goroutine:")
	done := make(chan bool)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic in goroutine: %v\n", r)
			}
			done <- true
		}()
		
		panic("Something went wrong in goroutine!")
	}()
	
	<-done
	fmt.Println("Goroutine completed")
}

func demonstrateBestPractices() {
	// Don't ignore errors
	user, err := findUser(2)
	if err != nil {
		// Handle error appropriately
		fmt.Printf("Could not find user: %v\n", err)
		return
	}
	fmt.Printf("Found user: %s\n", user.Name)
	
	// Error handling in loops
	userIDs := []int{1, 2, 3, 999}
	for _, id := range userIDs {
		user, err := findUser(id)
		if err != nil {
			// Log error but continue processing
			fmt.Printf("Warning: Could not find user %d: %v\n", id, err)
			continue
		}
		fmt.Printf("Processing user: %s\n", user.Name)
	}
	
	// Returning errors early
	err = complexOperation()
	if err != nil {
		fmt.Printf("Complex operation failed: %v\n", err)
	}
}

// Basic function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Predefined errors (package-level)
var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidID    = errors.New("invalid user ID")
)

// Function that uses predefined errors
func findUser(id int) (User, error) {
	if id <= 0 {
		return User{}, ErrInvalidID
	}
	
	// Simulate database lookup
	users := map[int]User{
		1: {ID: 1, Name: "Alice", Email: "alice@example.com", Age: 30},
		2: {ID: 2, Name: "Bob", Email: "bob@example.com", Age: 25},
	}
	
	user, exists := users[id]
	if !exists {
		return User{}, ErrUserNotFound
	}
	
	return user, nil
}

// Function that reads a file and returns error
func readFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return string(content), nil
}

// Function with custom validation error
func validateUser(user User) error {
	if user.Name == "" {
		return ValidationError{Field: "Name", Message: "name cannot be empty"}
	}
	if user.Age < 0 {
		return ValidationError{Field: "Age", Message: "age cannot be negative"}
	}
	if !isValidEmail(user.Email) {
		return ValidationError{Field: "Email", Message: "invalid email format"}
	}
	return nil
}

// Function that returns multiple errors
func validateUserComprehensive(user User) []error {
	var errors []error
	
	if user.Name == "" {
		errors = append(errors, ValidationError{Field: "Name", Message: "name cannot be empty"})
	}
	if user.Age < 0 {
		errors = append(errors, ValidationError{Field: "Age", Message: "age cannot be negative"})
	}
	if user.Age > 150 {
		errors = append(errors, ValidationError{Field: "Age", Message: "age seems unrealistic"})
	}
	if !isValidEmail(user.Email) {
		errors = append(errors, ValidationError{Field: "Email", Message: "invalid email format"})
	}
	
	return errors
}

// Simple email validation
func isValidEmail(email string) bool {
	return len(email) > 0 && email != "invalid-email"
}

// Function that returns wrapped error
func saveUser(user User) error {
	// Simulate database error
	originalErr := errors.New("connection timeout")
	return DatabaseError{
		Operation: "INSERT",
		Table:     "users",
		Err:       originalErr,
	}
}

// Function that demonstrates error wrapping
func processUserData(userID int) error {
	user, err := findUser(userID)
	if err != nil {
		return fmt.Errorf("failed to process user data for ID %d: %w", userID, err)
	}
	
	err = saveUser(user)
	if err != nil {
		return fmt.Errorf("failed to save user %s: %w", user.Name, err)
	}
	
	return nil
}

// Function that might panic
func riskyOperation(a, b int) int {
	if b == 0 {
		panic("division by zero in risky operation")
	}
	return a / b
}

// Safe wrapper that recovers from panic
func safeOperation(operation func() interface{}) (result interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
			result = nil
		}
	}()
	
	return operation()
}

// Complex operation with multiple error points
func complexOperation() error {
	// Simulate multiple operations that could fail
	for i := 0; i < 3; i++ {
		if err := stepOperation(i); err != nil {
			return fmt.Errorf("step %d failed: %w", i, err)
		}
	}
	return nil
}

func stepOperation(step int) error {
	// Simulate some operation that might fail
	time.Sleep(10 * time.Millisecond)
	if step == 2 {
		return errors.New("step 2 always fails in demo")
	}
	return nil
}