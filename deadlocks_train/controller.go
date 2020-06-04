package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
	"sync"
)

var (
	allTrains          [4]*Train
	it1, it2, it3, it4 = sync.Mutex{}, sync.Mutex{}, sync.Mutex{}, sync.Mutex{}
)

const trainLength = 70

func update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		DrawTracks(screen)
		DrawXTrain(screen, 0, 1, 10, 135)
		DrawYTrain(screen, 1, 1, 10, 185)
		DrawXTrain(screen, 2, -1, 310, 185)
		DrawYTrain(screen, 3, -1, 310, 135)
	}
	return nil
}

func main() {
	for i := 0; i < 4; i++ {
		allTrains[i] = &Train{id: i, back: 0, front: trainLength}
	}

	go moveTrain(300,
		[]*Crossing{{position: 125, intersection: &it1}, {position: 175, intersection: &it2}},
		allTrains[0])

	go moveTrain(300,
		[]*Crossing{{position: 125, intersection: &it2}, {position: 175, intersection: &it3}},
		allTrains[1])

	go moveTrain(300,
		[]*Crossing{{position: 125, intersection: &it3}, {position: 175, intersection: &it4}},
		allTrains[2])

	go moveTrain(300,
		[]*Crossing{{position: 125, intersection: &it4}, {position: 175, intersection: &it1}},
		allTrains[3])

	if err := ebiten.Run(update, 320, 320, 2, "Trains in a box"); err != nil {
		log.Fatal(err)
	}
}
