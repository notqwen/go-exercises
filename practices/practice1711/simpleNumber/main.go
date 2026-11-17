package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	fmt.Scan(&number)

	if number < 2 {
		fmt.Println("Введено неверное число")
		os.Exit(1)
	}

	var isNotSimple bool
	for i := 2; i < number; i++ {
		if number%i == 0 {
			isNotSimple = true
		}
	}

	if isNotSimple {
		fmt.Println("Составное")
	} else {
		fmt.Println("Простое")
	}
}
