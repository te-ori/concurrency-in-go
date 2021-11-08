package main

import (
	"fmt"
	"sync"
)

func Sample() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new isnstacne.")
			return struct{}{}
		},
	}

	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
}

func AnotherSample() {
	var numCalcCreate int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcCreate += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()

			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(func(m interface{}) interface{} {
				fmt.Printf("%+v\n", m)
				return m
			}(mem))
		}()
	}

	wg.Wait()
	fmt.Printf("%+v calculators were created.\n", numCalcCreate)
}
