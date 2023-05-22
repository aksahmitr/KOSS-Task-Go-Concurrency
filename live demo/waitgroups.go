package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(s string, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println(s)
	time.Sleep(time.Second)
	fmt.Println("OVER!")
}

func main() {
	var w sync.WaitGroup
	for i := 0; i < 5; i++ {
		w.Add(1)
		go worker("Hello!", &w)
	}
	w.Wait()
}
