package common

import "sync"

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
