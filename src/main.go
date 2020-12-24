package main

import (
	"fmt"

	"github.com/InoFlexin/serverbase/base"
)

func main() {
	boot := base.Boot{Protocol: "tcp",
		Port:        ":5092",
		ServerName:  "test_server",
		Callback:    nil,
		ReceiveSize: 1024}

	fmt.Println(boot)
}
