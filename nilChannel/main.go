package main

import (
	"fmt"
	"time"
)

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)
	for {
		select {
		case number := <-c:
			sum = sum + number
		case <-t.C:
			c = nil
			fmt.Println(sum)
		}
	}
}

func send(c chan int) {
	for {
		c <- 1
	}
}

func main() {
	c := make(chan int)
	go send(c)
	go add(c)

	time.Sleep(3 * time.Second)
}
