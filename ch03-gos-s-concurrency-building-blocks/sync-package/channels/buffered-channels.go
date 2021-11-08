package channels

import (
	"fmt"
	"time"
)

func BufferedChannels() {
	intStream := make(chan int, 8)

	go func() {
		defer close(intStream)
		for i := 0; i < 16; i++ {
			intStream <- i

			fmt.Printf("%v streamed\n", i)
		}
	}()

	for c := range intStream {
		fmt.Printf("%v readed\n", c)
		time.Sleep(time.Second * 1)
	}
}
