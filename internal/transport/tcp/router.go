package tcp

import (
	"bufio"
	"encoding/json"
	"log"
	"net"

	"github.com/7Maliko7/chat_server/internal/model/client"
	"github.com/7Maliko7/chat_server/internal/model/message"
	"github.com/7Maliko7/chat_server/internal/service/chat"
)

type router struct {
	Service *chat.Service
	Ch      chan message.Model
}

func newRouter(s *chat.Service) *router {
	return &router{
		Service: s,
		Ch:      make(chan message.Model),
	}
}

func (r *router) Handle(conn net.Conn) error {
	for {
		buf, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			if err.Error() == "EOF" {
				log.Print("connection lost")
				break
			}
			return err
		}

		var msg ConnectRequest
		err = json.Unmarshal(buf, &msg)
		if err != nil {
			return err
		}

		err = r.Service.ConnectChat(client.New(conn, r.Ch, msg.Name, "tcp"))
		if err != nil {
			return err
		}

		break
	}
	return nil
}

func (r *router) ListenTcp() error {
	for msg := range r.Ch {
		var msgBody MessageRequest
		err := json.Unmarshal(msg.Message, &msgBody)
		if err != nil {
			return err
		}

		if msgBody.Broadcast {
			r.Service.Broadcast(msg.From, msgBody.Message)
		} else {
			r.Service.Send(msg.From, msgBody.To, msgBody.Message)
		}
	}
	return nil
}
