package router

import (
	"net"
)

var routers map[string]func(json string) string = make(map[string]func(json string) string)

type RouteHandlerStruct struct {
	Json       string
	Key        string
	Connection net.Conn
}

func Route(path string, fn func(s string) string) {
	routers[path] = fn
}

func RouteHandle(routeInformation *RouteHandlerStruct) string {
	fn := routers[routeInformation.Key]

	if fn != nil {
		return fn(routeInformation.Json)
	}

	return ""
}
