/*
Задание 3

Считайте с консоли одно целое число. 
Если оно делится на 3 и на 5 — выведите "FizzBuzz", 
если только на 3 — "Fizz", только на 5 — "Buzz", 
иначе выведите само число. Используйте % и if / else if / else.
*/

package main

import "fmt"

func main(){
	var num int
	fmt.Scan(&num)

	if num % 3 == 0 && num % 5 == 0 {
		fmt.Print("FizzBuzz")
	} else if num % 3 == 0 {
		fmt.Print("Fizz")
	} else if num % 5 == 0 {
		fmt.Print("Buzz")
	} else {
		print(num)
	}
}