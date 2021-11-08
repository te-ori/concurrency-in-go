package channels

import (
	"bytes"
	"fmt"
	"os"
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

// ??? This is an example of an optimization that can be useful under the right
// conditions: if a goroutine making writes to a channel has knowledge of how
// many writes it will make, it can be useful to create a buffered channel whose
// capacity is the number of writes to be made, and then make those writes
// as quickly as possible. There are, of course, caveats, and we’ll cover them
// in the next chapter.
func AnotherSample() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Procedure done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}
