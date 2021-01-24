package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func f1(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}
	time.Sleep(time.Microsecond)
	return f1(n-1) + f1(n-2)
}

func f2(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	time.Sleep(time.Microsecond)
	return fn[n]
}

func primeN1(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func primeN2(n int) bool {
	k := math.Floor(float64(n/2) + 1)
	for i := 2; i < int(k); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// create cpu profiling file
	cpuProf, err := os.Create("/tmp/cpuProfiling.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	pprof.StartCPUProfile(cpuProf)
	defer pprof.StopCPUProfile()
	total := 0
	for i := 2; i < 100000; i++ {
		n := primeN1(i)
		if n {
			total = total + 1
		}
	}
	fmt.Printf("Total Prime number N1: %d\n", total)
	// second method
	total = 0
	for i := 2; i < 100000; i++ {
		n := primeN2(i)
		if n {
			total = total + 1
		}
	}
	fmt.Printf("Total prime number with method n2: %d\n", total)

	// fib
	fmt.Println("fib----------------")
	for i := 0; i < 20; i++ {
		f := f1(i)
		fmt.Print(f, " ")
	}
	fmt.Println()
	for i := 0; i < 20; i++ {
		f := f2(i)
		fmt.Print(f, " ")
	}
	fmt.Println()
	runtime.GC()
	// memory profiling
	memory, err := os.Create("/tmp/memoryProfiling.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer memory.Close()
	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("%d created\t", i)
	}
	err = pprof.WriteHeapProfile(memory)
	if err != nil {
		fmt.Println(err)
		return
	}
}
