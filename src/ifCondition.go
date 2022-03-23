package main

import (
	"fmt"
	"log"
	"strconv"
)

func convertValue(valueS string) {
	value, err := strconv.Atoi(valueS)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d %T", value, value)
}

func main() {
	var value1 int8 = 1
	var value2 int8 = 2
	var result bool

	if value1 != value2 {
		result = true
		fmt.Println(result)
	}

	if value1 != value2 && value1 == value2 {
		result = true
		fmt.Println(result)
	} else {
		result = false
		fmt.Println(result)
	}

	if value1 != value2 || value1 == value2 {
		result = true
		fmt.Println(result)
	} else {
		result = false
		fmt.Println(result)
	}

	convertValue("8")
}
