package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ipLookup bool
	var rdnsLookup bool
	flag.BoolVar(&ipLookup, "ip-lookup", false, "")
	flag.BoolVar(&rdnsLookup, "rdns", false, "")
	flag.Parse()
	for scanner.Scan() {
		if ipLookup {
			data, _ := net.LookupIP(strings.TrimSpace(scanner.Text()))
			for _, d := range data {
				fmt.Println(d)
			}
		}
		if rdnsLookup {
			data, _ := net.LookupAddr(strings.TrimSpace(scanner.Text()))
			for _, d := range data {
				fmt.Println(d)
			}
		}
	}
}
