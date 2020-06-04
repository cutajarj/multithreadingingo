package main

import (
	"sort"
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

func lockIntersectionsInDistance(id, pos int, crossings []*Crossing) {
	var intersectionsToLock []*Intersection
	for _, crossing := range crossings {
		if pos > crossing.position-trainLength-1 && crossing.intersection.lockedBy != id {
			intersectionsToLock = append(intersectionsToLock, crossing.intersection)
		}
	}

	sort.Slice(intersectionsToLock, func(i, j int) bool {
		return intersectionsToLock[i].id < intersectionsToLock[j].id
	})

	for _, it := range intersectionsToLock {
		it.mutex.Lock()
		it.lockedBy = id
		time.Sleep(10 * time.Millisecond)
	}

}

func moveTrain(id int, distance int, crossings []*Crossing) {
	//time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	train := trains[id]
	for train.front < distance {
		train.back += 1
		train.front += 1
		for _, crossing := range crossings {
			if train.front == crossing.position-1 {
				lockIntersectionsInDistance(id, train.front, crossings)
			}
			if train.back == crossing.position+1 {
				crossing.intersection.mutex.Unlock()
				crossing.intersection.lockedBy = -1
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
