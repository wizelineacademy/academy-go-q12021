package main

import (
	"fmt"
	"sort"
	"sync"

	"github.com/jesus-mata/academy-go-q12021/infrastructure/worker"
)

func main() {
	limit := 2000
	itemsPerWorker := 10

	var wg sync.WaitGroup

	//jobs := make(chan worker.Job, limit)
	results := make(chan int, limit)

	workerPool := worker.NewWorkerPool(limit, itemsPerWorker, results, &wg)
	workerPool.Start()

	go func() {
		for i := 1; i <= 1000; i++ {
			job := worker.NewNewsJobFilter(i)
			workerPool.AddJob(job)
		}
		workerPool.ShutDown()
	}()

	var nums []int
	for v := range results {
		nums = append(nums, v)
		if len(nums) == limit {
			workerPool.Stop()
			break
		}
	}
	workerPool.Stop()
	sort.Ints(nums)
	fmt.Printf("Resut: %v - len: %v\n", nums, len(nums))

}
