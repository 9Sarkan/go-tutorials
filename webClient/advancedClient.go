package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}
	URL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpData, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("status code: ", httpData.StatusCode)
	headers, _ := httputil.DumpResponse(httpData, false)
	fmt.Println(string(headers))
	contentType := httpData.Header.Get("Content-Type")
	fmt.Println("ContentType : ", contentType)
	charset := strings.SplitAfter(contentType, "charset=")
	if len(charset) > 1 {
		fmt.Println("Charset: ", charset[0])
	}
	if httpData.ContentLength == -1 {
		fmt.Println("Unknown content type!")
	} else {
		fmt.Println("ContentLength: ", httpData.ContentLength)
	}

	r := httpData.Body
	defer httpData.Body.Close()
	length := 0
	var buffer [1024]byte
	for {
		n, err := r.Read(buffer[0:])
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}
	fmt.Printf("Calculated body data length: \t%d\n", length)
}
