# Lesson 03: Functions and Methods

## Learning Objectives
- Understand function syntax and declaration
- Learn about parameters and return values
- Explore variadic functions and closures
- Understand methods and receivers
- Work with higher-order functions

## Key Concepts

### Function Declaration

**Basic function:**
```go
func functionName(parameter type) returnType {
    // function body
    return value
}
```

**Multiple parameters and return values:**
```go
func divide(a, b int) (int, int) {
    return a/b, a%b
}
```

**Named return values:**
```go
func rectangleStats(w, h float64) (area, perimeter float64) {
    area = w * h
    perimeter = 2 * (w + h)
    return // naked return
}
```

### Variadic Functions

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

### Methods

Methods are functions with a receiver:

**Value receiver:**
```go
func (p Person) Greet() string {
    return "Hello, " + p.Name
}
```

**Pointer receiver (can modify the struct):**
```go
func (p *Person) HaveBirthday() {
    p.Age++
}
```

### Anonymous Functions and Closures

```go
// Anonymous function
multiply := func(a, b int) int {
    return a * b
}

// Closure
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

### Higher-Order Functions

Functions that take other functions as parameters:

```go
func calculate(a, b int, op func(int, int) int) int {
    return op(a, b)
}
```

## Running the Code

```bash
cd lesson03-functions-methods
go run main.go
```

## Try It Yourself
1. Create a function that calculates the area of different shapes
2. Write a method for a custom struct
3. Implement a recursive function (like Fibonacci)
4. Create a closure that maintains state
5. Experiment with defer statements