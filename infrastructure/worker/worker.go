package worker

import (
	"sync"
)

type Worker struct {
	ID             int
	jobs           <-chan Job
	results        chan<- int
	quit           chan bool
	wg             *sync.WaitGroup
	itemsPerWorker int
}

func NewWorker(id int, jobs <-chan Job, wg *sync.WaitGroup, results chan<- int, itemsPerWorker int) *Worker {
	return &Worker{
		ID:             id,
		jobs:           jobs,
		results:        results,
		quit:           make(chan bool),
		itemsPerWorker: itemsPerWorker,
		wg:             wg,
	}
}

func (w *Worker) Start() {
	count := 0
	w.wg.Add(1)
	defer w.wg.Done()
	for {
		select {
		case n, ok := <-w.jobs:
			if !ok {
				return
			}
			if n.Apply() {
				//fmt.Printf("=>Worker %v processing \n", w.ID)
				//time.Sleep(3 * time.Second)
				count++
				w.results <- n.GetData()
			}
		case <-w.quit:
			return
		}
		if count == w.itemsPerWorker {
			//fmt.Println("Break with value!!! ", count)
			return
		}
	}

}

// Stop quits the worker
func (wr *Worker) Stop() {
	go func() {
		wr.quit <- true
	}()
}
