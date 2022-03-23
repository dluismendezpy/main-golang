package main

import "fmt"

func main() {
	values := make(map[string]int)
	values["Python"] = 100
	values["Golang"] = 99
	values["Java"] = 98

	fmt.Println(values)

	for i, value := range values {
		fmt.Printf("%v %v\n", i, value)
	}
}
