package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func set(value int) {
	writeValue <- value
}
func read() int {
	return <-readValue
}
func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d - ", value)
		case readValue <- value:
		}
	}
}
func main() {
	number := flag.Int("number", 10, "random numbers")
	flag.Parse()

	num := *number
	fmt.Printf("going to create %d random numbers!\n", num)
	rand.Seed(time.Now().Unix())
	go monitor()
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn(10 * num))
		}()
	}

	wg.Wait()
	fmt.Printf("\nLast Number:\t%d\n", read())
}
