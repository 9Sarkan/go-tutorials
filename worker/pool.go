package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}
type Data struct {
	job    Client
	square int
}

var (
	size    = 10
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func worker(w *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		d := Data{job: c, square: square}
		data <- d
		time.Sleep(time.Second)
	}
	w.Done()
}

func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}

func create(n int) {
	for i := 0; i < n; i++ {
		client := Client{i, i}
		clients <- client
	}
	close(clients)
}

func main() {
	job := flag.Int("job", 10, "count of jobs")
	worker := flag.Int("worker", 10, "count of workers")
	flag.Parse()

	go create(*job)
	finished := make(chan interface{})
	go func() {
		for i := range data {
			fmt.Printf("client id:\t%d\n", i.job.id)
			fmt.Printf("%d square:\t%d\n", i.job.integer, i.square)
			fmt.Println("------------")
		}
		finished <- true
	}()
	makeWP(*worker)
	fmt.Printf(": %v\n", <-finished)
}
