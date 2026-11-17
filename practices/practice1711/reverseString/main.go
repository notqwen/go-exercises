package main

import "fmt"

func main(){
	var s, resultS string
	fmt.Scan(&s)

	for i := len(s) - 1; i >= 0; i-- {
		resultS += string(s[i])
	}
	fmt.Println(resultS)
}