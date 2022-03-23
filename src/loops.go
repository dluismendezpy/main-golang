package main

import "fmt"

func main() {
	var counter int32 = 0
	var counterF int32 = 0
	var i int16

	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	for {
		fmt.Println(counterF)
		counterF++
	}
}
