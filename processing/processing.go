package uniq

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Options struct {
	CFlagStated bool
	DFlagStated bool
	UFlagStated bool
	FFlagStated int
	SFlagStated int
	IFlagStated bool
}

func SetArgs() (Options, error) {
	var options Options

	flag.BoolVar(&options.CFlagStated, "c", false, "Сount the number of occurrences of a string and"+
		" print this number before the string separated by a space.")
	flag.BoolVar(&options.DFlagStated, "d", false, "Output only those lines that are repeated "+
		"in the input data.")
	flag.BoolVar(&options.UFlagStated, "u", false, "Output only those lines that are not "+
		"repeated in the input data.")
	flag.IntVar(&options.FFlagStated, "f", 0, "Ignore the first numFields of fields in a row. "+
		"A field in a string is a non-empty set of characters separated by a space.")
	flag.IntVar(&options.SFlagStated, "s", 0, "Ignore the first numFields of fields in a row. "+
		"A field in a string is a non-empty set of characters separated by a space.")
	flag.BoolVar(&options.IFlagStated, "i", false, "Ignore the case of letters.")

	flag.Parse()

	if (options.CFlagStated && options.DFlagStated) ||
		(options.CFlagStated && options.UFlagStated) ||
		(options.DFlagStated && options.UFlagStated) {
		err := fmt.Errorf("parameters 'c', 'd' and 'u' cannot be used at the same time")
		fmt.Println("Use Case: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		flag.PrintDefaults()
		return options, err
	} else if options.SFlagStated < 0 || options.FFlagStated < 0 {
		err := fmt.Errorf("num_fields or num_chars cannot be negative")
		fmt.Println("Use Case: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		flag.PrintDefaults()
		return options, err
	}

	return options, nil
}

func (options Options) LinesProcessing(lines []string) []string {
	var processedLines []string
	comparisonFunction := comparisonStrings
	if options.IFlagStated {
		comparisonFunction = strings.EqualFold
	}
	if options.CFlagStated {
		processedLines = cFlagLinesProcessing(lines)
	} else if options.DFlagStated {
		processedLines = dFlagLinesProcessing(lines)
	} else if options.UFlagStated {
		processedLines = uFlagLinesProcessing(lines)
	} else {
		processedLines = otherFlagsStatedProcessing(lines, options.FFlagStated, options.SFlagStated, comparisonFunction)
	}

	return processedLines
}

func otherFlagsStatedProcessing(linesWithoutNumFieldsAndChars []string, numFields int, numChars int,
	stringsComparison func(lineOne, lineTwo string) bool) []string {
	space := ""
	if numFields != 0 {
		space = " "
	}

	linesWithNumFieldsAndChars := make([]string, len(linesWithoutNumFieldsAndChars))
	for index, value := range linesWithoutNumFieldsAndChars {
		fields := strings.Fields(value)
		if len(fields) < numFields || utf8.RuneCountInString(strings.Join(fields[numFields:], " ")) < numChars {
			linesWithNumFieldsAndChars[index] = value
			linesWithoutNumFieldsAndChars[index] = ""
		} else {
			linesWithNumFieldsAndChars[index] = strings.Join(fields[:numFields], " ") +
				space +
				strings.Join(fields[numFields:], " ")[:numChars]

			linesWithoutNumFieldsAndChars[index] = strings.Join(fields[numFields:], " ")[numChars:]
		}
	}

	var processedLines []string
	for index, value := range linesWithoutNumFieldsAndChars {
		if index == 0 || !stringsComparison(linesWithoutNumFieldsAndChars[index-1], value) {
			processedLines = append(processedLines, linesWithNumFieldsAndChars[index]+value)
		}
	}

	return processedLines
}
func comparisonStrings(lineOne, lineTwo string) bool {
	return lineOne == lineTwo
}

func uFlagLinesProcessing(lines []string) []string {
	var processedLines []string
	count := 1
	for i := 1; i < len(lines); i++ {
		if lines[i-1] != lines[i] {
			if count == 1 {
				processedLines = append(processedLines, lines[i-1])
			}
			count = 1
		} else {
			count++
		}
	}

	if count == 1 {
		processedLines = append(processedLines, lines[len(lines)-1])
	}

	return processedLines
}

func dFlagLinesProcessing(lines []string) []string {
	var processedLines []string
	count := 0
	for i := 1; i < len(lines); i++ {
		if lines[i-1] == lines[i] && count == 0 {
			processedLines = append(processedLines, lines[i])
			count++
		} else {
			count = 0
		}
	}

	return processedLines
}

func cFlagLinesProcessing(lines []string) []string {
	count := 1
	var processedLines []string
	for i := 1; i < len(lines); i++ {
		if lines[i-1] != lines[i] {
			processedLines = append(processedLines, strconv.Itoa(count)+" "+lines[i-1])
			count = 1
		} else {
			count++
		}
	}

	processedLines = append(processedLines, strconv.Itoa(count)+" "+lines[len(lines)-1])

	return processedLines
}
