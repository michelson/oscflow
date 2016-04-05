package ws

import (
	"golang.org/x/net/websocket"
	//"html/template"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	listenAddr = "localhost:8081" // server address
)

var (
	pwd, _ = os.Getwd()
	//RootTemp      = template.Must(template.ParseFiles(pwd + "/chat.html"))
	JSON          = websocket.JSON           // codec for JSON
	Message       = websocket.Message        // codec for string, []byte
	ActiveClients = make(map[ClientConn]int) // map containing clients
)

// Initialize handlers and websocket handlers
func init() {
	//http.HandleFunc("/", RootHandler)
	http.Handle("/echo", websocket.Handler(SockServer))
}

// Client connection consists of the websocket and the client ip
type ClientConn struct {
	websocket *websocket.Conn
	clientIP  string
}

// WebSocket server to handle chat between clients
func SockServer(ws *websocket.Conn) {
	var err error
	var clientMessage string
	// use []byte if websocket binary type is blob or arraybuffer
	// var clientMessage []byte

	// cleanup on server side
	defer func() {
		if err = ws.Close(); err != nil {
			log.Println("Websocket could not be closed", err.Error())
		}
	}()

	client := ws.Request().RemoteAddr
	log.Println("Client connected:", client)
	sockCli := ClientConn{ws, client}
	ActiveClients[sockCli] = 0
	log.Println("Number of clients connected ...", len(ActiveClients))

	// for loop so the websocket stays open otherwise
	// it'll close after one Receieve and Send
	for {
		if err = Message.Receive(ws, &clientMessage); err != nil {
			// If we cannot Read then the connection is closed
			log.Println("Websocket Disconnected waiting", err.Error())
			// remove the ws client conn from our active clients
			delete(ActiveClients, sockCli)
			log.Println("Number of clients still connected ...", len(ActiveClients))
			return
		}

		clientMessage = sockCli.clientIP + " Said: " + clientMessage
		for cs, _ := range ActiveClients {
			if err = Message.Send(cs.websocket, clientMessage); err != nil {
				// we could not send the message to a peer
				log.Println("Could not send message to ", cs.clientIP, err.Error())
			}
		}
	}
}

func Listener() {
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func FakeSender() {
	//senbd websocket messages
	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan struct{})
	//var msg = &Message{Author: "miguel", Body: "Hello is there anybody in there?"}
	go func() {
		for {
			select {
			case <-ticker.C:
				//server.SendAll() <- msg
				SendAll()
				//Writter()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func SendAll() {
	clientMessage := "Chupli!!"
	for cs, _ := range ActiveClients {
		if err := Message.Send(cs.websocket, clientMessage); err != nil {
			// we could not send the message to a peer
			log.Println("Could not send message to ", cs.clientIP, err.Error())
		}
	}
}

type JMessage struct {
	Author string        `json:"author"`
	Body   []interface{} `json:"body"`
}

func (self *JMessage) String() string {
	fmt.Println("message received")
	fmt.Println(self.Author)
	fmt.Println(self.Body)
	return self.Author + " says "
}

func Sender(msg []interface{}) {

	m := &JMessage{
		Author: "ds",
		Body:   msg,
	}

	fmt.Println("msg[0]:")
	fmt.Println(msg[0])

	for cs, _ := range ActiveClients {
		if err := JSON.Send(cs.websocket, m); err != nil {
			// we could not send the message to a peer
			log.Println("Could not send message to ", cs.clientIP, err.Error())
		}
	}
}

/*
import (
	//"fmt"
	"golang.org/x/net/websocket"
	//"io"
	"log"
	"net/http"
	"time"
	//"os"
)

// Chat server.
type Server struct {
	path         string
	clients      []*Client
	addClient    chan *Client
	removeClient chan *Client
	sendAll      chan *Message
	messages     []*Message
}

// Create new chat server.
func NewServer(path string) *Server {
	clients := make([]*Client, 0)
	addClient := make(chan *Client)
	removeClient := make(chan *Client)
	sendAll := make(chan *Message)
	messages := make([]*Message, 0)
	return &Server{path, clients, addClient, removeClient, sendAll, messages}
}

func (self *Server) AddClient() chan<- *Client {
	return (chan<- *Client)(self.addClient)
}

func (self *Server) RemoveClient() chan<- *Client {
	return (chan<- *Client)(self.removeClient)
}

func (self *Server) SendAll() chan<- *Message {
	return (chan<- *Message)(self.sendAll)
}

func (self *Server) Messages() []*Message {
	msgs := make([]*Message, len(self.messages))
	copy(msgs, self.messages)
	return msgs
}

// Listen and serve.
// It serves client connection and broadcast request.
func (self *Server) Listen() {

	log.Println("Listening server...")

	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		client := NewClient(ws, self)
		self.addClient <- client
		client.Listen()
		defer ws.Close()
	}
	http.Handle(self.path, websocket.Handler(onConnected))
	log.Println("Created handler")

	for {
		select {

		// Add new a client
		case c := <-self.addClient:
			log.Println("Added new client")
			self.clients = append(self.clients, c)
			for _, msg := range self.messages {
				c.Write() <- msg
			}
			log.Println("Now", len(self.clients), "clients connected.")

		// remove a client
		case c := <-self.removeClient:
			log.Println("Remove client")
			for i := range self.clients {
				if self.clients[i] == c {
					self.clients = append(self.clients[:i], self.clients[i+1:]...)
					break
				}
			}

		// broadcast message for all clients
		case msg := <-self.sendAll:
			log.Println("Send all:", msg)
			self.messages = append(self.messages, msg)
			for _, c := range self.clients {
				c.Write() <- msg
			}
		}
	}
}

// This example demonstrates a trivial echo server.
var server = NewServer("/echo")

func Sender(msg *Message) {

	server.SendAll() <- msg
}

func Listener() {

	//server := NewServer("/echo")
	go server.Listen()

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}

func FakeSender() {
	//senbd websocket messages
	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan struct{})
	var msg = &Message{Author: "miguel", Body: "Hello is there anybody in there?"}
	go func() {
		for {
			select {
			case <-ticker.C:
				server.SendAll() <- msg
				//Writter()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}*/
