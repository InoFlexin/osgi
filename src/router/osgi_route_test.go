package router

import (
	"net"
	"sync"
	"testing"

	"github.com/InoFlexin/serverbase/auth"
	"github.com/InoFlexin/serverbase/base"
	"github.com/InoFlexin/serverbase/client"
)

type MyClientMessage base.Message //Client Message 타입 정의
type MyMessage base.Message

var testWg sync.WaitGroup = sync.WaitGroup{}
var testResult string

func (m MyMessage) OnMessageReceive(message *base.Message, client net.Conn) {
	testResult = RouteHandle(message.Key, message.Json)
	testWg.Done()
}

func (m MyMessage) OnConnect(message *base.Message, client net.Conn) {
}

func (m MyMessage) OnClose(err error) {
}

func (m MyClientMessage) OnMessageReceive(message *base.Message, client net.Conn) {
}

func (m MyClientMessage) OnConnect(message *base.Message, client net.Conn) {
}

func (m MyClientMessage) OnClose(err error) {
}

func _initializeRouters() {
	clientHandler := func(s string) string {
		return "pong"
	}

	Route(auth.GetKey("client"), clientHandler)
}

func TestClientRun(t *testing.T) {
	wg := sync.WaitGroup{}

	auth.RegisterKey("server", "osgi_server@"+auth.GenerateKey(20))
	auth.RegisterKey("client", "osgi_test_client@"+auth.GenerateKey(20))
	_initializeRouters()

	ev := MyMessage{} // 서버 이벤트 선언
	boot := base.Boot{Protocol: "tcp",
		Port:        ":5092",
		ServerName:  "test_server",
		Callback:    ev,
		ReceiveSize: 1024,
		Complex:     true}
	// server boot option 설정

	wg.Add(1) // synchronized gorutine
	go base.ServerStart(boot, &wg)
	wg.Wait()

	event := MyClientMessage{}
	clientBoot := client.ClientBoot{Protocol: "tcp",
		HostAddr:   "localhost",
		HostPort:   ":5092",
		Callback:   event,
		BufferSize: 1024}

	wg.Add(1) // synchronized goroutine
	go client.ConnectServer(&clientBoot, &wg)
	wg.Wait()

	testWg.Add(1)
	go client.Write("ping")
	testWg.Wait()

	if testResult != "pong" {
		t.Error("don't received 'pong' message")
	}
}
