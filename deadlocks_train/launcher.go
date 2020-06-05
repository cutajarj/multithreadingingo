package main

import (
	. "github.com/cutajarj/multithreadingingo/deadlocks_train/arbitrator"
	"github.com/hajimehoshi/ebiten"
	"log"
	"sync"
)

const trainLength = 70

func update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		DrawTracks(screen)
		DrawIntersections(screen)
		DrawTrains(screen)
	}
	return nil
}

func main() {
	for i := 0; i < 4; i++ {
		Trains[i] = &Train{Id: i, TrainLength: trainLength, Front: 0}
	}

	for i := 0; i < 4; i++ {
		Intersections[i] = &Intersection{Id: i, Mutex: sync.Mutex{}, LockedBy: -1}
	}

	go MoveTrain(0, 300, []*Crossing{{Position: 125, Intersection: Intersections[0]}, {Position: 175, Intersection: Intersections[1]}})

	go MoveTrain(1, 300, []*Crossing{{Position: 125, Intersection: Intersections[1]}, {Position: 175, Intersection: Intersections[2]}})

	go MoveTrain(2, 300, []*Crossing{{Position: 125, Intersection: Intersections[2]}, {Position: 175, Intersection: Intersections[3]}})

	go MoveTrain(3, 300, []*Crossing{{Position: 125, Intersection: Intersections[3]}, {Position: 175, Intersection: Intersections[0]}})

	if err := ebiten.Run(update, 320, 320, 3, "Trains in a box"); err != nil {
		log.Fatal(err)
	}
}
