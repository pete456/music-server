package net

import (
	"music-server/internal/song"
	"fmt"
	"os"
)

func SendData(s string) *os.File {
        f := song.OpenAudioData(s)
	return f
}

/*
func SendThumbnail(name string) *os.File {
	return 
}
*/
func encodeInfo(i *song.SongInfo) string {
	return fmt.Sprintf("name=%s;title=%s;artist=%s",i.Name,i.Title,i.Artist)
}

func SendInfo(s string) string {
	si := song.OpenInfo(s)
	d := encodeInfo(si)
	return d
}
