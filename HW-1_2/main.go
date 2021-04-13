package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	operators = "*+-()/"
)

func main() {
	expression, err := getCmdArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	polishStr, err := getPolishNotation(expression)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := getResult(polishStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func getCmdArgs() (string, error) {
	if len(os.Args) == 1 {
		return "", errors.New("Empty comand line arguments")
	} else if len(os.Args) > 2 {
		return "", errors.New("Expected one comand line argument")
	}
	return os.Args[1], nil
}

func isOperator(char rune) bool {
	return strings.ContainsRune(operators, char)
}

func isDigit(char rune) bool {
	_, err := strconv.Atoi(string(char))
	return err == nil
}

func checklastDigit(char rune) bool {
	_, err := strconv.Atoi(string(char))
	return err == nil
}

func getPriority(char rune) int {
	switch char {
	case '(':
		return 0
	case ')':
		return 1
	case '+':
		return 2
	case '-':
		return 2
	case '*':
		return 3
	case '/':
		return 3
	}
	return 4
}

func getPolishNotation(inputStr string) (string, error) {
	stack := NewRuneStack()
	var polishStr string
	inputStr = strings.ReplaceAll(inputStr, " ", "")
	inputStr = strings.ReplaceAll(inputStr, "\"", "")
	lastDigit := false

	for _, char := range inputStr {
		if isOperator(char) {
			lastDigit = false
			switch char {
			case '(':
				stack.Push(char)
			case ')':
				r, err := stack.Pop()
				if err != nil {
					return "", err
				}
				for r != '(' {
					polishStr += " " + string(r)
					r, err = stack.Pop()
					if err != nil {
						return "", err
					}
				}
			default:
				if !stack.IsEmpty() {
					if getPriority(char) <= getPriority(stack.Top()) {
						polishStr += " " + string(char)
						continue
					}
				}
				stack.Push(char)
			}
		} else if isDigit(char) {
			if !lastDigit {
				polishStr += " "
				lastDigit = true
			}
			polishStr += string(char)
		} else {
			return "", errors.New("Wrong symbols in expression")
		}
	}
	for !stack.IsEmpty() {
		r, err := stack.Pop()
		if err != nil {
			return "", err
		}
		polishStr += " " + string(r)
	}
	return polishStr, nil
}

func getResult(polishStr string) (int, error) {
	var result int
	var currentNumber string
	stack := NewIntStack()
	lastDigit := false

	for _, char := range polishStr {
		if isOperator(char) {
			currentNumber = ""
			lastDigit = false
			r, err := stack.Pop()
			if err != nil {
				return 0, err
			}
			l, err := stack.Pop()
			if err != nil {
				return 0, err
			}
			switch char {
			case '+':
				result = l + r
			case '-':
				result = l - r
			case '*':
				result = l * r
			case '/':
				if r == 0 {
					return 0, errors.New("Division by zero")
				}
				result = l / r
			}
			stack.Push(result)
		} else if isDigit(char) {
			if !lastDigit {
				lastDigit = true
				currentNumber = ""
			}
			currentNumber += string(char)
		} else {
			if currentNumber != "" {
				number, _ := strconv.Atoi(currentNumber)
				stack.Push(number)
				currentNumber = ""
			}
		}
	}
	return stack.Top(), nil
}
