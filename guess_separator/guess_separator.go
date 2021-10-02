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
	numberOfReadLines, lines := readFileUpToNLines(os.Args[1], 5)
	separatorsByLines := getSepListFromLines(lines, separators, numberOfReadLines)
	lineSeparator := guessSep(separators, separatorsByLines, numberOfReadLines)
	printResult(lineSeparator)
}

func readFileUpToNLines(fileName string, maxLines int) (int, []string) {
	var file *os.File
	var err error
	if file, err = os.Open(fileName); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	lines := make([]string, maxLines)
	i := 0
	for ; i < maxLines; i++ {
		line, err := rd.ReadString('\n')
		if line != "" {
			lines[i] = line
		}
		if err != nil {
			if err == io.EOF {
				break
			}
		}
	}
	return i, lines[:i]
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
func guessSep(sepList []string, sepByLines [][]int, maxLine int) string {
	for sepIdx := range sepList {
		same := true
		target := sepByLines[sepIdx][0]
		for lineIdx := 1; lineIdx < maxLine; lineIdx++ {
			if sepByLines[sepIdx][lineIdx] != target {
				same = false
				break
			}
		}
		if target > 0 && same {
			return sepList[sepIdx]
		}
	}
	return ""
}
func printResult(separator string) {
	switch separator {
	case "\t":
		fmt.Printf("tab-separated\n")
	case "":
		fmt.Printf("white-space separated or not separated at all\n")
	default:
		fmt.Printf("%s-separated\n", separator)
	}
}
