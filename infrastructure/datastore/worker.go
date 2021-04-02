package datastore

import (
	"log"
	"sync"
)

type Worker struct {
	id                int
	key               chan int
	latestIdx         chan int
	odd               bool
	itemsProcessed    int
	maxItemsPerWorker int
	keys              []int
	shutdown          chan struct{}
	wg                *sync.WaitGroup
}

func (w Worker) Work() {
	log.Println("Starting up worker", w.id)
	for {
		select {
		case <-w.shutdown:
			log.Println("Bye from worker", w.id)
			w.wg.Done()
			return
		case idx := <-w.latestIdx:
			requiredType := false
			for !requiredType {
				idx++
				if len(w.keys) <= idx {
					log.Println("worker", w.id, "no more valid ids.")
					w.wg.Done()
					w.key <- -1
					return
				}
				requiredType = (w.keys[idx] % 2) == 0
				if w.odd {
					requiredType = !requiredType
				}
			}
			log.Println("worker", w.id, "sending", w.keys[idx])
			w.key <- idx
			w.itemsProcessed++
			if w.itemsProcessed == w.maxItemsPerWorker {
				log.Println("worker", w.id, "max items reached.")
				w.wg.Done()
				return
			}
		}
	}
}
