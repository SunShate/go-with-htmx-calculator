package main

import (
	"fmt"
	"os"
	"strings"

	"calculator/conversion"
)

func main() {
	fmt.Println("Interface will be sooner")

	arithmeticExpression := strings.Split(os.Args[1], " ")

	// arithmetic expression to postfix expression converter
	postfixExpression, isError := conversion.BuildPostfix(arithmeticExpression)

	if isError {
		return
	}

	fmt.Println("\nPostfix notation:", postfixExpression)

	// postfix calculation
	result, isError := conversion.CalculateExpr(postfixExpression)
	fmt.Printf("\nResult: %d\nError: %v\n", result, isError)
}
