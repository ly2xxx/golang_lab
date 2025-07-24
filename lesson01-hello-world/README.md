# Lesson 01: Hello World and Basic Syntax

## Learning Objectives
- Understand the basic structure of a Go program
- Learn about packages and imports
- Use the fmt package for output
- Write your first Go program

## Key Concepts

### Package Declaration
Every Go file starts with a package declaration. The `main` package is special - it tells Go that this is an executable program.

```go
package main
```

### Imports
The `import` statement tells Go which packages your program needs.

```go
import "fmt"
```

### Main Function
The `main` function is the entry point of your program.

```go
func main() {
    // Your code here
}
```

## Running the Code

```bash
go run main.go
```

## Expected Output
```
Hello, World!
Hello from Go!
Hello Gopher!
Golang was first released in 2009
```

## Try It Yourself
1. Modify the greeting message
2. Add your own name to the output
3. Try different formatting with Printf