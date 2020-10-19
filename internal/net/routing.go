package net

import (
	"music-server/config"
	"fmt"
	"bufio"
	"strings"
	"io"
)

// Handles retrieving data from disk and sends it to user
func HandleData(s Server,r []string) {
        switch r[0] {
        case "audio":
		d := SendData(r[1])
		io.Copy(s.Conn,d)
        case "thumbnail":
                //io.Copy(c,SendThumbnail(r[1]))
        case "info":
		fmt.Fprintf(s.Conn, SendInfo(r[1]))
	default:
		fmt.Fprintf(s.Conn,config.RESOURCE_NOT_FOUND)
        }
}

func HandleShelf(s Server, r []string) {
	SendShelf(r)
}

// Breaks down the first slot of user request
func ParseUserRequest(s Server) {
        scn := bufio.NewScanner(s.Conn)
        scn.Scan()
        t := scn.Text()
	l := strings.Split(t,":")
        header := l[0]
        tail := l[1:]
        switch header {
        case "data":
                HandleData(s,tail)
	case "shelf":
		HandleShelf(s,tail)
	default:
		fmt.Fprintf(s.Conn,config.SERVICE_NOT_FOUND)
        }
}

