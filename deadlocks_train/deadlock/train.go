package deadlock

import (
	. "github.com/cutajarj/multithreadingingo/deadlocks_train/common"
	"time"
)

func MoveTrain(train *Train, distance int, crossings []*Crossing) {
	//time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	for train.Front < distance {
		train.Front += 1
		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				crossing.Intersection.Mutex.Lock()
				crossing.Intersection.LockedBy = train.Id
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
