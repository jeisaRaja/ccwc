package main

import (
	"flag"
	"fmt"
	"os"
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
	var outputBytes = flag.Bool("bytes", false, "to output the number of bytes")
	var fileSource = flag.String("file", "", "locate the file source")

	flag.Parse()

	dat, err := os.ReadFile(*fileSource)
	check(err)

	ccwc := New(dat)

	if *outputBytes {
		ccwc.countBytes()
	}

	for _, number := range ccwc.FinalOutput {
		fmt.Printf("%d  ", number)
    fmt.Println("")
	}
}

func (wc *CCWC) countBytes() {
	wc.FinalOutput = append(wc.FinalOutput, len(wc.Input))
}
