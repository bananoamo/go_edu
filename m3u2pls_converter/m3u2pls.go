package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Song struct {
	Name   string
	Length int
	Path   string
}

func hasSuffix(s string, extension *string) bool {
	if strings.HasSuffix(s, ".m3u") {
		*extension = "m3u"
		return true
	} else if strings.HasSuffix(s, ".pls") {
		*extension = "pls"
		return true
	} else {
		return false
	}
}

func getSongInfo(s string) (seconds int, name string) {
	const separator = ","
	if i := strings.IndexAny(s, "-0123456789"); i > -1 {
		s = s[i:]
		if j := strings.IndexAny(s, separator); j > -1 {
			name = s[j+len(separator):]
			var err error
			if seconds, err = strconv.Atoi(s[:j]); err != nil {
				log.Printf("failed to read duration of song '%s' : %v\n", name, err)
				seconds = -1
			}
		}
	}
	return seconds, name
}

func mapDirSeparator(c rune) rune {
	if c == '/' || c == '\\' {
		return filepath.Separator
	}
	return c
}

func parseM3uFile(s string) (songs []Song) {
	var song Song

	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#EXTM3U") || line == "" {
			continue
		}
		if strings.HasPrefix(line, "#EXTINF:") {
			song.Length, song.Name = getSongInfo(line)
		} else {
			song.Path = strings.Map(mapDirSeparator, line)
		}
		if song.Name != "" && song.Path != "" && song.Length != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}

func writeM3uToFile(songs []Song) {
	fmt.Println("[playlist]")
	for i, song := range songs {
		i++
		fmt.Printf("File%d=%s\n", i, song.Path)
		fmt.Printf("Title%d=%s\n", i, song.Name)
		fmt.Printf("Length%d=%d\n", i, song.Length)
	}
	fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
}

func hasPrefix(s string, variants ...string) bool {
	for _, variant := range variants {
		if strings.HasPrefix(s, variant) {
			return true
		}
	}
	return false
}

func parsePlsFile(s string) (songs []Song) {
	var song Song
	var numberOfSong int = 1
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || hasPrefix(line, "[playlist]", "Version=", "NumberOfEntries=") {
			continue
		}

		var separator = "="
		var numberOfSongString = strconv.Itoa(numberOfSong)
		if strings.HasPrefix(line, "File"+string(numberOfSongString)+separator) {
			i := strings.IndexAny(line, separator)
			line = line[i+len(separator):]
			song.Path = strings.Map(mapDirSeparator, line)
		}
		if strings.HasPrefix(line, "Title"+string(numberOfSongString)+separator) {
			i := strings.IndexAny(line, separator)
			song.Name = line[i+len(separator):]
		}
		if strings.HasPrefix(line, "Length"+string(numberOfSongString)+separator) {
			i := strings.IndexAny(line, separator)
			line := line[i+len(separator):]
			if seconds, err := strconv.Atoi(line); err != nil {
				log.Printf("failed to read duration of song '%s' : %v\n", song.Name, err)
				seconds = -1
				song.Length = seconds
			} else {
				song.Length = seconds
			}
		}
		if song.Name != "" && song.Path != "" && song.Length != 0 {
			songs = append(songs, song)
			song = Song{}
			numberOfSong++
		}
	}
	return songs
}

func writePlsToFile(songs []Song) {
	fmt.Println("#EXTM3U")
	for _, song := range songs {
		fmt.Printf("#EXTINF:%d,%s\n", song.Length, song.Name)
		fmt.Printf("%s\n", song.Path)
	}
}

func main() {
	fileExtension := ""

	if len(os.Args) == 1 || !hasSuffix(os.Args[1], &fileExtension) {
		fmt.Printf("usage: %s playlist <*.m3u|*.pls>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		if err := strings.Compare(fileExtension, "m3u"); err == 0 {
			m3uToFile := parseM3uFile(string(rawBytes))
			writeM3uToFile(m3uToFile)
		} else {
			plsToFile := parsePlsFile(string(rawBytes))
			writePlsToFile(plsToFile)
		}
	}
}
