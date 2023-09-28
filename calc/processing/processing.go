package processing

import (
	"calc/utilities"
	"errors"
	"regexp"
	"strconv"
)

func containsOnlyValidSymbols(line string) bool {
	regex := regexp.MustCompile(`^[+\-/*\d.()]+$`)

	return regex.MatchString(line)
}

func calculateInterfaces(left, right interface{}, operator string) interface{} {
	leftFloat, leftIsFloat := left.(float64)
	rightFloat, rightIsFloat := right.(float64)

	if leftIsFloat && rightIsFloat {
		switch operator {
		case "+":
			return leftFloat + rightFloat
		case "-":
			return leftFloat - rightFloat
		case "*":
			return leftFloat * rightFloat
		case "/":
			if rightFloat != 0 {
				return leftFloat / rightFloat
			} else {
				return nil
			}
		}
	}
	return nil
}

func priority(operation interface{}) int {
	if operation == "+" || operation == "-" {
		return 1
	} else if operation == "*" || operation == "/" {
		return 2
	} else {
		return -1
	}
}

func processOperations(stackNumbers *utilities.Stack, operation interface{}) {
	right := stackNumbers.Peek()
	stackNumbers.Pop()
	left := stackNumbers.Peek()
	stackNumbers.Pop()

	switch operation {
	case "+":
		stackNumbers.Push(calculateInterfaces(left, right, "+"))
	case "-":
		stackNumbers.Push(calculateInterfaces(left, right, "-"))
	case "/":
		stackNumbers.Push(calculateInterfaces(left, right, "/"))
	case "*":
		stackNumbers.Push(calculateInterfaces(left, right, "*"))
	}

}
func IsNotNumber(value interface{}) bool {
	return value == "+" || value == "-" || value == "*" || value == "/" || value == "(" || value == ")"
}

func Calc(inputLine string) (interface{}, error) {
	if !containsOnlyValidSymbols(inputLine) {
		err := errors.New("invalid symbols are given")
		return nil, err
	}
	stackNumbers := utilities.NewStack()
	stackOperations := utilities.NewStack()

	for i := 0; i < len(inputLine); i++ {
		value := string(inputLine[i])
		if value == "(" {
			stackOperations.Push(value)
		} else if value == ")" {
			for stackOperations.Peek() != "(" {
				if stackOperations.IsEmpty() {
					err := errors.New("expression is not written correctly")
					return nil, err
				}
				processOperations(stackNumbers, stackOperations.Peek())
				stackOperations.Pop()
			}
			stackOperations.Pop()
		} else if IsNotNumber(value) {
			for !stackOperations.IsEmpty() && priority(stackOperations.Peek()) >= priority(value) {
				processOperations(stackNumbers, stackOperations.Peek())
				stackOperations.Pop()
			}
			stackOperations.Push(value)
		} else {
			operand := ""
			for i < len(inputLine) && !IsNotNumber(string(inputLine[i])) {
				operand += string(inputLine[i])
				i++
			}
			i--
			floatValue, _ := strconv.ParseFloat(operand, 64)
			stackNumbers.Push(floatValue)
		}
	}
	for !stackOperations.IsEmpty() {
		processOperations(stackNumbers, stackOperations.Peek())
		stackOperations.Pop()
	}
	return stackNumbers.Peek(), nil
}
