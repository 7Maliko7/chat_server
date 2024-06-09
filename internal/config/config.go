package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	HttpPort  string `env:"HTTP_PORT"`
	TcpPort   string `env:"TCP_PORT"`
	UdpPort   string `env:"UDP_PORT"`
	UseHttp   bool   `env:"USE_HTTP"`
	UseTcp    bool   `env:"USE_TCP"`
	UseUdp    bool   `env:"USE_UDP"`
	EnableTLS bool   `env:"ENABLE_TLS"`
	CertFile  string `env:"CERT_FILE"`
	KeyFile   string `env:"KEY_FILE"`
}

func New(path string) (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	c := Config{}
	if err := envconfig.Process(context.Background(), &c); err != nil {
		return nil, err
	}

	return &c, nil
}
