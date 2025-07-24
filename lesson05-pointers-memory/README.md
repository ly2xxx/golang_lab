# Lesson 05: Pointers and Memory Management

## Learning Objectives
- Understand pointer declaration and dereferencing
- Learn the difference between value and pointer receivers
- Master memory allocation with `new()` and `make()`
- Understand Go's memory model and garbage collection
- Work with nil pointers safely

## Key Concepts

### Pointer Basics

**Declaration and usage:**
```go
var x int = 42
var p *int    // Pointer to int
p = &x        // Address-of operator
fmt.Println(*p) // Dereference operator
```

**Zero value of pointers is `nil`:**
```go
var p *int
fmt.Println(p == nil) // true
```

### Pointer vs Value Receivers

**Value receiver (doesn't modify original):**
```go
func (c Counter) Increment() {
    c.Value++ // Only modifies copy
}
```

**Pointer receiver (modifies original):**
```go
func (c *Counter) Increment() {
    c.Value++ // Modifies original
}
```

### Memory Allocation

**Using `new()` - returns pointer to zero value:**
```go
p := new(int)    // *p is 0
s := new(string) // *s is ""
```

**Using `make()` - for slices, maps, channels:**
```go
slice := make([]int, 5, 10) // length=5, capacity=10
map1 := make(map[string]int)
chan1 := make(chan int)
```

**Using address operator `&`:**
```go
s := &Student{Name: "Alice", Age: 20}
```

### Pointer Arithmetic

Go has limited pointer arithmetic compared to C/C++:
- No pointer arithmetic with `+`, `-` operators
- Use `unsafe` package for low-level operations (not recommended)
- Slices provide safe array-like access

### Function Parameters

**Pass by value (copies the value):**
```go
func modify(x int) {
    x = 100 // Doesn't affect original
}
```

**Pass by pointer (can modify original):**
```go
func modify(x *int) {
    *x = 100 // Modifies original
}
```

### Reference Types

Some types are already "reference-like":
- Slices (header points to underlying array)
- Maps
- Channels
- Interfaces
- Functions

### Memory Safety

**Always check for nil:**
```go
if ptr != nil {
    fmt.Println(*ptr)
}
```

**Go's garbage collector:**
- Automatic memory management
- No manual memory deallocation needed
- Handles circular references

### Unsafe Package

The `unsafe` package allows:
- Pointer arithmetic
- Type conversions
- Memory layout inspection

**Use with extreme caution:**
```go
size := unsafe.Sizeof(myStruct)
offset := unsafe.Offsetof(myStruct.field)
```

## Running the Code

```bash
cd lesson05-pointers-memory
go run main.go
```

## Best Practices

1. **Use pointers when:**
   - Need to modify the receiver in methods
   - Dealing with large structs (avoid copying)
   - Need to represent "optional" values

2. **Avoid pointers when:**
   - Working with small values
   - Don't need to modify the original
   - Can use value semantics

3. **Always check for nil before dereferencing**

4. **Prefer value semantics when possible**

## Try It Yourself
1. Create a linked list using pointers
2. Implement a binary tree with pointer-based nodes
3. Write functions that demonstrate the difference between value and pointer parameters
4. Create a struct with both value and pointer receivers
5. Experiment with slice modifications through functions