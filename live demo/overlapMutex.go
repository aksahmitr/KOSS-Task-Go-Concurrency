package main

import (
	"fmt"
	"sync"
)

func main() {

	cnt := 0

	var wg sync.WaitGroup

	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				m.Lock()
				cnt++
				m.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(cnt)
}
