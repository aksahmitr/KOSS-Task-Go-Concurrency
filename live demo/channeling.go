package main

import (
	"fmt"
	"time"
)

func speaker(input chan string) {
	time.Sleep(time.Second * 3)
	input <- "Hello!"
}

func main() {

	messages := make(chan string)

	go speaker(messages)

	fmt.Println(<-messages)
}
