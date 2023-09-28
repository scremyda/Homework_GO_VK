package main

import (
	"bufio"
	"calc/processing"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')
	if err != nil {
		err = errors.New("ReadString encounters an error before finding a delimiter")
		fmt.Println("Error happened while reading:", err)
		return
	}

	textWithoutSpaces := strings.ReplaceAll(inputString, " ", "")

	cleanedText := strings.TrimSpace(textWithoutSpaces)

	answer, err := processing.Calc(cleanedText)
	if err != nil {
		fmt.Println("Error happened while calculating:", err)
		return
	} else if answer == nil {
		err = errors.New("expression is not written correctly")
		fmt.Println("Error happened while calculating:", err)
		return
	}

	fmt.Println(answer)
}
