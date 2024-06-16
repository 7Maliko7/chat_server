package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/7Maliko7/chat_server/internal/config"
	"github.com/7Maliko7/chat_server/internal/service/chat"
	"github.com/7Maliko7/chat_server/internal/transport/http"
	"github.com/7Maliko7/chat_server/internal/transport/tcp"
)

func main() {
	cfg, err := config.New(".env")
	if err != nil {
		log.Fatal(err)
	}

	chatService := chat.New(cfg.AuthToken)

	wg := sync.WaitGroup{}

	if cfg.UseHttp {
		transport := http.New(cfg.HttpPort, chatService)
		wg.Add(1)
		go func() {
			err := transport.Start(cfg.EnableTLS, cfg.CertFile, cfg.KeyFile)
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}()
	}

	if cfg.UseTcp {
		transportTCP := tcp.New("tcp", cfg.TcpPort, chatService)
		wg.Add(1)
		go func() {
			transportTCP.Start()
			wg.Done()
		}()
	}

	fmt.Println("Start server...")
	wg.Wait()

}
