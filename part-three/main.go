package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func(){
		c <- 42
		c <- 43
		c <- 44
		}()
		
	v, ok := <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)

	close(c)
	
	v, ok = <-c
	fmt.Println(v, ok)
}
