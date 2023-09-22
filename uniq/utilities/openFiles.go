package utilities

import (
	"fmt"
	"os"
)

func OpenFiles(argsWithoutFlags []string) (inputFile, outputFile *os.File) {
	if len(argsWithoutFlags) > 0 {
		var err error
		inputFile, err = os.Open(argsWithoutFlags[0])
		if err != nil {
			err = fmt.Errorf("can't open input file: %s", argsWithoutFlags[0])
			fmt.Println("Error happened while opening file: ", err)
			return
		}
		if len(argsWithoutFlags) == 2 {
			outputFile, err = os.Create(argsWithoutFlags[1])
			if err != nil {
				err = fmt.Errorf("can't create output file: %s", argsWithoutFlags[1])
				fmt.Println("Error happened while creating file: ", err)
				return
			}
		}
	}
	return inputFile, outputFile
}
