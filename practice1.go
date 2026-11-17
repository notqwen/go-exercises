package main

import "fmt"

func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				fmt.Print(string(str1[i]), string(' '))
				break
			}
		}
	}
}