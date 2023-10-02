package main

import (
	"bufio"
	"calc/processing"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error happened while reading:", err)
		return
	}

	textWithoutSpaces := strings.ReplaceAll(inputString, " ", "")

	cleanedText := strings.TrimSpace(textWithoutSpaces)

	answer, err := processing.Calc(cleanedText)
	if err != nil {
		fmt.Println("Error happened while calculating:", err)
		return
	}

	fmt.Println(answer)
}
