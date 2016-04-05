package main

import (
	"fmt"
	"github.com/hypebeast/go-osc/osc"
	//gui "gui"
	//web "web"
	ws "ws"
)

// osc listener
// websocket listener
func main() {
	fmt.Println("init oscflow")

	addr := "127.0.0.1:8765"
	server := &osc.Server{Addr: addr}

	server.Handle("/message/address", func(msg *osc.Message) {
		//osc.PrintMessage(msg)
		fmt.Println(msg)
		var m = &ws.Message{
			Author: "miguel",
			Body:   "this is from osc mdf!"}

		ws.Sender(m)
	})

	go ws.Listener()

	ws.FakeSender()

	//go web.Example()

	//go gui.Example()

	//if err := qml.Run(run); err != nil {
	//	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	//	os.Exit(1)
	//}

	server.ListenAndServe()
}

/*
// osc client
func client() {
	client := osc.NewClient("localhost", 8765)
	msg := osc.NewMessage("/osc/address")
	msg.Append(int32(111))
	msg.Append(true)
	msg.Append("hello")
	client.Send(msg)
}
*/
