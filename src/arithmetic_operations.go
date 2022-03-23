package main

import "fmt"

func main() {
	var a float64 = 20
	var b float64 = 20

	result_s := a + b
	result_r := a - b
	resut_m := a * b
	result_d := a / b
	resut_mo := 10 % 10

	fmt.Println(result_s, result_r, resut_m, result_d, resut_mo)
}
