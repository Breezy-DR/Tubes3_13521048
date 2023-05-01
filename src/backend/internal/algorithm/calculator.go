package algorithm

import (
	"errors"
	"math"
	"strconv"
)

func CalculateExpression(expression string) (float64, error) {
	var (
		numbers []float64
		ops     []uint8
	)

	expressionLen := len(expression)
	for i := 0; i < expressionLen; i++ {
		if expression[i] == ' ' {
			continue
		}

		el := expression[i]

		if el >= '0' && el <= '9' {
			counter := i
			for i < expressionLen && (expression[i] >= '0' && expression[i] <= '9') {
				i++
			}
			res, _ := strconv.ParseFloat(expression[counter:i], 32)
			numbers = append(numbers, res)
			i -= 1
		} else if el == '(' {
			ops = append(ops, el)
		} else if el == ')' {
			m := len(numbers) - 1 // Top element of numbers
			n := len(ops) - 1     // Top element of ops
			for m >= 1 && n >= 0 && ops[n] != '(' {
				right := numbers[m]
				left := numbers[m-1]
				numbers = numbers[:(m - 1)]

				op := ops[n]
				ops = ops[:n]

				calc, err := basicCalc(left, right, op)

				if err != nil {
					return -1, err
				}
				numbers = append(numbers, calc)
				m = len(numbers) - 1
				n = len(ops) - 1
			}

			if n >= 0 && ops[n] == '(' {
				ops = ops[0:n]
			} else {
				return -1, errors.New("brackets do not match")
			}

		} else {
			m := len(numbers) - 1 // Top element of numbers
			n := len(ops) - 1     // Top element of ops
			for n >= 0 && getPrec(ops[n]) >= getPrec(el) {
				right := numbers[m]
				left := numbers[m-1]
				numbers = numbers[:(m - 1)]

				op := ops[n]
				ops = ops[:n]

				calc, err := basicCalc(left, right, op)

				if err != nil {
					return -1, err
				}
				numbers = append(numbers, calc)
				m = len(numbers) - 1
				n = len(ops) - 1
			}
			ops = append(ops, el)
		}

	}

	m := len(numbers) - 1 // Top element of numbers
	n := len(ops) - 1     // Top element of ops
	for n >= 0 && m >= 1 {
		right := numbers[m]
		left := numbers[m-1]
		numbers = numbers[:(m - 1)]

		op := ops[n]
		ops = ops[:n]

		calc, err := basicCalc(left, right, op)

		if err != nil {
			return -1, err
		}
		numbers = append(numbers, calc)

		m = len(numbers) - 1
		n = len(ops) - 1
	}

	if len(ops) > 0 || len(numbers) != 1 {
		return 0, errors.New("invalid format of expression")
	}

	return numbers[len(numbers)-1], nil
}

func getPrec(op uint8) int {
	if op == '^' {
		return 3
	}
	if op == '*' || op == '/' {
		return 2
	}
	if op == '+' || op == '-' {
		return 1
	}
	return 0
}

func basicCalc(left, right float64, op uint8) (float64, error) {
	if op == '^' {
		return math.Pow(left, right), nil
	}
	if op == '*' {
		return left * right, nil
	}
	if op == '/' {
		return left / right, nil
	}
	if op == '-' {
		return left - right, nil
	}
	if op == '+' {
		return left + right, nil
	}

	// we use error because we couldn't simply return -1 since there might be an actual result of -1
	return -1, errors.New("invalid operator")
}
