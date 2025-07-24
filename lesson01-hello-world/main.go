// Lesson 01: Hello World and Basic Syntax
// This lesson covers the fundamental structure of a Go program

package main

import "fmt"

// main is the entry point of every Go program
func main() {
	// Simple hello world
	fmt.Println("Hello, World!")
	
	// Different ways to print
	fmt.Print("Hello ")
	fmt.Print("from ")
	fmt.Println("Go!")
	
	// Printf for formatted output
	fmt.Printf("Hello %s!\n", "Gopher")
	
	// Variables in action
	name := "Golang"
	year := 2009
	fmt.Printf("%s was first released in %d\n", name, year)
}