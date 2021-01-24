package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var BUFFERSIZE int
var FILESIZE int

func random(a, b int) int {
	return rand.Intn(b-a) + a
}

func createBuffer(buf *[]byte, count int) {
	*buf = make([]byte, count)
	if count == 0 {
		return
	}
	for i := 0; i < count; i++ {
		intByte := byte(random(0, 100))
		if len(*buf) > count {
			return
		}
		*buf = append(*buf, intByte)
	}
}

func create(dst string, b, f int) error {
	_, err := os.Stat(dst)
	if err == nil {
		return fmt.Errorf("file %s is already exists", dst)
	}
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	buf := make([]byte, 0)
	for {
		createBuffer(&buf, b)
		buf = buf[:b]
		if _, err := destination.Write(buf); err != nil {
			return err
		}
		if f < 0 {
			break
		}
		f = f - len(buf)
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) != 3 {
		panic("You have to send a file size and a buffer size to this app")
	}
	BUFFERSIZE, err := strconv.Atoi(args[1])
	FILESIZE, errF := strconv.Atoi(args[2])
	if err != nil || errF != nil {
		panic("invalid data")
	}
	dst := "/tmp/go/tempFile"
	err = create(dst, BUFFERSIZE, FILESIZE)
	if err != nil {
		panic(err)
	}
	err = os.Remove(dst)
	if err != nil {
		panic(err)
	}
}
