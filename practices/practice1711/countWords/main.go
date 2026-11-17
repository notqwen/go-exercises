package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Предложение: ")
	scanner.Scan()
	
	str := scanner.Text()
	
	words := len(strings.Fields(str))
	fmt.Print("Количество слов ", words)
}