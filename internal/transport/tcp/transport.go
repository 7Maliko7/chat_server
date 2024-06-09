package tcp

import (
	"log"
	"net"
	"net/netip"

	"github.com/7Maliko7/chat_server/internal/service/chat"
)

type Transport struct {
	Network string
	Adress  string
	router  *router
}

func New(network, adress string, s *chat.Service) *Transport {
	return &Transport{
		Network: network,
		Adress:  adress,
		router:  newRouter(s),
	}
}

func (t *Transport) Start() {
	go t.router.ListenTcp()

	if t.Network == "udp" {
		addr, err := netip.ParseAddrPort(t.Adress)
		if err != nil {
			log.Fatal(err)
		}

		udpaddr := net.UDPAddrFromAddrPort(addr)
		ln, err := net.ListenUDP(t.Network, udpaddr)
		if err != nil {
			log.Fatal(err)
		}

		for {
			go t.router.Handle(ln)
		}
	} else {
		ln, err := net.Listen(t.Network, t.Adress)
		if err != nil {
			log.Fatal(err)
		}

		for {
			conn, _ := ln.Accept()
			go t.router.Handle(conn)
		}
	}

}
