package main

import "fmt"

type figure interface {
	area() float64
}

type square struct {
	base float64
}

type triangle struct {
	base   float64
	height float64
}

func (mySquare square) area() float64 {
	return mySquare.base * mySquare.base
}

func (myTriangle triangle) area() float64 {
	return myTriangle.base * myTriangle.height / 2
}

func calculate(myFigure figure) float64 {
	return myFigure.area()
}

func main() {
	var mySquare square = square{base: 14.2}
	var myTriangle triangle = triangle{base: 63.2, height: 9}

	fmt.Printf("%f\n", calculate(mySquare))
	fmt.Printf("%f\n", calculate(myTriangle))
}
