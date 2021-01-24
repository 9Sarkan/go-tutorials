package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	delay := random(0, 5)
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Fprintf(w, "delay:\t%d\n", delay)
	fmt.Fprintf(w, "serving:\t%s\n", r.URL.Path)
	fmt.Printf("served:\t%s\n", r.Host)
}

func main() {
	PORT := ":8081"
	if len(os.Args) == 1 {
		fmt.Println("going to use default port, ", PORT)
	} else {
		PORT = ":" + os.Args[1]
	}
	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
