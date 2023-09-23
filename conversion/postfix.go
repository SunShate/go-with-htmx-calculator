package conversion

import (
	"fmt"
	"strconv"

	"calculator/operation"
	"calculator/utils"
)

const rightRoundBracket = ")"
const leftRoundBracket = "("

func BuildPostfix(infixExpr []string) ([]string, bool) {
	var postfixExpr []string

	var opStack utils.Stack

	for i := 0; i < len(infixExpr); i++ {
		lexem := infixExpr[i]
		switch lexem {
		case leftRoundBracket:
			opStack.Push(lexem)
		case rightRoundBracket:
			element, _ := opStack.Pop()

			for !opStack.IsEmpty() &&
				element != leftRoundBracket {

				postfixExpr = append(postfixExpr, element.(string))
				element, _ = opStack.Pop()
			}
		case "+", "-", "*", "/":
			element := opStack.Peek()
			if !opStack.IsEmpty() &&
				operation.OperationPriorities[element.(string)] >= operation.OperationPriorities[lexem] {

				postfixExpr = append(postfixExpr, element.(string))
				opStack.Pop()
			}
			opStack.Push(lexem)
		default:
			if utils.NumberRegexp.MatchString(lexem) {
				postfixExpr = append(postfixExpr, lexem)
			} else {
				fmt.Println("ERROR: UNSUPPORTED SYMBOL")
				return nil, true
			}
		}
	}

	for !opStack.IsEmpty() {
		element, _ := opStack.Pop()
		postfixExpr = append(postfixExpr, element.(string))
	}

	return postfixExpr, false
}

func CalculateExpr(postfixExpr []string) (uint64, bool) {
	var stack utils.Stack
	var firstValueMet, secondValueMet, isError bool

	for i := 0; i < len(postfixExpr); i++ {
		lexem := postfixExpr[i]

		switch lexem {
		case "+", "-", "*", "/": // operations
			if secondValueMet {
				secondValue, _ := stack.Pop()
				firstValue, _ := stack.Pop()

				_operation := operation.New(firstValue.(uint64), secondValue.(uint64), uint8(lexem[0]))

				fmt.Println("Operation:", _operation)

				computeResult := operation.Compute(_operation)

				stack.Push(computeResult)

				if len(stack) < 1 {
					firstValueMet, secondValueMet = false, false
				}
			}
		default:
			number, _ := strconv.Atoi(lexem) // convert string to int

			if firstValueMet {
				secondValueMet = true
			} else {
				firstValueMet = true
			}

			stack.Push(uint64(number))
		}
	}

	result, _ := stack.Pop()

	if !stack.IsEmpty() {
		isError = true
		fmt.Println("ERROR: INVALID EXPRESSION")
	}

	return result.(uint64), isError
}
