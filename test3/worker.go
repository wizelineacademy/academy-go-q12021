package main

import (
	"context"
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	limit := 50
	items_per_worker := 3

	var wg sync.WaitGroup
	//wg.Add(limit)
	jobs := make(chan int, 2)
	results := make(chan int, limit)

	//go worker(jobs, results)
	//go worker(jobs, results)
	//go worker(jobs, results)
	//go worker(jobs, results)

	pollSize := calculatePoolSize(limit, items_per_worker)
	for i := 0; i < pollSize; i++ {
		go workerSingel(i, jobs, results, &wg, items_per_worker)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for i := 0; i < 2000; i++ {
		/*go func(n int) {
			for {
				select {
				case jobs <- n:
					fmt.Printf("Worker sent %d\n", n)
					return
				default:
				}
			}

		}(i)*/
		//jobs <- i
		select {
		case jobs <- i:
			fmt.Println("sent data", i)
		default:
			//fmt.Println("no message sent")
		}
	}

	close(jobs)

	var nums []int
	for v := range results {
		nums = append(nums, v)
		if len(nums) == limit {
			break
		}
	}
	sort.Ints(nums)
	fmt.Printf("Resut: %v - len: %v\n", nums, len(nums))
}

func calculatePoolSize(limit, items_per_worker int) int {
	return limit/items_per_worker + 1
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	for n := range jobs {
		fmt.Printf("=>Worker %v processing %v\n", id, n)
		if isEven(n) {
			results <- n
		}
	}
	fmt.Printf("=>Worker %v STOPPED!!!! \n", id)
}

func workerSingel(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup, items_per_worker int) {
	count := 0
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case n, ok := <-jobs:
			if !ok {
				return
			}
			if isEven(n) {
				fmt.Printf("=>Worker %v processing %v\n", id, n)
				//time.Sleep(3 * time.Second)
				count++
				results <- n
			}

		}
		if count == items_per_worker {
			break
		}
	}
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

//=======================
type Task struct {
	fn  func(context.Context, chan<- interface{}) (context.Context, error)
	ctx context.Context
	Err error
}

func NewTask(fn func(context.Context, chan<- interface{}) (context.Context, error), ctx context.Context) *Task {
	return &Task{fn: fn, ctx: ctx}
}

func (task *Task) run(results chan<- interface{}) interface{} {
	ctx, err := task.fn(task.ctx, results)
	task.Err = err
	return ctx
}

//
//===========================
type Worker struct {
	ID      int
	tasks   <-chan *Task
	results chan<- interface{}
	quit    chan bool
}

func NewPooledExecutor(id int, tasks <-chan *Task, results chan<- interface{}) *Worker {
	return &Worker{
		ID:      id,
		tasks:   tasks,
		results: results,
		quit:    make(chan bool),
	}
}

func (wr *Worker) Start(wg *sync.WaitGroup) {
	for {
		select {
		case task := <-wr.tasks:
			fmt.Printf("Running task with worker %v - goid: %v\n", wr.ID, goid())
			task.run(wr.results)
		case <-wr.quit:
			wg.Done()
			return
		}
	}
}

// Stop quits the worker
func (wr *Worker) Stop() {
	fmt.Printf("Closing worker %d\n", wr.ID)
	go func() {
		wr.quit <- true
	}()
}

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
