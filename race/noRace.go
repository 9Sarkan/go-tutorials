package main

import (
	"flag"
	"fmt"
	"sync"
)

var aMutex sync.Mutex

func main() {
	num := flag.Int("n", 10, "number")
	flag.Parse()
	var wg sync.WaitGroup
	k := make(map[int]int)
	for i := 0; i <= *num; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			aMutex.Lock()
			k[j] = j
			aMutex.Unlock()
		}(i)
	}

	wg.Wait()
	k[2] = 22
	fmt.Printf("map value:\t%v\n", k)
}
