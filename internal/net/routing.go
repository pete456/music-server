package net

import (
	"music-server/config"
	"fmt"
	"net"
	"bufio"
	"strings"
)

// Handles retrieving data from disk and sends it to user
func HandleRequest(c net.Conn, r []string) {
        switch r[0] {
        case "audio":
                SendData(c,r[1])
        case "thumbnail":
                SendThumbnail(c,r[1])
        case "info":
                SendInfo(c,r[1])
	default:
		fmt.Fprintf(c,config.RESOURCE_NOT_FOUND)
        }
}

// Breaks down the first slot of user request
func ParseUserRequest(c net.Conn) {
        s := bufio.NewScanner(c)
        s.Scan()
        t := s.Text()
        l := strings.Split(t,";")
        header := l[0]
        tail := l[1:]
        switch header {
        case "request":
                HandleRequest(c,tail)
	default:
		fmt.Fprintf(c,config.SERVICE_NOT_FOUND)
        }
        c.Close()
}

