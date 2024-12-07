package main

import (
	"aoc2024/src/util"
	"fmt"
)

type Reader struct{}

// onLine appends a line to the slice
func (r *Reader) OnLine(line string) {
	fmt.Println(line)
}

// collect returns the collected lines
func (r *Reader) Done() {
	fmt.Println("final done")
}

func main() {
	// Create a new StringSliceReader
	reader := &Reader{}

	// Call ReadFile with the StringSlice
	filePath := "./src/day1/input.txt"
	util.ReadFile(filePath, reader)
}
