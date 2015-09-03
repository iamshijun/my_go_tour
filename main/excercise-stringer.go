package main

import "fmt"

/*
type IpAddr4 struct {
	P1, P2, P3, P4 byte
}

func (addr IpAddr4) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", addr.P1, addr.P2, addr.P3, addr.P4)
}
*/

type IPAddr [4]byte

func (addr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", addr[0], addr[1], addr[2], addr[3])
}

func main() {
	//addr := IpAddr4{192, 168, 1, 2}

	addr := IPAddr([...]byte{192, 168, 1, 3})
	// if we remove method String off IPAddr we got [192 168 1 3] default stringer in []T
	fmt.Println(addr)

	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	//retrieve map  , a do like a Map.Entry
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
