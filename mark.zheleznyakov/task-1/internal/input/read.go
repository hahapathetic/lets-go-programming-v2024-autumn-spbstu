package input

import (
	"fmt"
	"strconv"
	"github.com/mrqiz/task-1/internal/strings"
)

func readNumber(zeroAllowed bool) float64 {
  var input string
	var result float64

	for {
		fmt.Printf("gimme a number: ")
		fmt.Scanln(&input)

		n, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("this is not a number")
			continue
		}

		result = n

		if !zeroAllowed && result == 0 {
			fmt.Println("no zero is allowed")
			continue
		}

		break	
	}

	return result
}

func readOperator() rune {
	var operator string

	for {
		fmt.Printf("gimme an operator: ")
		fmt.Scanln(&operator)
		
		allowedOperators := []string{"+", "-", "*", "/"}

		if !strings.Has(allowedOperators, operator) {
			fmt.Println("this is not an operator, ok?")
			continue
		}

		break
	}

	return []rune(operator)[0]
}

func Read() (float64, float64, rune) {
	lOperand := readNumber(true)
	operator := readOperator()
	rOperand := readNumber(string(operator) != "/")

	return lOperand, rOperand, operator
}
