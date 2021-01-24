package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "serving: %s\n", r.URL.Path)
	fmt.Printf("served: %s\n", r.Host)
}
func timeHandler(w http.ResponseWriter, r *http.Request) {
	body := "The current time is : "
	t := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "<h1>%s</h1>\n", body)
	fmt.Fprintf(w, "<h2>%s</h2>\n", t)
	fmt.Fprintf(w, "serving: %s\n", r.URL.Path)
	fmt.Printf("served: %s\n", r.Host)
}

func main() {
	args := os.Args
	port := ":8081"
	if len(args) == 2 {
		portString, err := strconv.Atoi(args[1])
		if err == nil {
			port = fmt.Sprintf(":%d", portString)
		}
	}
	fmt.Printf("start web server in port:\t %s\n", port)

	// config Handlers
	http.HandleFunc("/", myHandler)
	http.HandleFunc("/time", timeHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
