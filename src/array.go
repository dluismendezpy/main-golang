package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) bool {
	var textLower string = strings.ToLower(text)
	var isPalindrome bool = false
	var textReserse string

	for i := len(textLower) - 1; i >= 0; i-- {
		textReserse += string(textLower[i])
	}

	if textReserse == textLower {
		isPalindrome = true
	}

	return isPalindrome
}

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

	messages := []string{"Python", "Golang", "JavaScript"}
	for i, value := range messages {
		fmt.Printf("%d %s\n", i, value)
	}

	fmt.Printf("%s is palindrome? %t", "amor", isPalindrome("Ama"))
}
