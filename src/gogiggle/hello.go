package main

import "fmt"

func printer(input chan string) {
	msg := <-input
	fmt.Println(msg)
}

func hello() {
	var messages = make(chan string)
	go printer(messages)

	messages <- "hello world"
}
