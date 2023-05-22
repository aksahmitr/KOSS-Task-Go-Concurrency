package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var cnt int32

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddInt32(&cnt, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(cnt)
}
