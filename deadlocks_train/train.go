package main

import (
	"sync"
	"time"
)

var (
	trains        [4]*Train
	intersections [4]*Intersection
)

type Intersection struct {
	id       int
	mutex    sync.Mutex
	lockedBy int
}

type Crossing struct {
	position     int
	intersection *Intersection
}

type Train struct {
	id    int
	back  int
	front int
}

func moveTrain(id int, distance int, crossings []*Crossing) {
	//time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	train := trains[id]
	for train.front < distance {
		train.back += 1
		train.front += 1
		for _, crossing := range crossings {
			if train.front == crossing.position-1 {
				crossing.intersection.mutex.Lock()
				crossing.intersection.lockedBy = id
			}
			if train.back == crossing.position+1 {
				crossing.intersection.mutex.Unlock()
				crossing.intersection.lockedBy = -1
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
