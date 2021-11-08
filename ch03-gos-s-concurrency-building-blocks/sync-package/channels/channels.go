package channels

import (
	"fmt"
	"time"
)

func SimpleChannel() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()

	fmt.Println(<-stringStream)
}

// `range` bir `channel` ile kullanıldığında bir tür `iterator` oluşturmakta.
// Bu iterator:
// 1. `channel` açıkken ve `channel`'a bir değer yazılmamışsa, `channel`'a bir
//    değer yazılana kadar veya `channel`'ın kapatıldığı -`close(intStream)`-
//    sinyali gelene kadar burada bekler.
// 2. `channel`'a veri aktarıldığı sinyali gelirse döngünün içine girerek ilgili
//    komutları çalıştırır ve 1. adımdaki duruma döner
// 3. Eğer kapatma sinyali gelirse döngüyü bitirir ve döngüden hemen sonraki
//    kodları çağırır.
func ForRangeChannels() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v", integer)
	}
}

// Eğer bir `channel` kapatılmadan bir `goroutine`'den çıkılırsa
func ForRangeWithUnclosedChannels() {
	fmt.Println("ForRangeWithUnclosedChannels")
	intStream := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			intStream <- i

		}

		time.Sleep(2 * time.Second)
		fmt.Println("Go routine completed")
	}()

	for integer := range intStream {
		fmt.Printf("%v", integer)
	}

	fmt.Println("reading channel complete")
}
