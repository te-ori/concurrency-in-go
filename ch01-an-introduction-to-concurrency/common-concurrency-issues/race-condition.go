package main

import (
	"fmt"
	"sync"
	"time"
)

// *Race Condition* belirli bir sırayla çalışması zorunlu olan operasyonların
// bu sırayla çalışmasının garanti altına alınmadığı durumlarda ortaya çıkar
// [ch01]
func race_condition() {
	stepCount := 100

	data := 0
	go func() {
		for data < stepCount {
			// Paralel çalışma ihtimali olan bir `goroutine` içerisinde `data`
			// nın değeri değiştiriliyor
			data++

			time.Sleep(200 * time.Nanosecond)
		}
	}()

	for i := 0; i < stepCount; i++ {
		// Paralelde değeri değiştirlen `data` değişkeninin değeri burada okunmaya
		// çalışılıyor. Yukarıdaki `goroutine` ile bu kodun paralel çalışma ihtimali
		// var. Eğer paralel çalışma gerçekleşirse değişkenin değerinin ne zaman
		// değiştirileceği ve ne zaman okunacağı işletim sisteminin inisiyatifine
		// kalmış ve bu çalıştırma işleminin sırası her çalıştırmada farklılık
		// göstrebilir.
		fmt.Printf("the data is %v.\n", data)
		time.Sleep(20 * time.Nanosecond)
	}
}

// En iyi çözüm değil!
func solve_race_condition_with_mutex() {
	var memoryAccess sync.Mutex
	data := 0
	stepCount := 100

	go func() {
		for data < stepCount {
			memoryAccess.Lock()
			data++
			time.Sleep(200 * time.Nanosecond)
			memoryAccess.Unlock()
		}
	}()

	for i := 0; i < stepCount; i++ {
		memoryAccess.Lock()
		fmt.Printf("the data is %v.\n", data)
		time.Sleep(20 * time.Nanosecond)
		memoryAccess.Unlock()
	}
}
