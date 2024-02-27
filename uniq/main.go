package main

import (
	"flag"
	"fmt"
	"uniq/processing"
	"uniq/utilities"
)

func main() {
	options, err := utilities.SetArgs()

	if err != nil {
		fmt.Println("Error happened: ", err)
		return
	}

	argsWithoutFlags := flag.Args()

	inputFile, outputFile := utilities.OpenFiles(argsWithoutFlags)
	if len(argsWithoutFlags) > 0 {
		if inputFile != nil {
			defer func() {
				if err = inputFile.Close(); err != nil {
					err = fmt.Errorf("error closing file: %s", argsWithoutFlags[0])
					fmt.Println("Error happened while closing input file: ", err)
				}
			}()
		}
	}
	if len(argsWithoutFlags) == 2 {
		if outputFile != nil {
			defer func() {
				if err = outputFile.Close(); err != nil {
					err = fmt.Errorf("error closing file: %s", argsWithoutFlags[1])
					fmt.Println("Error happened while closing output file: ", err)
				}
			}()
		}
	}

	inputData, err := utilities.GetData(inputFile)
	if err != nil {
		fmt.Println("Error happened while reading: ", err)
		return
	}

	processedLines := uniq.LinesProcessing(options, inputData)

	err = utilities.WriteData(outputFile, processedLines)
	if err != nil {
		fmt.Println("Error happened while writing: ", err)
		return
	}
}
