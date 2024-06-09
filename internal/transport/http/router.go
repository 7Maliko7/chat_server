package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/7Maliko7/chat_server/internal/model/client"
	"github.com/7Maliko7/chat_server/internal/model/message"
	"github.com/7Maliko7/chat_server/internal/service/chat"
)

type router struct {
	Service  *chat.Service
	Ch       chan message.Model
	upgrader websocket.Upgrader
}

func newRouter(s *chat.Service) *router {
	return &router{
		Service: s,
		Ch:      make(chan message.Model),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (r *router) ListenWebsocket() error {
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

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		u := req.URL.Path
		switch u {
		case "/connect":
			conn, err := r.upgrader.Upgrade(w, req, nil)
			if err != nil {
				log.Println(err)
				return
			}

			clientName := req.URL.Query().Get("name")

			err = r.Service.ConnectChat(client.New(conn, r.Ch, clientName, "websocket"))
			if err != nil {
				fmt.Fprint(w, http.Response{
					StatusCode: http.StatusInternalServerError,
				})
			}
			fmt.Fprint(w, http.Response{
				StatusCode: http.StatusOK,
			})

		}
		fmt.Fprint(w, http.Response{StatusCode: http.StatusMethodNotAllowed})
	}
}
