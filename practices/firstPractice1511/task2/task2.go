/*
Задание 2 

Считайте с консоли два вещественных числа и строку-оператор (+, -, *, /). 
Через switch выполните действие. Для деления добавьте проверку на ноль через if. 
Выведите результат через fmt.Println.
*/

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