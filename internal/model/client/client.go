package client

import (
	"bufio"
	"bytes"
	"log"
	"net"

	"github.com/7Maliko7/chat_server/internal/model/message"
	"github.com/7Maliko7/chat_server/pkg/errors"
	"github.com/gorilla/websocket"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Model struct {
	conn     interface{}
	Name     string
	Ch       chan message.Model
	connType string
}

type Message struct {
	From string `json:"from"`
	Text string `json:"text"`
}

func New(c interface{}, ch chan message.Model, n, t string) *Model {
	return &Model{
		conn:     c,
		Name:     n,
		Ch:       ch,
		connType: t,
	}
}

func (m *Model) Listen() error {
	switch m.connType {
	case "websocket":
		return m.listenWebsocket()
	case "tcp":
		return m.ListenTcp()
	}
	return nil
}

func (m *Model) listenWebsocket() error {
	c, ok := m.conn.(*websocket.Conn)
	if !ok {
		return errors.ErrUnexpectedConnecion
	}
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		if msg != nil {

			msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))

			m.Ch <- message.Model{From: m.Name, Message: msg}
		}
	}

	return nil
}

func (m *Model) ListenTcp() error {
	c, ok := m.conn.(net.Conn)
	if !ok {
		return errors.ErrUnexpectedConnecion
	}

	for {
		buf, err := bufio.NewReader(c).ReadBytes('\n')
		if err != nil {
			break
		}

		if buf != nil {
			buf = bytes.TrimSpace(bytes.Replace(buf, newline, space, -1))
			m.Ch <- message.Model{From: m.Name, Message: buf}
		}
	}

	return nil
}

func (m *Model) Send(message Message) error {
	switch m.connType {
	case "websocket":
		c, ok := m.conn.(*websocket.Conn)
		if ok {
			err := c.WriteJSON(message)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
