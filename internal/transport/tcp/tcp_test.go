package tcp

import (
	"testing"

	"github.com/7Maliko7/chat_server/internal/service/chat"
)

var TransportTest *Transport

func TestTransport(t *testing.T) {
	ChatService := chat.New()
	TransportTest = New("tcp", "8081", ChatService)
}
