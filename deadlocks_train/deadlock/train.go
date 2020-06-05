package deadlock

import (
	"sync"
	"time"
)

var (
	Trains        [4]*Train
	Intersections [4]*Intersection
)

type Intersection struct {
	Id       int
	Mutex    sync.Mutex
	LockedBy int
}

type Crossing struct {
	Position     int
	Intersection *Intersection
}

type Train struct {
	Id          int
	TrainLength int
	Front       int
}

func MoveTrain(id int, distance int, crossings []*Crossing) {
	//time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	train := Trains[id]
	for train.Front < distance {
		train.Front += 1
		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				crossing.Intersection.Mutex.Lock()
				crossing.Intersection.LockedBy = id
			}
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				crossing.Intersection.Mutex.Unlock()
				crossing.Intersection.LockedBy = -1
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
