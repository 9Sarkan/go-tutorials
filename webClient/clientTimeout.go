package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var timeout = time.Duration(time.Second)

// Timeout for client connections
func Timeout(network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}
func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: %s URL Timeout\n", filepath.Base(args[0]))
		return
	}
	tempURL, err := url.Parse(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	timeoutTemp, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	timeout = time.Duration(time.Duration(timeoutTemp) * time.Second)
	t := http.Transport{Dial: Timeout}
	client := http.Client{Transport: &t}
	data, err := client.Get(tempURL.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer data.Body.Close()
	_, err = io.Copy(os.Stdout, data.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}
