package main

import (
	"fmt"
	"sync"
	"time"
)

type Crossing struct {
	position     int
	intersection *sync.Mutex
}

type Train struct {
	id    int
	back  int
	front int
}

func moveTrain(distance int, crossings []*Crossing, train *Train) {
	//time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	for train.front < distance {
		fmt.Println("Train", train.id, "at position", train.front)
		train.back += 1
		train.front += 1
		for _, crossing := range crossings {
			if train.front == crossing.position-1 {
				fmt.Println("Train", train.id, "locking")
				crossing.intersection.Lock()
			}
			if train.back == crossing.position+1 {
				fmt.Println("Train", train.id, "Unlocking")
				crossing.intersection.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
