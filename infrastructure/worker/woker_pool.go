package worker

import (
	"sync"
)

type WorkerPool struct {
	pollSize       int
	itemsPerWorker int
	jobQueue       chan Job
	results        chan<- int
	Workers        []*Worker
	wg             *sync.WaitGroup
	wgSendJobs     *sync.WaitGroup
}

func calculatePoolSize(limit, items_per_worker int) int {
	return limit/items_per_worker + 1
}

func NewWorkerPool(limit, itemsPerWorker int, results chan<- int, wg *sync.WaitGroup) *WorkerPool {

	jobQueue := make(chan Job, limit*itemsPerWorker)

	pollSize := calculatePoolSize(limit, itemsPerWorker)
	return &WorkerPool{pollSize: pollSize, jobQueue: jobQueue, results: results, itemsPerWorker: itemsPerWorker, wg: wg, wgSendJobs: &sync.WaitGroup{}}
}

func (p *WorkerPool) Start() {
	for i := 0; i < p.pollSize; i++ {
		worker := NewWorker(i, p.jobQueue, p.wg, p.results, p.itemsPerWorker)
		p.Workers = append(p.Workers, worker)
		go func(worker *Worker) {
			worker.Start()
		}(worker)
	}

	go func() {
		p.wg.Wait()
		close(p.results)
	}()
}

// AddTask adds a task to the pool
func (p *WorkerPool) AddJob(job Job) {
	p.wgSendJobs.Add(1)
	go func() {
		defer p.wgSendJobs.Done()
		p.jobQueue <- job
	}()

}

func (p *WorkerPool) ShutDown() {
	go func() {
		p.wgSendJobs.Wait()
		close(p.jobQueue)
		//p.Stop()
	}()
}

func (p *WorkerPool) Stop() {

	for i := range p.Workers {
		p.Workers[i].Stop()
	}
}
