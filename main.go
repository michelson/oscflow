package main

import "github.com/hypebeast/go-osc/osc"

func main() {
	addr := "127.0.0.1:8765"
	server := &osc.Server{Addr: addr}

	server.Handle("/message/address", func(msg *osc.Message) {
		osc.PrintMessage(msg)
	})

	server.ListenAndServe()
}

func client() {
	client := osc.NewClient("localhost", 8766)
	msg := osc.NewMessage("/osc/address")
	msg.Append(int32(111))
	msg.Append(true)
	msg.Append("hello")
	client.Send(msg)
}
