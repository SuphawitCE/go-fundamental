package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

type Hits struct {
	count int
	sync.Mutex
}

func serve(waitGroup *sync.WaitGroup, hits *Hits, iteration int) {
	wait()
	hits.Lock()
	defer hits.Unlock()
	defer waitGroup.Done()
	hits.count += 1
	fmt.Println("served iteration", iteration)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var waitGroup sync.WaitGroup

	hitCounter := Hits{}
	for i := 0; i < 20; i++ {
		iteration := i
		waitGroup.Add(1)
		go serve(&waitGroup, &hitCounter, iteration)
	}
	fmt.Printf("waiting for goroutines...\n\n")
	waitGroup.Wait()

	hitCounter.Lock()
	totalHits := hitCounter.count
	hitCounter.Unlock()

	fmt.Printf("\ntotal hits = %d\n", totalHits)
}
