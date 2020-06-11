package arbitrator

import (
	. "github.com/cutajarj/multithreadingingo/deadlocks_train/common"
	"sync"
	"time"
)

var (
	Waiter = sync.Mutex{}
	Cond   = sync.NewCond(&Waiter)
)

func allFree(intersectionsToLock []*Intersection) bool {
	for _, it := range intersectionsToLock {
		if it.LockedBy >= 0 {
			return false
		}
	}
	return true
}

func lockIntersectionsInDistance(id, reserveStart int, reserveEnd int, crossings []*Crossing) {
	var intersectionsToLock []*Intersection
	for _, crossing := range crossings {
		if reserveEnd >= crossing.Position &&
			reserveStart <= crossing.Position &&
			crossing.Intersection.LockedBy != id {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}
	Waiter.Lock()
	for !allFree(intersectionsToLock) {
		Cond.Wait()
	}
	for _, it := range intersectionsToLock {
		it.LockedBy = id
		time.Sleep(10 * time.Millisecond)
	}
	Waiter.Unlock()
}

func MoveTrain(train *Train, distance int, crossings []*Crossing) {
	//time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	for train.Front < distance {
		train.Front += 1
		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				lockIntersectionsInDistance(train.Id, crossing.Position, crossing.Position+train.TrainLength, crossings)
			}
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				Waiter.Lock()
				crossing.Intersection.LockedBy = -1
				Cond.Broadcast()
				Waiter.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
