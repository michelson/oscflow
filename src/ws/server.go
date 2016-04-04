package ws

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
	"time"
	//"os"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	fmt.Println("ola")

	//fmt.Println(ws)

	var in []byte
	if err := websocket.Message.Receive(ws, &in); err != nil {
		return
	}
	fmt.Printf("WEbsotcket message Received: %s\n", string(in))
	websocket.Message.Send(ws, in)

	//this might not be necessary
	io.Copy(ws, ws)
}

// This example demonstrates a trivial echo server.
func Listener() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func Writter() {

	origin := "http://localhost/"
	url := "ws://localhost:12345/echo"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])

}

func FakeSender() {
	//senbd websocket messages
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				Writter()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
