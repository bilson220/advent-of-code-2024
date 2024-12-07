package util

import (
	"bufio"
	"os"
)

type LineReader[T any] interface {
	onLine(line string)
	collect() *T
}

func ReadFileLines[T any](filePath string, reader LineReader[T]) (*T, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reader.onLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reader.collect(), nil
}
