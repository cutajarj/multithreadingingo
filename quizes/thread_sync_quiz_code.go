package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock   = sync.Mutex{}
	rwLock = sync.RWMutex{}
)

func oneTwoThreeA() {
	lock.Lock()
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Millisecond)
	}
	lock.Unlock()
}

func StartThreadsA() {
	for i := 1; i <= 2; i++ {
		go oneTwoThreeA()
	}
	time.Sleep(1 * time.Second)
}

func oneTwoThreeB() {
	rwLock.RLock()
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Millisecond)
	}
	rwLock.RLock()
}

func StartThreadsB() {
	for i := 1; i <= 2; i++ {
		go oneTwoThreeB()
	}
	time.Sleep(1 * time.Second)
}
