package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type spinLock int32

func (sl *spinLock) Lock() {
	for !atomic.CompareAndSwapInt32((*int32)(sl), 0, 1) {
		runtime.Gosched()
	}
}

func (sl *spinLock) Unlock() {
	atomic.StoreInt32((*int32)(sl), 0)
}

func NewSpinLock() sync.Locker {
	var lock spinLock
	return &lock
}
