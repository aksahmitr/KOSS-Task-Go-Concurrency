package main

import (
	"fmt"
	"time"
)

func first() {
	for i := 0; i < 10; i++ {
		fmt.Println("first")
		time.Sleep(time.Millisecond * 50)
	}
}

func second() {
	for i := 0; i < 10; i++ {
		fmt.Println("second")
		time.Sleep(time.Millisecond * 50)
	}
}

func third() {
	for i := 0; i < 10; i++ {
		fmt.Println("third")
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	go first()
	go second()
	go third()
	time.Sleep(time.Second)
}
