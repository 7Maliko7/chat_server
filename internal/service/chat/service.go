package chat

import (
	"fmt"

	"github.com/7Maliko7/chat_server/internal/model/client"
	"github.com/7Maliko7/chat_server/pkg/errors"
)

type Service struct {
	Conns     map[string]*client.Model
	AuthToken string
}

func New(AuthToken string) *Service {
	return &Service{Conns: make(map[string]*client.Model), AuthToken: AuthToken}
}

func (c *Service) ConnectChat(cl *client.Model, AuthToken string) error {
	if c.AuthToken != AuthToken {
		fmt.Printf("Invalid AuthToken")
		return errors.ErrInvalidAuthToken
	}
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
