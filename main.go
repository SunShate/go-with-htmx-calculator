package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"calculator/conversion"
)

type Expression struct {
	InfixExpr   string
	PostfixExpr string
	Result      uint64
}

func main() {
	fmt.Println("Interface will be sooner")

	arithmeticExpression := strings.Split(os.Args[1], " ")

	// convert arithmetic expression to postfix expression
	postfixExpression, isError := conversion.BuildPostfix(arithmeticExpression)

	if isError {
		return
	}

	fmt.Println("\nPostfix notation:", postfixExpression)

	// postfix calculation
	result, isError := conversion.CalculateExpr(postfixExpression)
	fmt.Printf("\nResult: %d\nError: %v\n", result, isError)

	expression := Expression{
		strings.Join(arithmeticExpression, ""),
		strings.Join(postfixExpression, ""),
		result,
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, expression)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
