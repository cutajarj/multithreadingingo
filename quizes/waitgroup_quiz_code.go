package main

import (
	"fmt"
	"sync"
)

func WaitGroupQuesA() {
	wg := sync.WaitGroup{}
	wg.Wait()
	fmt.Println("Done")
}

func Count() {
	wg := sync.WaitGroup{}
	x := 0
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go increment(&x, &wg)
	}
	wg.Wait()
	fmt.Printf("%d\n", x)
}

func increment(x *int, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		*x += 1
	}
	wg.Done()
}
