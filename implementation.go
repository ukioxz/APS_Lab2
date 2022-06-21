package lab2

import (
  "fmt"
	"strconv"
  "strings"
	"math"
)

func PostfixCalculation(input string) (string, error) {
  stack := []int{}
	result := 0
  var elements = strings.Fields(input)

  for _, el := range elements {
    switch el {
    case "+", "-", "*", "/", "^":
      topIdx := len(stack) -1
      b := stack[topIdx]
      a := stack[topIdx-1]
      stack = stack[0: topIdx-1]
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
        result = int(math.Pow(float64(a), float64(b)))
      }
    default:
      var err error
      result, err = strconv.Atoi(el)
      if err != nil {
        return "0", fmt.Errorf("invalid input")
      }
    }
    stack = append(stack, result)
  }
  if len(stack) != 1 {
    return "0", fmt.Errorf("empty string")
  }
	res:= strconv.Itoa(result);
  return res, nil
}