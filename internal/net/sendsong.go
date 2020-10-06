package net

import (
	"music-server/config"
	"music-server/internal/song"
	"net"
	"fmt"
	"io"
	"log"
)

func sendNoResourceError(c net.Conn) {
	if r := recover(); r != nil {
		log.Printf("Error finding resources\n")
		fmt.Fprintf(c,config.RESOURCE_NOT_FOUND)
	}
}

func SendData(c net.Conn, name string) {
	defer sendNoResourceError(c)
        f := song.OpenAudioData(name)
        io.Copy(c,f)
}

func SendThumbnail(c net.Conn, name string) {

}

func encodeInfo(i *song.SongInfo) string {
	return fmt.Sprintf("name=%s;title=%s;artist=%s",i.Name,i.Title,i.Artist)
}

func SendInfo(c net.Conn, item string) {
	defer sendNoResourceError(c)
	si := song.OpenInfo(item)
	d := encodeInfo(si)
	fmt.Fprintf(c,d)
}
