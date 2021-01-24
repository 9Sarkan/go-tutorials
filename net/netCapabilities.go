package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, i := range interfaces {
		fmt.Printf("Interface Name: %v\nInterface MTU: %v\nInterface Hardware Addr: %v\nInterface flags: %v\n", i.Name, i.MTU, i.HardwareAddr.String(), i.Flags.String())
	}
}
