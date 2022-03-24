package main

import "fmt"

type pc struct {
	ram        int8
	totalSpace int16
	brand      string
}

func (myPc pc) getBrand() string {
	var msj string = fmt.Sprintf("%v MSI", myPc.brand)
	return msj
}

func (myPc *pc) getMoreRam() {
	myPc.ram = myPc.ram * 2
}

func main() {
	var a int8 = 10
	b := &a

	fmt.Printf("%d\n%v\n%d", a, b, *b)

	*b = 20

	fmt.Printf("\n\n%d\n%v\n%d\n", a, b, *b)

	var myPc pc = pc{ram: 16, totalSpace: 1000, brand: "Hp"}
	fmt.Println(myPc)
	fmt.Println(myPc.getBrand())

	fmt.Println("\n", myPc)
	myPc.getMoreRam()

	fmt.Println("\n", myPc)
	myPc.getMoreRam()

	fmt.Println("\n", myPc)
}
