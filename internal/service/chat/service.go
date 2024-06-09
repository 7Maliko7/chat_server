package chat

import (
	"fmt"

	"github.com/7Maliko7/chat_server/internal/model/client"
)

type Service struct {
	Conns map[string]*client.Model
}

func New() *Service {
	return &Service{Conns: make(map[string]*client.Model)}
}

func (c *Service) ConnectChat(cl *client.Model) error {
	c.Conns[cl.Name] = cl
	fmt.Printf("%s connected\n", cl.Name)
	go cl.Listen()
	return nil
}

func (c *Service) Send(from, name, text string) error {
	conn := c.Conns[name]
	if conn == nil {
		fmt.Printf("%s not found\n", name)
		return nil
	}
	err := conn.Send(client.Message{From: from, Text: text})
	if err != nil {
		return err
	}

	return nil
}

func (c *Service) Broadcast(from, text string) error {
	for _, conn := range c.Conns {
		err := conn.Send(client.Message{From: from, Text: text})
		if err != nil {
			return err
		}
	}

	return nil
}
