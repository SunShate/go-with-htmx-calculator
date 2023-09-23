package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"calculator/operation"
	"calculator/postfix"
	"calculator/utils"
)

func main() {
	fmt.Println("Interface will add sooner")

	var stack utils.Stack
	var firstValueMet, secondValueMet bool

	// arithmetic expression to postfix expression converter
	arithmeticExpression := strings.Split(os.Args[1], " ")
	postfixExpression, isError := postfix.Build(arithmeticExpression)

	if isError {
		return
	}

	fmt.Println(postfixExpression)

	// postfix calculation
	for i := 0; i < len(postfixExpression); i++ {
		lexem := postfixExpression[i]

		switch lexem {
		case "+", "-", "*", "/": // operations
			if secondValueMet {
				secondValue, _ := stack.Pop()
				firstValue, _ := stack.Pop()

				_operation := operation.New(firstValue.(uint64), secondValue.(uint64), uint8(lexem[0]))

				fmt.Println(_operation)

				computeResult := operation.Compute(_operation)

				stack.Push(computeResult)

				if len(stack) < 1 {
					firstValueMet, secondValueMet = false, false
				}
			}
		default:
			var number int

			if utils.NumberRegexp.MatchString(lexem) {
				number, _ = strconv.Atoi(lexem) // convert char to int
			} else {
				return
			}

			if firstValueMet {
				secondValueMet = true
			} else {
				firstValueMet = true
			}

			stack.Push(uint64(number))
		}
	}

	fmt.Println(stack.Pop())

	if !stack.IsEmpty() {
		fmt.Println("ERROR: INVALID EXPRESSION")
	}
}
