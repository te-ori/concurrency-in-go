package main

import (
	"fmt"
	"sync"
)

func OnceDemo() {
	var count int

	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}

func OnceWithDifferentFuncs() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	// `Once`'ın kendisi fonksiyon bazında bir takip sağlamıyor! Aslında `Once`
	// kendisinin `Do` metodunun çağrılıp çağrılmadığını kontrol ediyor. Bu
	// nedenle bir `Once` instance'ının birden fazla çağrılmasının hiç bir etkisi
	// yoktur.
	//
	//
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count)
}

// @1. `onceA` `initA` fonksiyonu ile çağrılıyor. Bu noktada `onceA` kendisinin
//     tek sefer çalışmasını garanti etmek için bir tür `lock` başlatamkata.
// @2. `initA` fonksiyonu çalıştığında `onceB` çalıştırılmakta. `onceA` ve
//     `onceB` iki ayrı `once` olduğu için `onceB` de bir defaya mahsus olmak
//     üzere `initB` fonksiyonunu çalıştıracak
// @3. `onceA` tekrar `initA` fonksiyonu ile çalıştırılmak isteniyor. Bunun
//     içinde -muhtemelen- önce bir `lock` başlatıyor. Fakat `lock` @1 adımındaki
//     `lock` ile çakışıyor. @1'deki `lock` kapanmak için çağırdığı `initA`'nın
//     tamamlanmasını bekliyor. @3'de ki `Do` çağrısı da devam etmek için @1'deki
// 	   `lock`'un açılmasını bekliyor
// Böylece Deadlock oluşmuş oluyor.
func OnceDeadLock() {

	var onceA, onceB sync.Once
	var initB func()

	initA := func() {
		fmt.Println("A")
		onceB.Do(initB) // @2
	}

	initB = func() {
		fmt.Println("B")
		onceA.Do(initA) // @3
	}

	onceA.Do(initA) // @1
}
