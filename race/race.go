package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	num := flag.Int("n", 10, "number")
	flag.Parse()
	number := *num
	k := make(map[int]int)
	var wg sync.WaitGroup
	for i := 0; i <= number; i++ {
		wg.Add(1)
		go func() {
			k[i] = i
			defer wg.Done()
		}()
	}
	wg.Wait()
	k[2] = 22
	fmt.Printf("map value:\t%v\n", k)
}
