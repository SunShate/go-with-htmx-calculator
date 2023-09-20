package main

import (
	"fmt"
	"os"

	"calculator/operation"
)

func main() {
	fmt.Println("Interface will add sooner")
	fmt.Println(operation.HelloMom(os.Args[1]))

	_operation := operation.New(5.0, 2.1, '%')
	fmt.Println(operation.Compute(_operation))
}
