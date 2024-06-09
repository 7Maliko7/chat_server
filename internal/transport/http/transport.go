package http

import (
	"net/http"
	"time"

	"github.com/7Maliko7/chat_server/internal/service/chat"
)

type Transport struct {
	Server *http.Server
	router *router
}

func New(adress string, s *chat.Service) *Transport {
	r := newRouter(s)
	return &Transport{
		Server: &http.Server{
			Addr:           adress,
			Handler:        r,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		router: r,
	}
}

func (t *Transport) Start() error {
	go t.router.ListenWebsocket()
	return t.Server.ListenAndServe()
}
