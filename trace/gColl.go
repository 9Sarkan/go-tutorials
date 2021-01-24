package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func printMem(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Printf("Alloc:\t%d\n", mem.Alloc)
	fmt.Printf("Total Alloc:\t%d\n", mem.TotalAlloc)
	fmt.Printf("HeapAlloc:\t%d\n", mem.HeapAlloc)
	fmt.Printf("NumGC:\t%d\n", mem.NumGC)
	fmt.Println("--------------")
}

func main() {
	file, err := os.Create("/tmp/goTrace/trace.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = trace.Start(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer trace.Stop()
	var mem runtime.MemStats
	printMem(mem)
	for i := 0; i < 10; i++ {
		s := make([]byte, 1000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(50 * time.Millisecond)
	}
	printMem(mem)
	for i := 0; i < 20; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("failed to created!")
		}
		time.Sleep(50 * time.Millisecond)
	}
	printMem(mem)
}
