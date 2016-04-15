package ws

import "fmt"

type Message struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

func (self *Message) String() string {
	fmt.Println("message received")
	fmt.Println(self.Author)
	fmt.Println(self.Body)
	return self.Author + " says " + self.Body
}
