package main

import (
	"fmt"
	"net"
	"os"
)

func lookupIP(address string) ([]string, error) {
	hosts, err := net.LookupAddr(address)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return hosts, nil
}
func lookupHost(hostname string) ([]string, error) {
	address, err := net.LookupHost(hostname)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return address, nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		panic("you have to send a host or ip for lookup")
	}
	ipAddress := net.ParseIP(args[1])
	if ipAddress != nil {
		hosts, err := lookupIP(args[1])
		if err == nil {
			for _, i := range hosts {
				fmt.Println(i)
			}
		}
	} else {
		address, err := lookupHost(args[1])
		if err == nil {
			for _, i := range address {
				fmt.Println(i)
			}
		}
	}
}
