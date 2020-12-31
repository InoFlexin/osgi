package main

import (
	"net"
	"osgi"

	"sync"

	"github.com/InoFlexin/serverbase/base"
)

type MyMessage base.Message

//MyMessage.OnMesageReceive OSGI Gateway Socket Message Receive Callback
func (m MyMessage) OnMessageReceive(message *base.Message, client net.Conn) {
	// "{"ProxySequence: 1, Action: 1, Data: "data""}"
	osgiPacket := osgi.ToPacket(message.Json)
	//Router로 패킷을 넘겨줌.
}

func (m MyMessage) OnConnect(message *base.Message, client net.Conn) {
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
