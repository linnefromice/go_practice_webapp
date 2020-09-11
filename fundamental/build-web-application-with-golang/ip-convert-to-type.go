package main

import (
	"net"
	"os"
	"fmt"
)

// go run ip-convert-to-type.go 127.0.0.1 -> The address is  127.0.0.1
// go run ip-convert-to-type.go xxx.xxx.xxx.xxx -> Invalid address
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}