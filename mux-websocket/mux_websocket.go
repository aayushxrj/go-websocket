package muxwebsocket

import (
	"context"
	"fmt"
	"log"
	"net/http"
	// "strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
) 


var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,	
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "chrome-extension://cbcbkhdmedgianpaifchdaddpnmgnknn"
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err!= nil{
		w.Write([]byte(err.Error()))
		return 
	}
	defer conn.Close()

	// for{
	// 	conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
   	// 	time.Sleep(time.Second)
	// }

	// i := 0
	// for {
	// 	i++;
	// 	conn.WriteMessage(websocket.TextMessage, []byte("New message (#"+strconv.Itoa(i)+")"))
	// 	time.Sleep(time.Second)
	// }

	for{
		// Read message from client
		_, message, err := conn.ReadMessage()
		if err != nil{
			fmt.Println("Error: ", err)
			break
		}

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

func Mux(){
	router := mux.NewRouter()
	router.HandleFunc("/ws", wsHandler)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	go func(){
		fmt.Println("Server started at localhost:8080/ws")
		log.Fatal(server.ListenAndServe())
	}()

	time.Sleep(60 * time.Second)
	fmt.Println("Server is shutting down due to idle timeout...")
    _ = server.Shutdown(context.Background())
}