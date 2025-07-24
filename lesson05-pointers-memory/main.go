// Lesson 05: Pointers and Memory Management
// This lesson covers Go's pointer system and memory management

package main

import (
	"fmt"
	"unsafe"
)

// Struct for demonstrating pointer operations
type Counter struct {
	Value int
}

// Method with value receiver
func (c Counter) IncrementValue() {
	c.Value++ // This won't modify the original
	fmt.Printf("Inside IncrementValue (value receiver): %d\n", c.Value)
}

// Method with pointer receiver
func (c *Counter) IncrementPointer() {
	c.Value++ // This modifies the original
	fmt.Printf("Inside IncrementPointer (pointer receiver): %d\n", c.Value)
}

func main() {
	fmt.Println("=== Lesson 05: Pointers and Memory Management ===")
	
	// Basic pointer operations
	fmt.Println("\n--- Basic Pointers ---")
	
	// Declare a variable
	x := 42
	fmt.Printf("x = %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	
	// Declare a pointer
	var p *int
	fmt.Printf("Zero value of pointer: %v\n", p)
	
	// Assign address to pointer
	p = &x
	fmt.Printf("p points to address: %p\n", p)
	fmt.Printf("Value at address p points to: %d\n", *p)
	
	// Modify value through pointer
	*p = 100
	fmt.Printf("After modifying through pointer: x = %d\n", x)
	
	// Pointer to pointer
	pp := &p
	fmt.Printf("Address of p: %p\n", &p)
	fmt.Printf("pp points to: %p\n", pp)
	fmt.Printf("Value pp points to: %p\n", *pp)
	fmt.Printf("Value of what pp ultimately points to: %d\n", **pp)
	
	// Array and slice pointers
	fmt.Println("\n--- Array and Slice Pointers ---")
	
	arr := [3]int{10, 20, 30}
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Address of array: %p\n", &arr)
	fmt.Printf("Address of first element: %p\n", &arr[0])
	
	// Pointer to array
	arrPtr := &arr
	fmt.Printf("Through pointer: %v\n", *arrPtr)
	(*arrPtr)[1] = 200
	fmt.Printf("Modified array: %v\n", arr)
	
	// Slices (already reference types)
	slice := []int{1, 2, 3}
	fmt.Printf("Slice: %v\n", slice)
	modifySlice(slice)
	fmt.Printf("After function call: %v\n", slice)
	
	// Struct pointers
	fmt.Println("\n--- Struct Pointers ---")
	
	counter := Counter{Value: 5}
	fmt.Printf("Initial counter: %d\n", counter.Value)
	
	// Method with value receiver
	counter.IncrementValue()
	fmt.Printf("After IncrementValue: %d\n", counter.Value)
	
	// Method with pointer receiver
	counter.IncrementPointer()
	fmt.Printf("After IncrementPointer: %d\n", counter.Value)
	
	// Direct pointer operations
	counterPtr := &counter
	counterPtr.Value = 10 // Go automatically dereferences
	fmt.Printf("After direct modification: %d\n", counter.Value)
	
	// Function parameter passing
	fmt.Println("\n--- Function Parameter Passing ---")
	
	a := 5
	b := 10
	fmt.Printf("Before: a=%d, b=%d\n", a, b)
	
	// Pass by value
	swapValues(a, b)
	fmt.Printf("After swapValues: a=%d, b=%d\n", a, b)
	
	// Pass by pointer
	swapPointers(&a, &b)
	fmt.Printf("After swapPointers: a=%d, b=%d\n", a, b)
	
	// Dynamic memory allocation
	fmt.Println("\n--- Dynamic Memory Allocation ---")
	
	// Using new() - returns pointer to zero value
	numPtr := new(int)
	fmt.Printf("new(int): %p, value: %d\n", numPtr, *numPtr)
	*numPtr = 42
	fmt.Printf("After assignment: %d\n", *numPtr)
	
	// Using make() for slices, maps, channels
	slicePtr := make([]int, 3, 5)
	fmt.Printf("make([]int, 3, 5): %v, len=%d, cap=%d\n", slicePtr, len(slicePtr), cap(slicePtr))
	
	// Struct allocation
	counter2 := new(Counter)
	counter2.Value = 15
	fmt.Printf("New counter: %d\n", counter2.Value)
	
	// Address allocation with &
	counter3 := &Counter{Value: 20}
	fmt.Printf("Address operator allocation: %d\n", counter3.Value)
	
	// Memory size and alignment
	fmt.Println("\n--- Memory Information ---")
	demonstrateSizes()
	
	// Nil pointers
	fmt.Println("\n--- Nil Pointers ---")
	var nilPtr *int
	fmt.Printf("Nil pointer: %v\n", nilPtr)
	fmt.Printf("Is nil? %t\n", nilPtr == nil)
	
	// Checking for nil before dereferencing
	if nilPtr != nil {
		fmt.Printf("Value: %d\n", *nilPtr)
	} else {
		fmt.Println("Cannot dereference nil pointer")
	}
	
	// Pointer arithmetic (limited in Go)
	fmt.Println("\n--- Unsafe Pointers (Advanced) ---")
	unsafePointerDemo()
}

// Function that modifies slice (reference type)
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
}

// Function with value parameters (doesn't modify originals)
func swapValues(x, y int) {
	x, y = y, x
	fmt.Printf("Inside swapValues: x=%d, y=%d\n", x, y)
}

// Function with pointer parameters (modifies originals)
func swapPointers(x, y *int) {
	*x, *y = *y, *x
	fmt.Printf("Inside swapPointers: x=%d, y=%d\n", *x, *y)
}

// Demonstrate type sizes
func demonstrateSizes() {
	fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(int(0)))
	fmt.Printf("Size of float64: %d bytes\n", unsafe.Sizeof(float64(0)))
	fmt.Printf("Size of string: %d bytes\n", unsafe.Sizeof(string("")))
	fmt.Printf("Size of []int: %d bytes\n", unsafe.Sizeof([]int{}))
	fmt.Printf("Size of Counter struct: %d bytes\n", unsafe.Sizeof(Counter{}))
	fmt.Printf("Size of pointer: %d bytes\n", unsafe.Sizeof(&Counter{}))
	
	// Alignment
	counter := Counter{}
	fmt.Printf("Alignment of Counter: %d\n", unsafe.Alignof(counter))
	fmt.Printf("Offset of Value field: %d\n", unsafe.Offsetof(counter.Value))
}

// Unsafe pointer operations (advanced topic)
func unsafePointerDemo() {
	fmt.Println("\nWarning: This demonstrates unsafe operations!")
	
	// Convert between different pointer types using unsafe.Pointer
	x := int64(42)
	ptr := unsafe.Pointer(&x)
	
	// Convert to *int32 (this is unsafe!)
	int32Ptr := (*int32)(ptr)
	fmt.Printf("Original int64: %d\n", x)
	fmt.Printf("As int32: %d\n", *int32Ptr)
	
	// Convert pointer to uintptr for arithmetic
	addr := uintptr(ptr)
	fmt.Printf("Memory address as uintptr: %d (0x%x)\n", addr, addr)
	
	// Note: In real applications, avoid unsafe operations unless absolutely necessary
	// They break Go's type safety and can lead to undefined behavior
}