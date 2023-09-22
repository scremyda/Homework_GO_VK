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
	switch leftType := left.(type) {
	case int:
		switch rightType := right.(type) {
		case int:
			switch operator {
			case "+":
				return leftType + rightType
			case "-":
				return leftType - rightType
			case "*":
				return leftType * rightType
			case "/":
				if rightType != 0 {
					return float64(leftType) / float64(rightType)
				} else {
					return nil
				}
			}
		case float64:
			switch operator {
			case "+":
				return float64(leftType) + rightType
			case "-":
				return float64(leftType) - rightType
			case "*":
				return float64(leftType) * rightType
			case "/":
				if rightType != 0 {
					return float64(leftType) / rightType
				} else {
					return nil
				}
			}
		}
	case float64:
		switch rightType := right.(type) {
		case int:
			switch operator {
			case "+":
				return leftType + float64(rightType)
			case "-":
				return leftType - float64(rightType)
			case "*":
				return leftType * float64(rightType)
			case "/":
				if rightType != 0 {
					return leftType / float64(rightType)
				} else {
					return nil
				}
			}
		case float64:
			switch operator {
			case "+":
				return leftType + rightType
			case "-":
				return leftType - rightType
			case "*":
				return leftType * rightType
			case "/":
				if rightType != 0 {
					return leftType / rightType
				} else {
					return nil
				}
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

	if operation == "+" {
		stackNumbers.Push(calculateInterfaces(left, right, "+"))
	} else if operation == "-" {
		stackNumbers.Push(calculateInterfaces(left, right, "-"))
	} else if operation == "/" {
		stackNumbers.Push(calculateInterfaces(left, right, "/"))
	} else if operation == "*" {
		stackNumbers.Push(calculateInterfaces(left, right, "*"))
	}

}
func isOperand(value interface{}) bool {
	return value == "+" || value == "-" || value == "*" || value == "/" || value == "(" || value == ")"
}

func Calc(inputLine string) (interface{}, error) {
	if !containsOnlyValidSymbols(inputLine) {
		err := errors.New("invalid symbols are given")
		return nil, err
	}
	stackNumbers := new(utilities.Stack)
	stackOperations := new(utilities.Stack)

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
		} else if isOperand(value) {
			for !stackOperations.IsEmpty() && priority(stackOperations.Peek()) >= priority(value) {
				processOperations(stackNumbers, stackOperations.Peek())
				stackOperations.Pop()
			}
			stackOperations.Push(value)
		} else {
			operand := ""
			for i < len(inputLine) && !isOperand(string(inputLine[i])) {
				operand += string(inputLine[i])
				i++
			}
			i--
			valueInt, err := strconv.Atoi(operand)
			if err == nil {
				stackNumbers.Push(valueInt)
			} else {
				floatValue, _ := strconv.ParseFloat(operand, 64)
				stackNumbers.Push(floatValue)
			}

		}
	}
	for !stackOperations.IsEmpty() {
		processOperations(stackNumbers, stackOperations.Peek())
		stackOperations.Pop()
	}
	return stackNumbers.Peek(), nil
}
