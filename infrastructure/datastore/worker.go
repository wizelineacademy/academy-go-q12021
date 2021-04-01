package datastore

import (
	"fmt"
	"log"
	"sync"
)

type Worker struct{
	id int
	key chan int
	latestIdx chan int
	mod int
	itemsProcessed int
	maxItemsPerWorker int
	keys []int
	shutdown chan struct{}
	wg *sync.WaitGroup
}

func (w Worker) Work() {
	for {
		select {
		case <- w.shutdown:
			fmt.Printf("Bye from %d\n", w.id)
			w.wg.Done()
			return
		case idx := <-w.latestIdx :
			idx++
			for len(w.keys)>idx && w.keys[idx]%w.mod != 0 {
				log.Println("worker ", w.id, " ??? ", idx)
				idx++
			}
			if len(w.keys) > idx {
				log.Println("worker ", w.id, " sending ", w.keys[idx])
				w.key <- idx
				w.itemsProcessed++
			}
			if w.itemsProcessed == w.maxItemsPerWorker || len(w.keys) <= idx {
				log.Println("worker ", w.id, " max items reached.")
				w.wg.Done()
				return
			}
		}
	}
}