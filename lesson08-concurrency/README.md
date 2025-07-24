# Lesson 08: Concurrency with Goroutines and Channels

## Learning Objectives
- Understand goroutines and concurrent execution
- Master channels for communication between goroutines
- Learn channel directions and pipeline patterns
- Use select statements for multiplexing
- Implement worker pools
- Apply synchronization primitives
- Understand context for cancellation

## Key Concepts

### Goroutines

Goroutines are lightweight threads managed by the Go runtime:

**Starting a goroutine:**
```go
go functionName()

// Anonymous function
go func() {
    fmt.Println("Hello from goroutine")
}()

// Function with parameters
go func(name string) {
    fmt.Printf("Hello %s\n", name)
}("Alice")
```

**Waiting for goroutines:**
```go
var wg sync.WaitGroup

for i := 0; i < 3; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        // Work here
    }(i)
}

wg.Wait() // Wait for all goroutines
```

### Channels

Channels provide communication between goroutines:

**Unbuffered channels:**
```go
ch := make(chan string)

// Send (blocks until someone receives)
ch <- "message"

// Receive (blocks until someone sends)
msg := <-ch
```

**Buffered channels:**
```go
ch := make(chan int, 3) // Buffer size 3

ch <- 1  // Doesn't block
ch <- 2  // Doesn't block  
ch <- 3  // Doesn't block
ch <- 4  // Would block (buffer full)
```

**Closing channels:**
```go
close(ch)

// Check if channel is closed
value, ok := <-ch
if !ok {
    fmt.Println("Channel is closed")
}

// Range over channel (stops when closed)
for value := range ch {
    fmt.Println(value)
}
```

### Channel Directions

Restrict channel usage in function parameters:

```go
// Send-only channel
func sender(ch chan<- string) {
    ch <- "message"
}

// Receive-only channel  
func receiver(ch <-chan string) {
    msg := <-ch
}
```

### Select Statement

Multiplex multiple channel operations:

```go
select {
case msg1 := <-ch1:
    fmt.Println("From ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("From ch2:", msg2)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout")
default:
    fmt.Println("No channels ready")
}
```

### Worker Pool Pattern

```go
func workerPool() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collect results
    for r := 1; r <= 9; r++ {
        <-results
    }
}

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        // Process job
        results <- job * 2
    }
}
```

### Synchronization Primitives

**Mutex (mutual exclusion):**
```go
type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}
```

**RWMutex (readers-writer mutex):**
```go
type SafeData struct {
    mu   sync.RWMutex
    data map[string]string
}

func (d *SafeData) Read(key string) string {
    d.mu.RLock()
    defer d.mu.RUnlock()
    return d.data[key]
}

func (d *SafeData) Write(key, value string) {
    d.mu.Lock()
    defer d.mu.Unlock()
    d.data[key] = value
}
```

**Once (execute only once):**
```go
var once sync.Once

once.Do(func() {
    fmt.Println("This runs only once")
})
```

### Pipeline Pattern

```go
// Stage 1: Generate numbers
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

// Stage 2: Square numbers
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * n
        }
    }()
    return out
}

// Usage
numbers := generate(1, 2, 3, 4)
squares := square(numbers)
for result := range squares {
    fmt.Println(result)
}
```

## Concurrency Patterns

1. **Fan-out/Fan-in**: Distribute work among multiple goroutines, then collect results
2. **Pipeline**: Chain operations through channels
3. **Worker Pool**: Fixed number of workers processing jobs
4. **Pub/Sub**: Publisher sends to multiple subscribers
5. **Rate Limiting**: Control the rate of operations

## Best Practices

1. **Don't communicate by sharing memory; share memory by communicating**
2. **Always close channels when you're done sending**
3. **Use buffered channels judiciously**
4. **Prefer channels over locks when possible**
5. **Use context for cancellation and timeouts**
6. **Avoid goroutine leaks - ensure all goroutines can exit**
7. **Use `sync.WaitGroup` to wait for goroutines**

## Common Pitfalls

1. **Goroutine leaks**: Goroutines that never exit
2. **Deadlocks**: Goroutines waiting for each other
3. **Race conditions**: Unsynchronized access to shared data
4. **Channel deadlocks**: Sending/receiving on channels incorrectly

## Running the Code

```bash
cd lesson08-concurrency
go run main.go
```

## Try It Yourself
1. Implement a concurrent web scraper
2. Create a rate limiter using channels
3. Build a pub/sub system with goroutines
4. Implement a concurrent merge sort
5. Create a producer-consumer system with multiple producers and consumers