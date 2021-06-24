package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

func (i IPAddr) String() string {
	var ret string
	for c, v := range i {
		ret += strconv.Itoa(int(v))
		if l := len(i)-1; c<l {ret += "."}
	} 
	return ret
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
