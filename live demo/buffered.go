package main

import "fmt"

func main() {

	messages := make(chan string, 3)

	messages <- "First!"
	messages <- "Second!"
	messages <- "Third!"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
