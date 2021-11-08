package channels

import (
	"fmt"
)

func SimpleSelect() {
	c1 := make(chan int)

	go func() {
		defer close(c1)
		for i := 0; i < 10; i++ {
			c1 <- i
		}
	}()

	for i := 0; i < 1000; i++ {

		select {
		case <-c1:
			fmt.Printf("c1: %v\n", <-c1)
		default:
			fmt.Println("default")
		}
	}
}

func AnotherSelect() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}
