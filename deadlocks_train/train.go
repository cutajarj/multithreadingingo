package main

import (
	"sync"
	"time"
)

type Crossing struct {
	position     int
	intersection sync.Mutex
}

type Train struct {
	back  int
	front int
}

func moveTrain(length int, distance int, crossings []*Crossing, train *Train) {
	train.back = 0
	train.front = length
	for train.front < distance {
		train.back += 1
		train.front += 1
		for _, crossing := range crossings {
			if train.front == crossing.position {
				crossings[0].intersection.Lock()
			}
			if train.back == crossing.position {
				crossings[0].intersection.Unlock()
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}
