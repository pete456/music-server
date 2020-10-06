#!/bin/bash
root=~/programming/go/music-server
src=$root/cmd/music-server/music-server.go
output=$root/build/music-server
go build -o $output $src
