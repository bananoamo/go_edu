package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: ./%s [filename]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	separators := []string{"\t", "*", "|", "•"}
	numberOfReadLines, lines := readFileUpToNLines(os.Args[1], 10)
	separatorsByLines := getSepListFromLines(lines, separators, numberOfReadLines)
	fmt.Println(separatorsByLines)
}

func readFileUpToNLines(fileName string, maxLines int) (int, []string) {
	var file *os.File
	var err error
	if file, err = os.Open(fileName); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	lines := make([]string, 0, maxLines)
	i := 0
	for ; i < maxLines; i++ {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		lines = append(lines, line)
	}
	return i, lines
}
func getSepListFromLines(lines, sep []string, maxLines int) [][]int {
	sepList := make([][]int, len(sep))
	for sepIdx := range sep {
		sepList[sepIdx] = make([]int, maxLines)
		for lineIdx, line := range lines {
			sepList[sepIdx][lineIdx] = strings.Count(line, sep[sepIdx])
		}
	}
	return sepList
}