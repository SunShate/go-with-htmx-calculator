package operation

import (
	"fmt"
)

type operation struct {
	firstOperand  float64
	secondOperand float64
	_type         uint8
}

func New(firstOperand float64, secondOperand float64, _operation uint8) operation {
	op := operation{firstOperand, secondOperand, _operation}
	return op
}

func Compute(e operation) float64 {
	var result float64

	switch e._type {
	case 37:
		first := int64(e.firstOperand)
		second := int64(e.secondOperand)
		result = float64(first % second)
	case 43:
		result = e.firstOperand + e.secondOperand
	case 42:
		result = e.firstOperand * e.secondOperand
	case 45:
		result = e.firstOperand - e.secondOperand
	case 47:
		result = e.firstOperand / e.secondOperand
	default:
		fmt.Printf("Error: Unsupported operation! %v", string(e._type))
	}

	return result
}

// Hello returns a greeting for the named person.
func HelloMom(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, Mom. I'm your son %v!", name)
	return message
}
