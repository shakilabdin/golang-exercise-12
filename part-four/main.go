package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	fmt.Println("go routines:", runtime.NumGoroutine())

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(n int) {
				for i := 0; i < 10; i++ {
					c <- i*10 + n
				}
				wg.Done()
			}(i)
		}
		fmt.Println("go routines:", runtime.NumGoroutine())
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println("printing", v)
	}

	fmt.Println("END")

}
