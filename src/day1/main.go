package main

import (
	"aoc2024/src/util"
	"fmt"
)

type Reader struct {
	lines [][]int
}

func (r *Reader) OnLine(line string) {
	intArray := util.ConvertToIntArray(line)
	fmt.Println("Converted integer array:", intArray)

	r.lines = append(r.lines, intArray)
}
func (r *Reader) Collect() [][]int {
	fmt.Println("done")
	return r.lines
}

func main() {
	reader := &Reader{}

	filePath := "./src/day1/input.txt"
	util.ReadFileCollect(filePath, reader)
}
