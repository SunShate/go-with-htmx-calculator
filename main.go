package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"calculator/operation"
)

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(obj interface{}) {
	*s = append(*s, obj)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return os.DevNull, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return os.DevNull
	} else {
		index := len(*s) - 1
		return (*s)[index]
	}
}

func main() {
	fmt.Println("Interface will add sooner")

	var stack Stack
	var firstValueMet, secondValueMet bool

	// arithmetic expression to postfix expression converter

	postfixExpression := strings.Split(os.Args[1], " ")

	for i := 0; i < len(postfixExpression); i++ {
		lexem := postfixExpression[i]

		switch lexem {
		case "+", "-", "*", "/": // operations
			if secondValueMet {
				firstValue, _ := stack.Pop()
				secondValue, _ := stack.Pop()

				_operation := operation.New(firstValue.(uint64), secondValue.(uint64), uint8(lexem[0]))

				computeResult := operation.Compute(_operation)

				stack.Push(computeResult)

				if len(stack) < 1 {
					firstValueMet, secondValueMet = false, false
				}
			}
		default:
			number, _ := strconv.Atoi(lexem) // convert char to int

			if firstValueMet {
				secondValueMet = true
			} else {
				firstValueMet = true
			}

			stack.Push(uint64(number))
		}
	}

	fmt.Println(stack.Pop())
}
