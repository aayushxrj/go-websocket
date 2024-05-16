package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
) 

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,	
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Here we just allow the chrome extension client accessable (you should check this verify according to your client source)
		return r.Header.Get("Origin") == "chrome-extension://cbcbkhdmedgianpaifchdaddpnmgnknn"
	},
}

func wsHandler(ctx *gin.Context){
	// ctx.IndentedJSON(200, gin.H{
	// 	"message": "Hello, World!",
	// })
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err!= nil{
		ctx.JSON(500, gin.H{"error": err.Error()})
		return 
	}
	defer conn.Close()

	// for{
	// 	conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
   	// 	time.Sleep(time.Second)
	// }

	// sending notification to client
	i := 0
	for {
		i++;
		conn.WriteMessage(websocket.TextMessage, []byte("New message (#"+strconv.Itoa(i)+")"))
		time.Sleep(time.Second)
	}
}

func main() {
	router := gin.Default()

	router.GET("/ws", wsHandler)

	router.Run("localhost:8080")
}