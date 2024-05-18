package main

import (
	// ginwebsocket "github.com/aayushxrj/go-gin-gorilla-websocket/gin-websocket"
	ginmuxwebsocket "github.com/aayushxrj/go-gin-gorilla-websocket/mux-websocket"
	// nethttpwebsocket "github.com/aayushxrj/go-gin-gorilla-websocket/net-http-websocket"
)


func main() {
	// ginwebsocket.Gin()
	ginmuxwebsocket.Mux()
	// nethttpwebsocket.NetHttp()
}