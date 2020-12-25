package main

import (
	"fmt"
	"osgi"
)

func main() {
	p := osgi.CreatePacket(1, "strings")
	p2 := osgi.ToString(*p)

	fmt.Println(p)
	fmt.Println(p2)
	fmt.Println(osgi.ToPacket(p2))
}
