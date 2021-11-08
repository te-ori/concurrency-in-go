package channels

import "fmt"

func SimpleChannel() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()

	fmt.Println(<-stringStream)
}
