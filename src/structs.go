package main

import (
	"fmt"

	"./thepackage"
)

func main() {
	var myCar thepackage.Car = thepackage.Car{Brand: "Ferrari", Year: 2020}
	fmt.Println(myCar)
}
