package main

import (
	"fmt"
	"net"

	"sync"

	"github.com/InoFlexin/serverbase/base"
)

type MyMessage base.Message

//MyMessage.OnMesageReceive OSGI Gateway Socket Message Receive Callback
func (m MyMessage) OnMessageReceive(message *base.Message, client net.Conn) {
	fmt.Println("on message receive: "+message.Json+" action: %d", message.Action)
}

func (m MyMessage) OnConnect(message *base.Message, client net.Conn) {
	fmt.Println("on connect: "+message.Json+" action: %d", message.Action)
}

func (m MyMessage) OnClose(err error) {
}

func main() {
	wg := sync.WaitGroup{}

	ev := MyMessage{} // 서버 이벤트 선언
	boot := base.Boot{Protocol: "tcp",
		Port:        ":5092",
		ServerName:  "test_server",
		Callback:    ev,
		ReceiveSize: 1024,
		Complex:     false}
	// server boot option 설정

	wg.Add(1) // synchronized gorutine
	go base.ServerStart(boot, &wg)
	wg.Wait()
}
