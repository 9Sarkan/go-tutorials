package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		panic("you have to send a domain too!")
	}
	ns, err := net.LookupNS(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	mx, err := net.LookupMX(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("---------- NS -----------\n")
	for _, i := range ns {
		fmt.Println(i)
	}
	fmt.Printf("---------- MX -----------\n")
	for _, i := range mx {
		fmt.Println(i)
	}

}
