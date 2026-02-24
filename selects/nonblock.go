package main

import "fmt"

func nonblock() {
	channel := make(chan int)
	go sendValue(1, channel)

	select {
	case value := <-channel:
		fmt.Printf("received from chan : %d\n", value)
	default:
		fmt.Println("channel is not yet prepared (unblock)")
	}

	fmt.Println("program ends")
}
