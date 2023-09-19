package main

import (
	"flag"
	"fmt"
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

	defer func() {
		if err = inputFile.Close(); err != nil {
			err = fmt.Errorf("error closing file: %s", argsWithoutFlags[0])
			fmt.Println("Error happened while closing file: ", err)
			return
		}
	}()

	defer func() {
		if err = outputFile.Close(); err != nil {
			err = fmt.Errorf("error closing file: %s", argsWithoutFlags[1])
			fmt.Println("Error happened while closing file: ", err)
			return
		}
	}()

	inputData, err := utilities.GetData(inputFile)
	if err != nil {
		fmt.Println("Error happened while reading: ", err)
		return
	}

	processedLines := options.LinesProcessing(inputData)

	err = utilities.WriteData(outputFile, processedLines)
	if err != nil {
		fmt.Println("Error happened while writing: ", err)
		return
	}
}
