package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type Job struct {
	ID    int
	Value int
}

type Result struct {
	Value string
}

const (
	numJobs   = 5
	numWorker = 2
)

func worker(ctx context.Context, i int, jobQueue <-chan Job, resultQueue chan<- Result) {
	for {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			return
		case job, ok := <-jobQueue:
			if !ok {
				return
			}
			log.Printf("worker %d start job %d", i, job.ID)
			time.Sleep(time.Second * time.Duration(i))
			resultQueue <- Result{Value: fmt.Sprintf("worker %d finished job %d", i, job.ID)}
			log.Printf("worker %d completed job %d", i, job.ID)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	jobQueue := make(chan Job, numJobs)
	resultQueue := make(chan Result, numJobs)

	wg := sync.WaitGroup{}

	// start workers
	for i := 0; i <= numWorker; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(ctx, i, jobQueue, resultQueue)
		}(i)
	}

	for i := 0; i <= numJobs; i++ {
		jobQueue <- Job{
			ID:    i,
			Value: i,
		}
	}
	close(jobQueue)

	go func() {
		wg.Wait()
		close(resultQueue)
	}()

	for res := range resultQueue {
		log.Printf("res: %s", res.Value)
	}
}
