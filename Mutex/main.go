package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	m  sync.Mutex
	v1 int
)

func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

func change(x int) {
	m.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1%10 == 0 {
		v1 = v1 - 10*x
	}
	m.Unlock()
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("You have to enter an argument!")
		return
	}
	number, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var wg sync.WaitGroup
	fmt.Printf("-> %d\n", read())
	for i := 0; i < number; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			change(x)
			fmt.Printf("-> %d\n", read())
		}(i)
	}
	wg.Wait()
}
