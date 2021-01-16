package main

import (
	"fmt"
	"time"
)

func f1(a, b chan struct{}) {
	<-a
	time.Sleep(time.Second)
	fmt.Println("F1();")
	close(b)
}
func f2(a, b chan struct{}) {
	<-a
	fmt.Println("F2();")
	close(b)
}
func f3(a <-chan struct{}) {
	<-a
	fmt.Println("F3!")
}
func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go f3(z)
	go f2(y, x)
	go f3(z)
	go f1(x, z)

	close(y)

	time.Sleep(time.Second * 3)
}
