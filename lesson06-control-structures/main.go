// Lesson 06: Control Structures
// This lesson covers if/else, loops, switch statements, and control flow

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== Lesson 06: Control Structures ===")
	
	// If/else statements
	fmt.Println("\n--- If/Else Statements ---")
	demonstrateIfElse()
	
	// For loops
	fmt.Println("\n--- For Loops ---")
	demonstrateForLoops()
	
	// Switch statements
	fmt.Println("\n--- Switch Statements ---")
	demonstrateSwitch()
	
	// Range loops
	fmt.Println("\n--- Range Loops ---")
	demonstrateRange()
	
	// Control flow statements
	fmt.Println("\n--- Control Flow (break, continue, goto) ---")
	demonstrateControlFlow()
	
	// Select statement (for channels)
	fmt.Println("\n--- Select Statement ---")
	demonstrateSelect()
}

func demonstrateIfElse() {
	// Basic if statement
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}
	
	// If-else
	y := 3
	if y%2 == 0 {
		fmt.Println("y is even")
	} else {
		fmt.Println("y is odd")
	}
	
	// If-else if-else
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}
	
	// If with initialization statement
	if num := rand.Intn(100); num < 50 {
		fmt.Printf("Random number %d is less than 50\n", num)
	} else {
		fmt.Printf("Random number %d is 50 or greater\n", num)
	}
	
	// If with multiple conditions
	age := 25
	income := 50000
	if age >= 18 && income > 30000 {
		fmt.Println("Eligible for loan")
	}
	
	// If with type assertion
	var value interface{} = "hello"
	if str, ok := value.(string); ok {
		fmt.Printf("Value is a string: %s\n", str)
	}
}

func demonstrateForLoops() {
	// Traditional for loop
	fmt.Println("Traditional for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	
	// For loop as while
	fmt.Println("For loop as while:")
	counter := 0
	for counter < 3 {
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}
	
	// Infinite loop (break to exit)
	fmt.Println("Infinite loop with break:")
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Printf("Iteration: %d\n", i)
		i++
	}
	
	// Nested loops
	fmt.Println("Nested loops (multiplication table):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
	
	// Loop with continue
	fmt.Println("Loop with continue (skip even numbers):")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func demonstrateSwitch() {
	// Basic switch
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Middle of work week")
	case "Friday":
		fmt.Println("TGIF!")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Invalid day")
	}
	
	// Switch with initialization
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
	
	// Switch on type
	var value interface{} = 42
	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
	
	// Switch without expression (like if-else)
	num := 15
	switch {
	case num < 10:
		fmt.Println("Single digit")
	case num < 100:
		fmt.Println("Double digit")
	case num < 1000:
		fmt.Println("Triple digit")
	default:
		fmt.Println("Big number")
	}
	
	// Switch with fallthrough
	grade := 'B'
	switch grade {
	case 'A':
		fmt.Println("Excellent!")
		fallthrough
	case 'B':
		fmt.Println("Good job!")
		fallthrough
	case 'C':
		fmt.Println("You passed!")
	case 'D':
		fmt.Println("You barely passed")
	case 'F':
		fmt.Println("You failed")
	default:
		fmt.Println("Invalid grade")
	}
}

func demonstrateRange() {
	// Range over slice
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Range over slice:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	
	// Range with blank identifier (ignore index)
	fmt.Println("\nRange ignoring index:")
	for _, value := range numbers {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
	
	// Range ignoring value
	fmt.Println("\nRange ignoring value:")
	for index := range numbers {
		fmt.Printf("Index: %d ", index)
	}
	fmt.Println()
	
	// Range over string (iterates over runes)
	fmt.Println("\nRange over string:")
	str := "Hello, 世界!"
	for index, char := range str {
		fmt.Printf("Index: %d, Char: %c (Unicode: %U)\n", index, char, char)
	}
	
	// Range over map
	fmt.Println("\nRange over map:")
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	for key, value := range colors {
		fmt.Printf("Color: %s, Hex: %s\n", key, value)
	}
	
	// Range over channel (will block until channel is closed)
	fmt.Println("\nRange over channel:")
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
}

func demonstrateControlFlow() {
	// Break and continue in nested loops
	fmt.Println("Break and continue in nested loops:")
	
	outerLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Println("Breaking out of outer loop")
				break outerLoop // Label to break out of outer loop
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
	
	// Continue with label
	fmt.Println("\nContinue with label:")
	
	outerContinue:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				fmt.Printf("Skipping inner loop for i=%d, j=%d\n", i, j)
				continue outerContinue // Continue outer loop
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
	
	// Goto statement (use sparingly)
	fmt.Println("\nGoto statement (not recommended):")
	i := 0
	
loop:
	fmt.Printf("i = %d\n", i)
	i++
	if i < 3 {
		goto loop
	}
}

func demonstrateSelect() {
	// Select statement for channel operations
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	
	// Send values to channels
	ch1 <- "Channel 1"
	ch2 <- "Channel 2"
	
	// Select statement
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg2)
		default:
			fmt.Println("No channels ready")
		}
	}
	
	// Select with timeout
	timeout := time.After(1 * time.Second)
	select {
	case <-ch1:
		fmt.Println("Received from ch1")
	case <-timeout:
		fmt.Println("Timeout occurred")
	}
	
	// Non-blocking channel operation
	select {
	case ch1 <- "New message":
		fmt.Println("Sent message to ch1")
	default:
		fmt.Println("ch1 is full, cannot send")
	}
}