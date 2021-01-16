package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getRandom(min, max int, out chan int, end chan bool) {
	for {
		select {
		case out <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		case <-time.After(4 * time.Second):
			fmt.Println("time.after()")
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	numbers := make(chan int)
	end := make(chan bool)

	if len(os.Args) != 2 {
		fmt.Println("You have to send an arguments with to this program!")
		return
	}
	number, _ := strconv.Atoi(os.Args[1])

	go getRandom(0, 2*number, numbers, end)

	for i := 0; i < number; i++ {
		fmt.Printf("number got: %d\n", <-numbers)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("exiting...")
	end <- true
}
