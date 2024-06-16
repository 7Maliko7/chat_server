package http

import (
	"testing"
	"time"

	"github.com/7Maliko7/chat_server/internal/service/chat"
)

var TransportTest *Transport

func TestTransport(t *testing.T) {
	ChatService := chat.New("token")
	TransportTest = New(":443", ChatService)
}

func TestStart(t *testing.T) {
	go func() {
		err := TransportTest.Start(true, "../../../rootCA.pem", "../../../rootCA.key")
		if err != nil {
			t.Error(err)
		}
	}()
	time.Sleep(3 * time.Second)
}
