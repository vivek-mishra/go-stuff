package main

import "fmt"

func main() {
//	c := make(chan string, 1)
	done := make(chan string, 2)
	fmt.Println(cap(done))
	go func() {
		//c <- "Hello"
	//	c <- "World"

	done <- "done"
		done <- "deal"
		close(done)
	}()
	for i := 0; i < cap(done); i++ {
		fmt.Println(<-done)
	}
}