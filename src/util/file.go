package util

import (
	"bufio"
	"os"
)

type Reader interface {
	OnLine(line string)
}

type LineReader interface {
	OnLine(line string)
	Done()
}

type LineReaderCollector[T any] interface {
	OnLine(line string)
	Collect() T
}

type LineReaderCollectorP[T any] interface {
	OnLine(line string)
	Collect() *T
}

func readLines(filePath string, reader Reader) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reader.OnLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func ReadFile(filePath string, reader LineReader) error {
	readLines(filePath, reader)
	reader.Done()
	return nil
}

func ReadFileCollect[T any](filePath string, reader LineReaderCollector[T]) (T, error) {
	readLines(filePath, reader)
	return reader.Collect(), nil
}

func ReadFileCollectP[T any](filePath string, reader LineReaderCollectorP[T]) (*T, error) {
	readLines(filePath, reader)
	return reader.Collect(), nil
}
