package main

import (
	"fmt"
	"time"
)

type message struct {
	msj string
}

func (myMessage message) getMessage() {
	fmt.Printf("%s\n", myMessage.msj)
}

func main() {
	var myMsj message = message{msj: "Hello"}
	var myMsj2 message = message{msj: "World"}

	myMsj.getMessage()
	go myMsj2.getMessage()
	time.Sleep(time.Second * 1)
}
