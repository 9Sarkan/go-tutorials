package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(4 * time.Second)
		c1 <- "c1 ok"
	}()
	select {
	case value := <-c1:
		fmt.Println(value)
	case <-time.After(1 * time.Second):
		fmt.Println("c1 timeout")
	}
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "c2 done"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("c2 timeout!")
	}
}
