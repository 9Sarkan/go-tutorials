package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync/atomic"
)

var count int32

func addHandler(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt32(&count, 1)
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	temp := atomic.LoadInt32(&count)
	fmt.Println("count: ", temp)
	fmt.Fprintf(w, "<h1 align=\"center\">%d</h1>\n", count)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	http.HandleFunc("/", addHandler)
	http.HandleFunc("/counter", getCounter)
	http.ListenAndServe(":8080", nil)
}
