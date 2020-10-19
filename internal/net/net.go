package net

import (
	"music-server/config"
	"net"
	"log"
	"fmt"
)

type Server struct {
	Conn net.Conn
}

func InitServer() {
	l, err := net.Listen("tcp",fmt.Sprintf("%s:%d",config.IP,config.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	server := Server{}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		server.Conn=conn
		ParseUserRequest(server)
	}
}
