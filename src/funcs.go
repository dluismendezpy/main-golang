package main

import "fmt"

func printMessage(msj string) string {
	var result string = fmt.Sprintf("%s!", msj)
	return result
}

func printValue(number int8) int8 {
	return number
}

func makeOperations(num1 int16, num2 int16) (a, b int16) {
	var result_s int16 = num1 + num2
	var result_m int16 = num1 * num2

	return result_s, result_m
}

func main() {
	var value1, value2 int16 = makeOperations(5, 5)

	fmt.Println(printMessage("Hello World"))
	fmt.Println(printValue(5))
	fmt.Printf("%d %d", value1, value2)
}
