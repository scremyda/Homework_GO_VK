package main

import (
	"flag"
	"fmt"
	"os"
	"uniq/utilities"
)

func main() {
	options, err := utilities.SetArgs()

	if err != nil {
		fmt.Println("Error happened: ", err)
		return
	}

	argsWithoutFlags := flag.Args()

	var inputFile *os.File
	var outputFile *os.File
	if len(argsWithoutFlags) > 0 && len(argsWithoutFlags) < 3 {
		inputFile, err = os.Open(argsWithoutFlags[0])
		if err != nil {
			err = fmt.Errorf("can't open input file: %s", flag.Args()[0])
			fmt.Println("Error happened: ", err)
			return
		}
		defer func() {
			if err = inputFile.Close(); err != nil {
				err = fmt.Errorf("error closing file: %s", flag.Args()[0])
				fmt.Println("Error happened: ", err)
				return
			}
		}()
		if len(argsWithoutFlags) == 2 {
			outputFile, err = os.Create(argsWithoutFlags[1])
			if err != nil {
				err = fmt.Errorf("can't create output file: %s", flag.Args()[1])
				fmt.Println("Error happened: ", err)
				return
			}
			defer func() {
				if err = outputFile.Close(); err != nil {
					err = fmt.Errorf("error closing file: %s", flag.Args()[1])
					fmt.Println("Error happened: ", err)
					return
				}
			}()
		}
	}
	inputData, err := utilities.GetData(inputFile)
	if err != nil {
		fmt.Println("Error happened: ", err)
		return
	}

	processedLines := options.LinesProcessing(inputData)

	err = utilities.WriteData(outputFile, processedLines)
	if err != nil {
		fmt.Println("Error happened: ", err)
		return
	}
}
