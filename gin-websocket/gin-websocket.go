package ginwebsocket

import (
	"fmt"
	"net/http"
	// "time"

	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
) 

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,	
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
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
	// i := 0
	// for {
	// 	i++;
	// 	conn.WriteMessage(websocket.TextMessage, []byte("New message (#"+strconv.Itoa(i)+")"))
	// 	time.Sleep(time.Second)
	// }


	// Reading message from client and sending response to client
	for{
		// Read message from client
		_, message, err := conn.ReadMessage()
		if err != nil{
			fmt.Println("Error: ", err)
			break
		}

		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(message))
		fmt.Printf("\nReceived: %s\n", message)

		// Send response to client
		if string(message) == "ping"{
			message= []byte("pong")
		}
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil{
			fmt.Println("Error: ", err)
			break
		}

		fmt.Printf("Sent: %s\n", message)
	}
}

func Gin() {
	router := gin.Default()

	router.GET("/ws", wsHandler)

	router.Run("localhost:8080")
}