package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) == 1 || !strings.HasSuffix(os.Args[1], ".ini") {
		fmt.Printf("usage: %s file[.ini]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	iniData := parseFile(os.Args[1])
	data := parseIni(iniData)
	printIni(data)
}

func parseIni(lines []string) map[string]map[string]string {
	const separator = "="
	collection := make(map[string]map[string]string)
	group := "General"
	for _, line := range lines {
		if line == "" || strings.HasPrefix(line,";") {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			group = line[1:len(line) - 1]
		} else if i := strings.Index(line, separator); i > - 1 {
			key := line[:i]
			value := line[len(key) + 1:]
			if _, found := collection[group]; !found {
				collection[group] = make(map[string]string)
			}
			collection[group][key] = value
		} else {
			log.Print("error to parse file")
		}
	}
	return collection
}
func printIni(data map[string]map[string]string) {
	if len(data) == 0 || data == nil {
		fmt.Printf("*.ini file is empty\n")
	}
	sortGroups := make([]string, 0, len(data))
	for group := range data {
		sortGroups = append(sortGroups, group)
	}
	sort.Strings(sortGroups)
	for i, group := range sortGroups {
		fmt.Printf("[%s]\n", group)
		keySlice := make([]string, 0, len(data[group]))
		for key := range data[group] {
			keySlice = append(keySlice, key)
		}
		sort.Strings(keySlice)
		for _, key := range keySlice {
			fmt.Printf("%s=%s\n", key, data[group][key])
		}
		if i + 1 < len(sortGroups) {
			fmt.Printf("\n")
		}
	}
}
func parseFile(filename string) []string {
	var file *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return getLinesFromFile(bufio.NewScanner(file))
}
func getLinesFromFile(scanner *bufio.Scanner) []string {
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}