// Lesson 02: Variables, Constants, and Data Types
// This lesson covers Go's type system and variable declarations

package main

import "fmt"

func main() {
	fmt.Println("=== Lesson 02: Variables, Constants, and Data Types ===")
	
	// Variable declarations
	var name string = "Alice"
	var age int = 25
	var isStudent bool = true
	
	// Short variable declaration (type inference)
	height := 5.6
	city := "New York"
	
	// Multiple variable declarations
	var x, y, z int = 1, 2, 3
	a, b := "Hello", "World"
	
	// Zero values (default values when not initialized)
	var defaultInt int
	var defaultString string
	var defaultBool bool
	
	// Constants
	const pi = 3.14159
	const greeting = "Welcome to Go!"
	
	// Multiple constants
	const (
		statusOK = 200
		statusNotFound = 404
		statusError = 500
	)
	
	// Display values
	fmt.Printf("Name: %s, Age: %d, Student: %t\n", name, age, isStudent)
	fmt.Printf("Height: %.1f, City: %s\n", height, city)
	fmt.Printf("Coordinates: (%d, %d, %d)\n", x, y, z)
	fmt.Printf("Message: %s %s\n", a, b)
	
	// Zero values
	fmt.Printf("Default int: %d, string: '%s', bool: %t\n", defaultInt, defaultString, defaultBool)
	
	// Constants
	fmt.Printf("Pi: %f\n", pi)
	fmt.Printf("%s\n", greeting)
	fmt.Printf("HTTP Status Codes: OK=%d, NotFound=%d, Error=%d\n", statusOK, statusNotFound, statusError)
	
	// Data types demonstration
	demonstrateTypes()
}

func demonstrateTypes() {
	fmt.Println("\n=== Data Types Demo ===")
	
	// Integer types
	var int8Val int8 = 127
	var int16Val int16 = 32767
	var int32Val int32 = 2147483647
	var int64Val int64 = 9223372036854775807
	
	// Unsigned integer types
	var uint8Val uint8 = 255
	var uint16Val uint16 = 65535
	
	// Floating point types
	var float32Val float32 = 3.14
	var float64Val float64 = 3.141592653589793
	
	// String and rune types
	var str string = "Hello, 世界"
	var char rune = 'A'
	
	// Byte type (alias for uint8)
	var byteVal byte = 65 // ASCII value for 'A'
	
	fmt.Printf("Integer types: int8=%d, int16=%d, int32=%d, int64=%d\n", int8Val, int16Val, int32Val, int64Val)
	fmt.Printf("Unsigned types: uint8=%d, uint16=%d\n", uint8Val, uint16Val)
	fmt.Printf("Float types: float32=%f, float64=%f\n", float32Val, float64Val)
	fmt.Printf("String: %s\n", str)
	fmt.Printf("Rune (Unicode): %c (%d)\n", char, char)
	fmt.Printf("Byte: %c (%d)\n", byteVal, byteVal)
}