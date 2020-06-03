package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
	"sync"
)

var (
	allTrains [4]*Train
	green     = color.RGBA{10, 95, 30, 255}
	red       = color.RGBA{255, 50, 50, 255}
)

const trainLength = 60

func update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		for i := 0; i < 300; i++ {
			screen.Set(10+i, 135, green)
			screen.Set(185, 10+i, green)
			screen.Set(310-i, 185, green)
			screen.Set(135, 310-i, green)
		}
		screen.Set(10+allTrains[0].back, 135, red)
		screen.Set(10+allTrains[0].front, 135, red)
		screen.Set(185, 10+allTrains[1].back, red)
		screen.Set(185, 10+allTrains[1].front, red)
		screen.Set(310-allTrains[2].back, 185, red)
		screen.Set(310-allTrains[2].front, 185, red)
		screen.Set(135, 310-allTrains[1].back, red)
		screen.Set(135, 310-allTrains[1].front, red)
	}
	return nil
}

func main() {
	it1, it2, it3, it4 := sync.Mutex{}, sync.Mutex{}, sync.Mutex{}, sync.Mutex{}

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
