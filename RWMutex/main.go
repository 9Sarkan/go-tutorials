package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Secret Struct
type Secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

var password = Secret{password: ""}

func change(c *Secret, pass string) {
	c.RWM.Lock()
	fmt.Println("LChange")
	time.Sleep(5 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}

func show(c *Secret) string {
	c.RWM.RLock()
	fmt.Println("Show")
	time.Sleep(2 * time.Second)
	defer c.RWM.RUnlock()
	return c.password
}

func showWithLock(c *Secret) string {
	c.RWM.Lock()
	fmt.Println("ShowWithLock")
	time.Sleep(2 * time.Second)
	defer c.RWM.Unlock()
	return c.password
}

func main() {
	args := os.Args
	var showFunction = func(c *Secret) string { return "" }
	if len(args) != 2 {
		fmt.Println("going to use show")
		showFunction = show
	} else {
		fmt.Println("going to use show with lock")
		showFunction = showWithLock
	}
	var wg sync.WaitGroup
	fmt.Printf("Password:\t%s\n", showFunction(&password))
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("go pass:\t%s\n", showFunction(&password))
		}()
	}

	go func() {
		wg.Add(1)
		defer wg.Done()
		change(&password, "123456")
	}()

	wg.Wait()
	fmt.Printf("password:\t%s\n", showFunction(&password))
}
