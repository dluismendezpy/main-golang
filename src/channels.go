package main

import "fmt"

func printMessage(msj string, channel chan<- string) {
	channel <- msj
}

func main() {
	channel := make(chan string, 1)

	fmt.Println("Hello")

	go printMessage("World", channel)

	fmt.Println(<-channel)
}
