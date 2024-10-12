
## A Simple Concurrent Increment Example

Concurrent Increment Example
This Go program demonstrates a concurrency issue involving the atomicity of operations. It increments a shared counter (count) using 1 million goroutines. Without proper synchronization, this operation would result in a race condition.

Problem
In Go, count++ is `not an atomic operation`. Without synchronization, multiple goroutines could try to increment count simultaneously, leading to inconsistent results. This example uses a mutex (sync.Mutex) to ensure that increments are safely synchronized.

## Key Concepts
`sync.WaitGroup`: Used to ensure the main goroutine waits until all spawned goroutines complete their work before proceeding.

- `wg.Add(1)`: Adds a counter to track a goroutine.
- `wg.Done()`: Decrements the counter when a goroutine finishes.
- `wg.Wait()`: Blocks the main goroutine until the counter reaches zero.
- 
`sync.Mutex`: A lock to ensure mutual exclusion, preventing race conditions.

- `mu.Lock()`: Locks the mutex, ensuring only one goroutine can access the critical section.
- `mu.Unlock()`: Unlocks the mutex, allowing other goroutines to access the critical section.

`Goroutines`: Lightweight threads that execute functions concurrently.# locking-mutexes-go
