package main

import (
	"log"
	"net"
	"sync"

	"github.com/InoFlexin/osgi/router"
	"github.com/InoFlexin/serverbase/auth"
	"github.com/InoFlexin/serverbase/base"
)

type MyMessage base.Message

//MyMessage.OnMesageReceive OSGI Gateway Socket Message Receive Callback
func (m MyMessage) OnMessageReceive(message *base.Message, client net.Conn) {
	// "{"ProxySequence: 1, Action: 1, Data: "data""}"
	//Router로 패킷을 넘겨줌.
	router.RouteHandle(&router.RouteHandlerStruct{
		Key:        message.Key,
		Json:       message.Json,
		Connection: client,
	})
}

func (m MyMessage) OnConnect(message *base.Message, client net.Conn) {
	log.Println("Successfully client connected " + message.Key)
}

func (m MyMessage) OnClose(err error) {
}

func initializeRouters() {
	clientHandler := func(s string) string {
		return "pong"
	}

	router.Route("test_proxy", clientHandler)
}

func main() {
	wg := sync.WaitGroup{}

	initializeRouters()

	auth.RegisterKey("server", "osgi_server@"+auth.GenerateKey(20))

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
