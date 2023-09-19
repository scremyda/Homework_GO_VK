package uniq

import (
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

type CountOptions struct {
	line  string
	count int
}

func comparisonStrings(lineOne, lineTwo string) bool {
	return lineOne == lineTwo
}

func (options Options) LinesProcessing(lines []string) []string {
	var processedLines []string
	comparisonFunction := comparisonStrings

	if options.CFlagStated {
		for _, value := range uniq(lines, options.FFlagStated, options.SFlagStated, comparisonFunction) {
			processedLines = append(processedLines, strconv.Itoa(value.count)+" "+value.line)
		}
	} else if options.DFlagStated {
		for _, value := range uniq(lines, options.FFlagStated, options.SFlagStated, comparisonFunction) {
			if value.count > 1 {
				processedLines = append(processedLines, value.line)
			}
		}
	} else if options.UFlagStated {
		for _, value := range uniq(lines, options.FFlagStated, options.SFlagStated, comparisonFunction) {
			if value.count == 1 {
				processedLines = append(processedLines, value.line)
			}
		}
	} else if options.IFlagStated {
		comparisonFunction = strings.EqualFold
		for _, value := range uniq(lines, options.FFlagStated, options.SFlagStated, comparisonFunction) {
			processedLines = append(processedLines, value.line)
		}
	} else {
		for _, value := range uniq(lines, options.FFlagStated, options.SFlagStated, comparisonFunction) {
			processedLines = append(processedLines, value.line)
		}
	}

	return processedLines
}

func uniq(linesWithoutNumFieldsAndChars []string, numFields int, numChars int,
	stringsComparison func(lineOne, lineTwo string) bool) []CountOptions {
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

	count := 1
	var processedLines []CountOptions
	var index int
	for index = 1; index < len(linesWithoutNumFieldsAndChars); index++ {
		if !stringsComparison(linesWithoutNumFieldsAndChars[index-1], linesWithoutNumFieldsAndChars[index]) {
			processedLines = append(processedLines, CountOptions{linesWithNumFieldsAndChars[index-count] +
				linesWithoutNumFieldsAndChars[index-count], count})
			count = 1
		} else {
			count++
		}
	}
	if count == 1 {
		processedLines = append(processedLines, CountOptions{linesWithNumFieldsAndChars[index-count] +
			linesWithoutNumFieldsAndChars[index-count], count})
	}

	return processedLines
}
