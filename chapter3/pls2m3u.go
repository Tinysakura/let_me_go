package chapter3

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Song struct {
	Title string
	File string
	Seconds int
}

func WriteM3uPlaylist(songs []Song) {
	fmt.Println("EXTM3U")
	for _, song := range songs {
		buffer := bytes.NewBufferString("")
		buffer.WriteString("EXTINF:")
		buffer.WriteString(strconv.Itoa(song.Seconds))
		buffer.WriteString(",")
		buffer.WriteString(song.Title)
		buffer.WriteString("\r\n")
		buffer.WriteString(song.File)

		fmt.Println(buffer.String())
	}
}

func ReadPlsPlaylist(filepath string)(songs []Song) {
	bytes, e := ioutil.ReadFile(filepath)
	var fileContent string

	if e == nil {
		fileContent = string(bytes)
	}

	var song Song

	split := strings.Split(fileContent, "\n")
	for _, line := range split {
		line = strings.TrimSpace(line)

		if line == "[playlist]" {
			 continue
		}

		if strings.HasPrefix(line, "File") {
			song.File = strings.Split(line, "=")[1]
		}

		if strings.HasPrefix(line, "Title") {
			song.Title = strings.Split(line, "=")[1]
		}

		if  strings.HasPrefix(line, "Length") {
			seconds := strings.Split(line, "=")[1]
			i, e := strconv.Atoi(seconds)
			if e == nil {
				song.Seconds = i
			}

			songs = append(songs, song)
			song = Song{}
		}
	}

	return songs
}
