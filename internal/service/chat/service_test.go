package chat

import (
	"net"
	"testing"

	"github.com/7Maliko7/chat_server/internal/model/client"
	"github.com/7Maliko7/chat_server/internal/model/message"
)

var ChatService *Service

func TestService(t *testing.T) {
	ChatService = New()
}

func TestConnectChat(t *testing.T) {
	conn := net.TCPConn{}
	ch := make(chan message.Model)
	cl := client.New(conn, ch, "John", "tcp")

	err := ChatService.ConnectChat(cl)
	if err != nil {
		t.Error(err)
	}
}
func TestSend(t *testing.T) {
	err := ChatService.Send("Ann", "John", "Hi")
	if err != nil {
		t.Error(err)
	}
}

func TestSendNoClient(t *testing.T) {
	err := ChatService.Send("Ann", "Julie", "Hi")
	if err != nil {
		t.Error(err)
	}
}

func TestBroadcast(t *testing.T) {
	err := ChatService.Broadcast("Ann", "Hi")
	if err != nil {
		t.Error(err)
	}
}
