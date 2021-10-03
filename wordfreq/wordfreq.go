package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage ./%s filename\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	wordsList := map[string]int{}
	for _, fileName := range os.Args[1:] {
		updateWordListFromFiles(fileName, wordsList)
	}
	printSortedFreqByWord(wordsList)
	wordByFreq := invertWordList(wordsList)
	printSortedWordByFreq(wordByFreq)

}

func updateWordListFromFiles(fileName string, wordList map[string]int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("%s\n", err)
	}
	defer file.Close()
	readAndUpdateWordsFromFile(bufio.NewReader(file), wordList)
}

func readAndUpdateWordsFromFile(reader *bufio.Reader, wordList map[string]int) {
	for {
		line, err := reader.ReadString('\n')
		for _, word := range splitByNoLetters(strings.TrimSpace(line)) {
			if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 1 {
				wordList[strings.ToLower(word)] += 1
			}
		}
		if err != nil {
			if err != io.EOF {
				log.Printf("failed to finishe reading file %s\n", err)
			}
			break
		}
	}
}

func splitByNoLetters(line string) []string {
	notALetter := func (r rune) bool {return !unicode.IsLetter(r)}
	return strings.FieldsFunc(line, notALetter)
}

func printSortedFreqByWord(wordList map[string]int) {
	sortedWordList := make([]string, 0, len(wordList))
	wordWidth, freqWidth := 0, 0
	for word, freq := range wordList {
		sortedWordList = append(sortedWordList, word)
		if wordLen := utf8.RuneCountInString(word); wordLen > wordWidth {
			wordWidth = wordLen
		}
		if freqLen := len(fmt.Sprint(freq)); freqLen > freqWidth {
			freqWidth = freqLen
		}
	}
	sort.Strings(sortedWordList)
	gap := wordWidth + freqWidth - len("word") - len("freq")
	fmt.Printf("Word %*s%s\n", gap, " ", "Freq")
	for _, word := range sortedWordList {
		fmt.Printf("%-*s %*d\n", wordWidth, word, freqWidth ,wordList[word])
	}
}

func invertWordList(wordList map[string]int) map[int][]string {
	wordsByFreq := make(map[int][]string, len(wordList))
	for word, freq := range wordList {
		wordsByFreq[freq] = append(wordsByFreq[freq], word)
	}
	return wordsByFreq
}

func printSortedWordByFreq(wordList map[int][]string) {
	freqList := make([]int, 0, len(wordList))
	for freq := range wordList {
		freqList = append(freqList, freq)
	}
	sort.Ints(freqList)
	fmt.Printf("words by freq\n")
	for _, freq := range freqList {
		words := wordList[freq]
		sort.Strings(words)
		fmt.Printf("%d\t%s\n", freq, strings.Join(words, ", "))
	}
}