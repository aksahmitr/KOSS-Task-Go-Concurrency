package main

import (
	"fmt"
	"time"
)

func first(c chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println("first")
		time.Sleep(time.Millisecond * 50)
	}
	c <- true
}

func second(c chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println("second")
		time.Sleep(time.Millisecond * 50)
	}
	c <- true
}

func third(c chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println("third")
		time.Sleep(time.Millisecond * 50)
	}
	c <- true
}

func main() {
	channel := make(chan bool)
	go first(channel)
	go second(channel)
	go third(channel)
	<-channel
	<-channel
	<-channel
}
