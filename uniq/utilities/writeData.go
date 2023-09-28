package utilities

import (
	"errors"
	"fmt"
	"os"
)

func WriteData(outputFile *os.File, processedLines []string) error {
	if outputFile == nil {
		for _, line := range processedLines {
			fmt.Println(line)
		}
	} else {
		for _, line := range processedLines {
			_, err := outputFile.WriteString(line + "\n")
			if err != nil {
				err = errors.New("error while writing data")
				return err
			}
		}
	}

	return nil
}
