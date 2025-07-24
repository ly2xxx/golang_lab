# Lesson 04: Structs and Interfaces

## Learning Objectives
- Understand struct definition and initialization
- Learn about embedded structs (composition)
- Master interface definition and implementation
- Work with type assertions and type switches
- Explore the empty interface

## Key Concepts

### Structs

Structs are Go's way of creating custom types that group related data:

**Basic struct:**
```go
type Person struct {
    Name string
    Age  int
}
```

**Initialization:**
```go
// Named fields
p1 := Person{Name: "Alice", Age: 30}

// Positional (must match field order)
p2 := Person{"Bob", 25}

// Zero value initialization
var p3 Person
```

**Embedded structs (composition over inheritance):**
```go
type Employee struct {
    Person    // Embedded struct
    ID        int
    Salary    float64
}
```

### Interfaces

Interfaces define method signatures. Any type that implements all methods automatically satisfies the interface:

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

// Rectangle implements Shape interface
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}
```

### Interface Composition

Interfaces can embed other interfaces:

```go
type Describer interface {
    Describe() string
}

type ShapeDescriber interface {
    Shape     // Embedded interface
    Describer // Embedded interface
}
```

### Type Assertions and Type Switches

**Type assertion:**
```go
if rect, ok := shape.(Rectangle); ok {
    // shape is a Rectangle
    fmt.Println(rect.Width)
}
```

**Type switch:**
```go
switch v := value.(type) {
case Rectangle:
    fmt.Println("It's a rectangle")
case Circle:
    fmt.Println("It's a circle")
default:
    fmt.Printf("Unknown type: %T\n", v)
}
```

### Empty Interface

The empty interface `interface{}` can hold any type:

```go
var anything interface{}
anything = 42
anything = "hello"
anything = []int{1, 2, 3}
```

## Running the Code

```bash
cd lesson04-structs-interfaces
go run main.go
```

## Try It Yourself
1. Create a struct for a Book with methods for getting info
2. Define an interface for different types of vehicles
3. Implement the interface for Car and Bicycle structs
4. Use type assertions to handle different vehicle types
5. Create a function that accepts any type using empty interface