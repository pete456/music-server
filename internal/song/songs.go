package song

import (
	"music-server/config"
	"github.com/BurntSushi/toml"
	"os"
	"log"
)

type SongInfo struct {
	Name string
	Title string
	Artist string
	PlayTime int
}

func OpenAudioData(name string) *os.File {
        f, err := os.Open(config.Datapath+name+"/"+name+"."+config.Audioformat)
	if err != nil {
		log.Panic(err)
	}
        return f
}

func OpenInfo(item string) *SongInfo {
	f,err := os.Open(config.Datapath+item+"/index.toml")
	if err != nil {
		log.Panic(err)
	}
        si := SongInfo{Name:item}
        if _, err := toml.DecodeReader(f,&si); err != nil {
                log.Panic(err)
        }
	return &si
}

