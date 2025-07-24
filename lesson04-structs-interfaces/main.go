// Lesson 04: Structs and Interfaces
// This lesson covers Go's approach to object-oriented programming

package main

import (
	"fmt"
	"math"
)

// Basic struct definition
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

// Struct with embedded fields (composition)
type Address struct {
	Street   string
	City     string
	ZipCode  string
	Country  string
}

type Employee struct {
	Person    // Embedded struct (anonymous field)
	Address   // Embedded struct
	ID        int
	Salary    float64
	Department string
}

// Interface definition
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Interface for objects that can be described
type Describer interface {
	Describe() string
}

// Rectangle struct implementing Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct implementing Shape interface
type Circle struct {
	Radius float64
}

// Implementing Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Implementing Describer interface for Rectangle
func (r Rectangle) Describe() string {
	return fmt.Sprintf("Rectangle with width %.2f and height %.2f", r.Width, r.Height)
}

// Implementing Shape interface for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Implementing Describer interface for Circle
func (c Circle) Describe() string {
	return fmt.Sprintf("Circle with radius %.2f", c.Radius)
}

// Methods for Person struct
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// Method for Employee struct
func (e Employee) GetDetails() string {
	return fmt.Sprintf("Employee ID: %d, Name: %s, Department: %s, Salary: $%.2f",
		e.ID, e.FullName(), e.Department, e.Salary)
}

// Function that works with any Shape
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
	
	// Type assertion to check if shape also implements Describer
	if describer, ok := s.(Describer); ok {
		fmt.Printf("Description: %s\n", describer.Describe())
	}
}

// Interface composition
type ShapeDescriber interface {
	Shape     // Embedded interface
	Describer // Embedded interface
}

func main() {
	fmt.Println("=== Lesson 04: Structs and Interfaces ===")
	
	// Creating struct instances
	fmt.Println("\n--- Basic Structs ---")
	
	// Different ways to create structs
	person1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Email:     "john.doe@example.com",
	}
	
	// Using field names (order doesn't matter)
	person2 := Person{
		Age:       25,
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane.smith@example.com",
	}
	
	// Positional initialization (must match field order)
	person3 := Person{"Bob", "Johnson", 35, "bob.johnson@example.com"}
	
	fmt.Printf("Person 1: %s, Age: %d, Adult: %t\n", person1.FullName(), person1.Age, person1.IsAdult())
	fmt.Printf("Person 2: %s, Age: %d, Adult: %t\n", person2.FullName(), person2.Age, person2.IsAdult())
	fmt.Printf("Person 3: %s, Age: %d, Adult: %t\n", person3.FullName(), person3.Age, person3.IsAdult())
	
	// Embedded structs (composition)
	fmt.Println("\n--- Embedded Structs ---")
	
	employee := Employee{
		Person: Person{
			FirstName: "Alice",
			LastName:  "Brown",
			Age:       28,
			Email:     "alice.brown@company.com",
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "10001",
			Country: "USA",
		},
		ID:         1001,
		Salary:     75000.0,
		Department: "Engineering",
	}
	
	fmt.Println(employee.GetDetails())
	fmt.Printf("Lives in: %s, %s\n", employee.City, employee.Country)
	
	// Can access embedded fields directly
	fmt.Printf("Employee's full name: %s\n", employee.FullName())
	
	// Anonymous struct
	product := struct {
		Name  string
		Price float64
	}{
		Name:  "Laptop",
		Price: 999.99,
	}
	fmt.Printf("Product: %s, Price: $%.2f\n", product.Name, product.Price)
	
	// Interfaces demonstration
	fmt.Println("\n--- Interfaces ---")
	
	// Creating shapes
	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}
	
	// Using interface
	shapes := []Shape{rectangle, circle}
	
	for i, shape := range shapes {
		fmt.Printf("\nShape %d:\n", i+1)
		printShapeInfo(shape)
	}
	
	// Type assertion and type switch
	fmt.Println("\n--- Type Assertions and Switches ---")
	
	var shape Shape = Rectangle{Width: 10, Height: 5}
	
	// Type assertion
	if rect, ok := shape.(Rectangle); ok {
		fmt.Printf("It's a rectangle with width: %.2f\n", rect.Width)
	}
	
	// Type switch
	identifyShape(rectangle)
	identifyShape(circle)
	identifyShape("not a shape")
	
	// Empty interface
	fmt.Println("\n--- Empty Interface ---")
	demonstrateEmptyInterface()
}

// Function demonstrating type switch
func identifyShape(s interface{}) {
	switch v := s.(type) {
	case Rectangle:
		fmt.Printf("Rectangle: %.2f x %.2f\n", v.Width, v.Height)
	case Circle:
		fmt.Printf("Circle with radius: %.2f\n", v.Radius)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// Demonstrating empty interface (interface{})
func demonstrateEmptyInterface() {
	// Empty interface can hold any type
	var anything interface{}
	
	anything = 42
	fmt.Printf("anything = %v (type: %T)\n", anything, anything)
	
	anything = "hello"
	fmt.Printf("anything = %v (type: %T)\n", anything, anything)
	
	anything = []int{1, 2, 3}
	fmt.Printf("anything = %v (type: %T)\n", anything, anything)
	
	// Slice of empty interfaces
	mixedSlice := []interface{}{1, "hello", 3.14, true, Rectangle{2, 3}}
	fmt.Println("Mixed slice:")
	for i, item := range mixedSlice {
		fmt.Printf("  [%d]: %v (type: %T)\n", i, item, item)
	}
}