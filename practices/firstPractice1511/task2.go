package main

import "fmt"

func main() {
	var num1, num2 float64
	var op string

	fmt.Scan(&num1, &op, &num2)

	switch op {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Делить на ноль нельзя")
		} else {
			fmt.Println(num1 / num2)
		}
	default:
		fmt.Println("Неизвестный оператор")
	}
}