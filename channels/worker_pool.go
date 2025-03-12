package channels

import (
	"fmt"
	"math/rand"
	"time"
)

//  Worker Pool distributes jobs across a fixed number of worker Goroutines.
// A worker pool is a pattern where multiple worker Goroutines process tasks from a shared queue (channel)
// 🔹 Why Use a Worker Pool?
// Efficient Resource Utilization: Limits Goroutines to avoid excessive memory and CPU usage.
// Better Concurrency Control: Instead of launching Goroutines for every task, it distributes work among fixed workers.
// Scalability: Can handle a large number of tasks with a limited number of workers.

// Each worker listens on a job channel, processes tasks, and sends results to a result channel.
func workerGoroutine(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs { // Receive from jobs channel
		fmt.Printf("🛠️ Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2 // Send result
	}
}

// jobs <-chan int: Receive-only (workers can only read jobs).
// results chan<- int: Send-only (workers can only send results).
// Workers process jobs and send results back.

func WorkerPool() {
	fmt.Println("WorkPool Demo...")
	rand.Seed(time.Now().UnixNano()) // Seed for random sleep times

	const numWorkers = 3 // Number of worker Goroutines
	const numJobs = 5    // Number of tasks

	jobs := make(chan int, numJobs)    // Buffered channel for jobs
	results := make(chan int, numJobs) // Buffered channel for results

	// Start worker Goroutines
	for w := 1; w <= numWorkers; w++ {
		go workerGoroutine(w, jobs, results)
	}

	// Send jobs to the workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // No more jobs to process

	// Collect results
	for r := 1; r <= numJobs; r++ {
		fmt.Println("✅ Processed result:", <-results)
	}

	fmt.Println("🚀 All jobs processed!")
}
