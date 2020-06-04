package main

import (
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
		trains[i] = &Train{id: i, back: 0, front: trainLength}
	}

	for i := 0; i < 4; i++ {
		intersections[i] = &Intersection{id: i, mutex: sync.Mutex{}, lockedBy: -1}
	}

	go moveTrain(0, 300, []*Crossing{{position: 125, intersection: intersections[0]}, {position: 175, intersection: intersections[1]}})

	go moveTrain(1, 300, []*Crossing{{position: 125, intersection: intersections[1]}, {position: 175, intersection: intersections[2]}})

	go moveTrain(2, 300, []*Crossing{{position: 125, intersection: intersections[2]}, {position: 175, intersection: intersections[3]}})

	go moveTrain(3, 300, []*Crossing{{position: 125, intersection: intersections[3]}, {position: 175, intersection: intersections[0]}})

	if err := ebiten.Run(update, 320, 320, 3, "Trains in a box"); err != nil {
		log.Fatal(err)
	}
}
