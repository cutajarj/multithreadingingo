package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mlock = sync.Mutex{}
	cond  = sync.NewCond(&mlock)
)

func runChildThread() {
	mlock.Lock()
	fmt.Println("RunChildThread, lock acquired")
	cond.Signal()
	fmt.Println("RunChildThread, Waiting")
	cond.Wait()
	fmt.Println("RunChildThread, Running")
}

func RunMainThread() {
	mlock.Lock()
	fmt.Println("RunMainThread, lock acquired")
	go runChildThread()
	fmt.Println("RunMainThread, Waiting")
	cond.Wait()
	fmt.Println("RunMainThread, Running")
	cond.Signal()
	time.Sleep(10 * time.Second)
}
