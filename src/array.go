package main

import "fmt"

func main() {
	var numbers [3]int
	numbers[0], numbers[1], numbers[2] = 1, 2, 3
	fmt.Printf("%v %d %d\n", numbers, len(numbers), cap(numbers))

	numbers2 := []int{1, 2, 3}
	fmt.Printf("%v %d %d\n", numbers2, len(numbers2), cap(numbers2))

	numbers2 = append(numbers2, 4)
	fmt.Println(numbers2)

	newArray := []int{5, 6, 7}
	numbers2 = append(numbers2, newArray...)
	fmt.Printf("%v %d %d\n", numbers2, len(numbers2), cap(numbers2))
}
