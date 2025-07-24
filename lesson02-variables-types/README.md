# Lesson 02: Variables, Constants, and Data Types

## Learning Objectives
- Understand Go's variable declaration syntax
- Learn about Go's built-in data types
- Work with constants
- Understand zero values and type inference

## Key Concepts

### Variable Declaration

**Explicit type declaration:**
```go
var name string = "Alice"
var age int = 25
```

**Short variable declaration (type inference):**
```go
name := "Alice"
age := 25
```

**Multiple variables:**
```go
var x, y, z int = 1, 2, 3
a, b := "Hello", "World"
```

### Data Types

**Integer Types:**
- `int8`, `int16`, `int32`, `int64`
- `uint8`, `uint16`, `uint32`, `uint64`
- `int`, `uint` (platform dependent)
- `byte` (alias for uint8)
- `rune` (alias for int32, represents Unicode code point)

**Floating Point:**
- `float32`, `float64`

**Other Types:**
- `bool` (true/false)
- `string`

### Constants

```go
const pi = 3.14159
const greeting = "Hello"

// Grouped constants
const (
    red = "#FF0000"
    green = "#00FF00"
    blue = "#0000FF"
)
```

### Zero Values
Variables declared without initial values get zero values:
- `0` for numeric types
- `false` for boolean
- `""` (empty string) for strings

## Running the Code

```bash
cd lesson02-variables-types
go run main.go
```

## Try It Yourself
1. Create variables for storing personal information
2. Try different data types and see their limits
3. Create a constants group for your application settings
4. Experiment with type conversions