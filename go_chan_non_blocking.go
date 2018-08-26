package main

import (
	"fmt"
	)

func main() {
	messgaes := make(chan string)
	signals := make(chan bool)
	//go func() {messgaes <- "cfd"}()
	//go func() {signals <- true }()

	select {
	case msg := <-messgaes:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messgaes <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messgaes:
		fmt.Println("received message", msg)
	case sig := <- signals:
		fmt.Println("recevied signal", sig)
	default:
		fmt.Println("no activity")
	}

}
