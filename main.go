package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	JobID int
	Value int
}

func worker(id int, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d dang xu ly job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)

		results <- Result{
			JobID: job,
			Value: job * job,
		}
	}
}

func main() {
	const workerCount = 3
	const jobCount = 8

	jobs := make(chan int, jobCount)
	results := make(chan Result, jobCount)

	var wg sync.WaitGroup

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for job := 1; job <= jobCount; job++ {
		jobs <- job
	}
	close(jobs)

	wg.Wait()
	close(results)

	fmt.Println("Ket qua:")
	for result := range results {
		fmt.Printf("Job %d => %d\n", result.JobID, result.Value)
	}
}
