package conversion

import (
	"calculator/operation"
	"calculator/utils"
	"fmt"
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

			for element != nil &&
				element != leftRoundBracket {

				postfixExpr = append(postfixExpr, element.(string))
				element, _ = opStack.Pop()
			}
		case "+", "-", "*", "/":
			element := opStack.Peek()
			if element != nil &&
				operation.OperationPriorities[element.(string)] != 0 &&
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
