package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func f1(cc chan chan int, f chan bool) {
	c := make(chan int)
	cc <- c
	defer close(c)
	sum := 0
	select {
	case x := <-c:
		for i := 0; i <= x; i++ {
			sum = sum + i
		}
		c <- sum
	case <-f:
		return
	}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("You have to send an argument for this app!")
		return
	}
	number, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	cc := make(chan chan int)
	for i := 0; i <= number; i++ {
		f := make(chan bool)
		go f1(cc, f)
		ch := <-cc
		ch <- i
		for sum := range ch {
			fmt.Printf("sum(%d): %d\n", i, sum)
		}
		time.Sleep(time.Second)
		close(f)
	}
}
