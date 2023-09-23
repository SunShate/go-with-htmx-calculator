package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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

	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, strings.Join(arithmeticExpression, ""))
		io.WriteString(w, fmt.Sprintf("\nExpression result: %d\n", result))
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
