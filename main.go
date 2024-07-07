package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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
	var outputBytes = flag.Bool("c", false, "to output the number of bytes")
	var outputLines = flag.Bool("l", false, "to output the number of bytes")
	var fileSource = flag.String("f", "", "locate the file source")

	flag.Parse()

	dat, err := os.ReadFile(*fileSource)
	check(err)

	ccwc := New(dat)

	if *outputBytes {
		ccwc.countBytes()
	}
	if *outputLines {
		ccwc.countLines()
	}

	for _, number := range ccwc.FinalOutput {
		fmt.Printf("%d  ", number)
		fmt.Println("")
	}
}

func (wc *CCWC) countBytes() {
	wc.FinalOutput = append(wc.FinalOutput, len(wc.Input))
}

func (wc *CCWC) countLines() {
	lines := strings.Split(string(wc.Input), "\n")
	wc.FinalOutput = append(wc.FinalOutput, len(lines)-1)
}

func (wc *CCWC) countWords() {
	words := strings.Split(string(wc.Input), " ")
	wc.FinalOutput = append(wc.FinalOutput, len(words)-1)
}
