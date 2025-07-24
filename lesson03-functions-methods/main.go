// Lesson 03: Functions and Methods
// This lesson covers function definitions, parameters, return values, and methods

package main

import (
	"fmt"
	"math"
)

// Person struct for demonstrating methods
type Person struct {
	Name string
	Age  int
}

// Method with receiver - belongs to Person struct
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old", p.Name, p.Age)
}

// Method with pointer receiver - can modify the struct
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("%s just turned %d!\n", p.Name, p.Age)
}

// Method that returns multiple values
func (p Person) GetInfo() (string, int) {
	return p.Name, p.Age
}

func main() {
	fmt.Println("=== Lesson 03: Functions and Methods ===")
	
	// Simple function calls
	greetings := sayHello("Alice")
	fmt.Println(greetings)
	
	// Function with multiple parameters
	sum := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", sum)
	
	// Function with multiple return values
	quotient, remainder := divide(17, 5)
	fmt.Printf("17 / 5 = %d remainder %d\n", quotient, remainder)
	
	// Using named return values
	area, perimeter := rectangleStats(4, 6)
	fmt.Printf("Rectangle (4x6): Area = %.2f, Perimeter = %.2f\n", area, perimeter)
	
	// Variadic function (variable number of arguments)
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", total)
	
	// Anonymous function (function literal)
	multiply := func(a, b int) int {
		return a * b
	}
	result := multiply(6, 7)
	fmt.Printf("6 * 7 = %d\n", result)
	
	// Higher-order function (function that takes another function as parameter)
	operationResult := calculate(10, 5, add)
	fmt.Printf("Calculate with add: %d\n", operationResult)
	
	operationResult = calculate(10, 5, func(a, b int) int {
		return a * b
	})
	fmt.Printf("Calculate with multiply: %d\n", operationResult)
	
	// Methods demonstration
	fmt.Println("\n=== Methods Demo ===")
	
	// Create a Person instance
	person := Person{Name: "Bob", Age: 30}
	
	// Call method
	fmt.Println(person.Greet())
	
	// Method with multiple return values
	name, age := person.GetInfo()
	fmt.Printf("Person info: %s is %d years old\n", name, age)
	
	// Method with pointer receiver
	person.HaveBirthday()
	fmt.Println(person.Greet()) // Age should be incremented
	
	// Closure example
	counter := createCounter()
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())
}

// Simple function with one parameter and one return value
func sayHello(name string) string {
	return "Hello, " + name + "!"
}

// Function with multiple parameters
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// Function with named return values
func rectangleStats(width, height float64) (area float64, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // naked return - returns the named values
}

// Variadic function (accepts variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Higher-order function (takes a function as parameter)
func calculate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Function that returns a closure
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Recursive function example
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Function demonstrating defer statement
func demonstrateDefer() {
	fmt.Println("Start")
	defer fmt.Println("This will be printed last")
	defer fmt.Println("This will be printed second to last")
	fmt.Println("Middle")
	fmt.Println("End")
	// Deferred functions are executed in LIFO (Last In, First Out) order
}