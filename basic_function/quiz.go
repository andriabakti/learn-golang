package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println("Sum >", total)

	result, err := calculate(5, 7, "=")
	if (err != nil) {
		fmt.Println("Terjadi kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println("Calc >", result)
}

func sum(arr []int) int {
	var result int
	for _, item := range arr {
		result += item
	}
	return result
}

func calculate(first, second int, operator string) (int, error) {
	var res int
	var err error

	switch operator {
	case "+":
		res = first + second
	case "-":
		res = first - second
	case "*":
		res = first * second
	case "/":
		res = first / second
	default:
		err = errors.New("unknown operator")
	}
	return res, err
}