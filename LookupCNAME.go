package main

import (
	"net"
	"os"
	"fmt"
)

func main() {
	
	if len(os.Args) != 2 {

		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]

	addr, err := net.LookupCNAME(name)

	

	if err != nil {

		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	fmt.Println("Canonical address is", addr)

	os.Exit(0)
}