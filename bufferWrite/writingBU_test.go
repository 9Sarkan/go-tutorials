package main

import (
	"fmt"
	"os"
	"testing"
)

var ERR error

func benchmarkCreate(b *testing.B, buffersize, filesize int) {
	dst := "/tmp/go/file"
	var err error
	for i := 0; i < b.N; i++ {
		err = create(dst, buffersize, filesize)
	}
	ERR = err
	err = os.Remove(dst)
	if err != nil {
		fmt.Println(err)
	}
}

func Benchmark32Create(b *testing.B) {
	benchmarkCreate(b, 32, 10000000)
}
func Benchmark64Create(b *testing.B) {
	benchmarkCreate(b, 64, 10000000)
}
func Benchmark128Create(b *testing.B) {
	benchmarkCreate(b, 128, 10000000)
}
func Benchmark512Create(b *testing.B) {
	benchmarkCreate(b, 512, 10000000)
}
func Benchmark10000000Create(b *testing.B) {
	benchmarkCreate(b, 10000000, 10000000)
}
