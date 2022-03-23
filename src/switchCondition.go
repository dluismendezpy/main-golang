package main

import "fmt"

func main() {
	var result bool
	switch module := 4 % 2; module {
	case 0:
		result = true
	default:
		result = false
	}
	fmt.Println(result)
}
