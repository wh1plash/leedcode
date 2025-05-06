package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stringer interface {
	String() string
}

type IPAddr [4]byte

func (i IPAddr) String() string {
	var parts []string
	for _, v := range i {
		parts = append(parts, strconv.Itoa(int(v)))
	}
	return strings.Join(parts, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"google":   {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip.String())

	}
}
