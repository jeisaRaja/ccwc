package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

type CCWC struct {
	Input       []byte
	FinalOutput []int
}

func New(input []byte) *CCWC {
	return &CCWC{Input: input}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var outputBytes = flag.Bool("c", false, "output the number of bytes")
	var outputLines = flag.Bool("l", false, "output the number of lines")
	var outputWords = flag.Bool("w", false, "output the number of words")
	var outputCharacters = flag.Bool("m", false, "output the number of characters")
	var fileSource = flag.String("f", "", "locate the file source")

	flag.Parse()

	var input []byte

	if *fileSource != "" {
		dat, err := os.ReadFile(*fileSource)
		check(err)
		input = dat
	} else {
		dat, err := io.ReadAll(os.Stdin)
		check(err)
		input = dat
	}

	ccwc := New(input)

	if *outputBytes {
		ccwc.countBytes()
	}
	if *outputLines {
		ccwc.countLines()
	}
	if *outputWords {
		ccwc.countWords()
	}
	if *outputCharacters {
		ccwc.countCharacters()
	}

	if !*outputBytes && !*outputLines && !*outputWords && !*outputCharacters {
		ccwc.countLines()
		ccwc.countWords()
		ccwc.countBytes()
	}

	for _, number := range ccwc.FinalOutput {
		fmt.Printf("%d  ", number)
	}
	fmt.Println("")
}

func (wc *CCWC) countBytes() {
	wc.FinalOutput = append(wc.FinalOutput, len(wc.Input))
}

func (wc *CCWC) countLines() {
	lines := strings.Split(string(wc.Input), "\n")
	wc.FinalOutput = append(wc.FinalOutput, len(lines)-1)
}

func (wc *CCWC) countWords() {
	words := strings.Fields(string(wc.Input))
	wc.FinalOutput = append(wc.FinalOutput, len(words))
}

func (wc *CCWC) countCharacters() {
	lenCharacters := utf8.RuneCountInString(string(wc.Input))
	wc.FinalOutput = append(wc.FinalOutput, lenCharacters)
}
