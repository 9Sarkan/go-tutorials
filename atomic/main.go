package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
)

type atomCounter struct {
	val int64
}

func (c *atomCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	minusX := flag.Int("x", 100, "go routines")
	minusY := flag.Int("y", 200, "Value")
	flag.Parse()
	X := *minusX
	Y := *minusY
	var wg sync.WaitGroup
	counter := atomCounter{}
	for i := 0; i < X; i++ {
		wg.Add(1)
		go func(no int) {
			defer wg.Done()
			for i := 0; i < Y; i++ {
				atomic.AddInt64(&counter.val, 1)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println(counter.Value())
}
