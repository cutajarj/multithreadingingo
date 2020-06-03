package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock1 = sync.Mutex{}
	lock2 = sync.Mutex{}
)

func blueRobot() {
	for {
		fmt.Println("Blue: Acquiring first lock...")
		lock1.Lock()
		fmt.Println("Blue: Acquiring second lock...")
		lock2.Lock()
		fmt.Println("Blue: Locks Acquired")
		lock2.Unlock()
		lock1.Unlock()
	}
}

func redRobot() {
	for {
		fmt.Println("Red: Acquiring first lock...")
		lock2.Lock()
		fmt.Println("Red: Acquiring second lock...")
		lock1.Lock()
		fmt.Println("Red: Locks Acquired")
		lock1.Unlock()
		lock2.Unlock()
	}
}

func main() {
	go redRobot()
	go blueRobot()
	time.Sleep(20 * time.Second)
	fmt.Println("Done")
}
