/*
Задание 1.

Объявите четыре переменные через var с явным указанием типа: 
целочисленную a = 10, вещественную b = 3.14, строковую c = "Go" 
и логическую d = true. Выведите все четыре значения через fmt.Println 
в одной строке.
*/

package main

import "fmt"

func main(){
	var a int = 10
	var b float32 = 3.14
	var c string = "Go"
	var d bool = true
	fmt.Println(a, b, c, d)
}