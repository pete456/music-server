package net

import (
	"music-server/config"
	"net"
	"log"
	"fmt"
)

func InitServer() {
	l, err := net.Listen("tcp",fmt.Sprintf("%s:%d",config.IP,config.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		ParseUserRequest(conn)
	}
}
