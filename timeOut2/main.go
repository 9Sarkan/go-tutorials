package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func timeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	temp := make(chan bool)
	go func() {
		defer close(temp)
		time.Sleep(4 * time.Second)
		wg.Wait()
	}()

	select {
	case <-temp:
		return false
	case <-time.After(timeout):
		return true
	}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("you have to send an argument as timeout!")
		return
	}
	timeoutArg, _ := strconv.Atoi(args[1])
	duration := time.Duration(int32(timeoutArg)) * time.Millisecond
	var wg sync.WaitGroup

	wg.Add(1)
	if timeout(&wg, duration) {
		fmt.Println("Time Out!")
	} else {
		fmt.Println("OK!")
	}

	wg.Done()
	if timeout(&wg, duration) {
		fmt.Println("Time Out!")
	} else {
		fmt.Println("OK!")
	}
}
