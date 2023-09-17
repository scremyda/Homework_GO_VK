package utilities

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadLines(reader io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		err = fmt.Errorf("error while reading data")
		return nil, err
	}
	return lines, nil
}

func GetData(inputFile *os.File) ([]string, error) {
	if inputFile == nil {
		return ReadLines(os.Stdin)
	}
	return ReadLines(inputFile)
}
