package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ips, _ := net.LookupIP(strings.TrimSpace(scanner.Text()))
		for _, ip := range ips {
			fmt.Println(ip)
		}
	}
}
