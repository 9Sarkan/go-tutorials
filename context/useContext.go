package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	myURL string
	delay int = 5
	w     sync.WaitGroup
)

type myData struct {
	response *http.Response
	err      error
}

func connect(c context.Context) error {
	defer w.Done()
	data := make(chan myData, 1)
	tr := &http.Transport{}
	httpClient := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", myURL, nil)

	go func() {
		response, err := httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
			data <- myData{nil, err}
			return
		}
		pack := myData{response, nil}
		data <- pack
	}()
	select {
	case <-c.Done():
		tr.CancelRequest(req)
		<-data
		fmt.Println("Cancel the request...")
		return c.Err()
	case ok := <-data:
		err := ok.err
		response := ok.response
		if err != nil {
			fmt.Println("error from response: ", err)
			return err
		}
		defer response.Body.Close()
		realHTTPData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		fmt.Println("Response: ", realHTTPData)
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("you have to send a url with a delay time!")
		return
	}
	myURL := args[1]
	delay, _ := strconv.Atoi(args[2])
	fmt.Println("delay: ", delay)
	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Duration(delay)*time.Second)
	defer cancel()
	fmt.Printf("going to connect: %s\n", myURL)
	w.Add(1)
	go connect(c)
	w.Wait()
	fmt.Println("Exiting...")
}
