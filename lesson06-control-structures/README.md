# Lesson 06: Control Structures

## Learning Objectives
- Master if/else conditional statements
- Understand different forms of for loops
- Work with switch statements and type switches
- Use range for iteration
- Control program flow with break, continue, and goto
- Understand the select statement for channels

## Key Concepts

### If/Else Statements

**Basic syntax:**
```go
if condition {
    // code
} else if condition2 {
    // code
} else {
    // code
}
```

**If with initialization:**
```go
if x := getValue(); x > 0 {
    fmt.Println("Positive:", x)
}
```

### For Loops

Go has only one loop construct - the `for` loop, but it's versatile:

**Traditional C-style loop:**
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

**While-style loop:**
```go
for condition {
    // code
}
```

**Infinite loop:**
```go
for {
    // code
    if shouldBreak {
        break
    }
}
```

### Switch Statements

**Basic switch:**
```go
switch value {
case 1:
    fmt.Println("One")
case 2, 3:
    fmt.Println("Two or Three")
default:
    fmt.Println("Other")
}
```

**Switch without expression:**
```go
switch {
case x > 0:
    fmt.Println("Positive")
case x < 0:
    fmt.Println("Negative")
default:
    fmt.Println("Zero")
}
```

**Type switch:**
```go
switch v := value.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
default:
    fmt.Printf("Unknown type: %T\n", v)
}
```

**Fallthrough:**
```go
switch grade {
case 'A':
    fmt.Println("Excellent!")
    fallthrough
case 'B':
    fmt.Println("Good job!")
}
```

### Range

The `range` keyword iterates over various data structures:

**Slice/Array:**
```go
for index, value := range slice {
    fmt.Printf("%d: %v\n", index, value)
}
```

**Map:**
```go
for key, value := range myMap {
    fmt.Printf("%s: %v\n", key, value)
}
```

**String (iterates over runes):**
```go
for index, char := range "Hello" {
    fmt.Printf("%d: %c\n", index, char)
}
```

**Channel:**
```go
for value := range channel {
    fmt.Println(value)
}
```

### Control Flow

**Break:**
- Exits the nearest loop or switch
- Can use labels for nested structures

**Continue:**
- Skips to next iteration
- Can use labels for nested structures

**Labels:**
```go
outerLoop:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if condition {
            break outerLoop
        }
    }
}
```

**Goto (use sparingly):**
```go
goto label
// code
label:
// code here
```

### Select Statement

Used for channel operations:

```go
select {
case msg1 := <-ch1:
    fmt.Println("Received:", msg1)
case ch2 <- "hello":
    fmt.Println("Sent message")
case <-time.After(1 * time.Second):
    fmt.Println("Timeout")
default:
    fmt.Println("No channel operation ready")
}
```

## Running the Code

```bash
cd lesson06-control-structures
go run main.go
```

## Best Practices

1. **Use range for iteration when possible**
2. **Prefer switch over long if-else chains**
3. **Use labels sparingly and only when necessary**
4. **Avoid goto in most cases**
5. **Use select for concurrent programming with channels**
6. **Initialize variables in if statements when scope allows**

## Try It Yourself
1. Create a number guessing game using control structures
2. Implement FizzBuzz using different loop types
3. Create a menu system using switch statements
4. Write a function that processes different data types using type switch
5. Implement a simple calculator with switch statements