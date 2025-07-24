// Lesson 08: Concurrency with Goroutines and Channels
// This lesson covers Go's concurrency primitives and patterns

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Lesson 08: Concurrency with Goroutines and Channels ===")
	
	// Basic goroutines
	fmt.Println("\n--- Basic Goroutines ---")
	demonstrateBasicGoroutines()
	
	// Channels
	fmt.Println("\n--- Channels ---")
	demonstrateChannels()
	
	// Channel directions
	fmt.Println("\n--- Channel Directions ---")
	demonstrateChannelDirections()
	
	// Select statement
	fmt.Println("\n--- Select Statement ---")
	demonstrateSelect()
	
	// Worker pools
	fmt.Println("\n--- Worker Pools ---")
	demonstrateWorkerPool()
	
	// Synchronization primitives
	fmt.Println("\n--- Synchronization Primitives ---")
	demonstrateSynchronization()
	
	// Context for cancellation
	fmt.Println("\n--- Context and Cancellation ---")
	demonstrateContext()
}

func demonstrateBasicGoroutines() {
	// Sequential execution
	fmt.Println("Sequential execution:")
	start := time.Now()
	for i := 0; i < 3; i++ {
		slowTask(fmt.Sprintf("task-%d", i))
	}
	fmt.Printf("Sequential took: %v\n", time.Since(start))
	
	// Concurrent execution with goroutines
	fmt.Println("\nConcurrent execution:")
	start = time.Now()
	var wg sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(taskName string) {
			defer wg.Done()
			slowTask(taskName)
		}(fmt.Sprintf("concurrent-task-%d", i))
	}
	
	wg.Wait() // Wait for all goroutines to complete
	fmt.Printf("Concurrent took: %v\n", time.Since(start))
	
	// Anonymous goroutine
	go func() {
		fmt.Println("Anonymous goroutine executed")
	}()
	
	// Give goroutine time to execute
	time.Sleep(100 * time.Millisecond)
}

func demonstrateChannels() {
	// Unbuffered channel
	fmt.Println("Unbuffered channel:")
	ch := make(chan string)
	
	// Send in a goroutine (prevents blocking)
	go func() {
		ch <- "Hello from goroutine!"
	}()
	
	// Receive from channel
	message := <-ch
	fmt.Printf("Received: %s\n", message)
	
	// Buffered channel
	fmt.Println("\nBuffered channel:")
	bufferedCh := make(chan int, 3)
	
	// Can send without blocking (up to buffer size)
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	
	fmt.Println("Sent 3 values to buffered channel")
	
	// Receive values
	for i := 0; i < 3; i++ {
		value := <-bufferedCh
		fmt.Printf("Received: %d\n", value)
	}
	
	// Channel with range and close
	fmt.Println("\nChannel with range:")
	numberCh := make(chan int, 5)
	
	// Send numbers in a goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			numberCh <- i * i // Send squares
		}
		close(numberCh) // Close channel when done
	}()
	
	// Range over channel (stops when closed)
	for num := range numberCh {
		fmt.Printf("Square: %d\n", num)
	}
}

func demonstrateChannelDirections() {
	// Channel directions for function parameters
	ch := make(chan string, 1)
	
	// Start producer and consumer
	go producer(ch) // Send-only channel in function
	go consumer(ch) // Receive-only channel in function
	
	time.Sleep(2 * time.Second)
	
	// Pipeline pattern
	fmt.Println("\nPipeline pattern:")
	numbers := generateNumbers(5)
	squares := squareNumbers(numbers)
	printNumbers(squares)
}

// Send-only channel parameter
func producer(ch chan<- string) {
	for i := 0; i < 3; i++ {
		ch <- fmt.Sprintf("Message %d", i+1)
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

// Receive-only channel parameter
func consumer(ch <-chan string) {
	for msg := range ch {
		fmt.Printf("Consumed: %s\n", msg)
	}
}

// Pipeline functions
func generateNumbers(count int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= count; i++ {
			ch <- i
		}
	}()
	return ch
}

func squareNumbers(input <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for num := range input {
			ch <- num * num
		}
	}()
	return ch
}

func printNumbers(input <-chan int) {
	for num := range input {
		fmt.Printf("Pipeline result: %d\n", num)
	}
}

func demonstrateSelect() {
	// Select with multiple channels
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	
	// Send to channels with different timing
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Channel 1"
	}()
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "Channel 2"
	}()
	
	// Select receives from whichever channel is ready first
	select {
	case msg1 := <-ch1:
		fmt.Printf("Received from ch1: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Received from ch2: %s\n", msg2)
	}
	
	// Select with timeout
	timeout := time.After(2 * time.Second)
	select {
	case msg := <-ch1:
		fmt.Printf("Received: %s\n", msg)
	case <-timeout:
		fmt.Println("Timeout occurred")
	}
	
	// Non-blocking select with default
	select {
	case msg := <-ch1:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("No message available")
	}
}

func demonstrateWorkerPool() {
	// Create job and result channels
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	
	// Start workers
	numWorkers := 3
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}
	
	// Send jobs
	numJobs := 9
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Collect results
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		
		// Simulate work
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		
		// Send result
		results <- job * 2
	}
}

func demonstrateSynchronization() {
	// Mutex for protecting shared data
	fmt.Println("Mutex example:")
	counter := &SafeCounter{}
	var wg sync.WaitGroup
	
	// Start multiple goroutines that increment counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
			fmt.Printf("Goroutine %d finished\n", n)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.Value())
	
	// Once example
	fmt.Println("\nOnce example:")
	var once sync.Once
	initFunction := func() {
		fmt.Println("This will only be printed once!")
	}
	
	// Call multiple times, but function executes only once
	for i := 0; i < 3; i++ {
		once.Do(initFunction)
	}
	
	// RWMutex example
	fmt.Println("\nRWMutex example:")
	data := &SafeData{data: make(map[string]string)}
	
	// Multiple readers
	for i := 0; i < 3; i++ {
		go func(id int) {
			for j := 0; j < 3; j++ {
				value := data.Read("key")
				fmt.Printf("Reader %d read: %s\n", id, value)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	
	// One writer
	go func() {
		for i := 0; i < 3; i++ {
			data.Write("key", fmt.Sprintf("value-%d", i))
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	time.Sleep(2 * time.Second)
}

// Thread-safe counter using mutex
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// Thread-safe data structure using RWMutex
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
	fmt.Printf("Wrote %s=%s\n", key, value)
}

func demonstrateContext() {
	// Context with cancellation
	fmt.Println("Context with cancellation:")
	
	// This is a simplified example - in real code you'd import "context"
	// For this lesson, we'll simulate context behavior with channels
	
	cancel := make(chan struct{})
	done := make(chan bool)
	
	// Start a cancelable operation
	go func() {
		defer func() { done <- true }()
		
		for i := 0; i < 10; i++ {
			select {
			case <-cancel:
				fmt.Println("Operation cancelled!")
				return
			default:
				fmt.Printf("Working... step %d\n", i+1)
				time.Sleep(200 * time.Millisecond)
			}
		}
		fmt.Println("Operation completed!")
	}()
	
	// Cancel after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		close(cancel)
	}()
	
	<-done
	fmt.Println("Context demonstration finished")
}

// Helper function that simulates slow work
func slowTask(name string) {
	fmt.Printf("Starting %s\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("Completed %s\n", name)
}