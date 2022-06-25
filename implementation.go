package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func PostfixCalculation(input string) (string, error) {
	stack := []float64{}
	var result float64 = 0
	var elements = strings.Fields(input)

	for _, el := range elements {
		switch el {
		case "+", "-", "*", "/", "^":
			id := len(stack) - 1
			b := stack[id]
			a := stack[id-1]
			stack = stack[0 : id-1]
			switch el {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			case "^":
				result = math.Pow(a, b)
			}
		default:
			var err error
			result, err = strconv.ParseFloat(el, 64)
			if err != nil {
				return "0", fmt.Errorf("invalid input")
			}
		}
		stack = append(stack, result)
	}
	if len(stack) != 1 {
		return "0", fmt.Errorf("syntax error")
	}
	res := fmt.Sprint(result)
	return res, nil
}
