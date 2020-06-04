package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math"
)

var (
	green = color.RGBA{10, 125, 50, 255}
	red   = color.RGBA{225, 50, 50, 255}
)

func DrawTracks(screen *ebiten.Image) {
	for i := 0; i < 300; i++ {
		screen.Set(10+i, 135, green)
		screen.Set(185, 10+i, green)
		screen.Set(310-i, 185, green)
		screen.Set(135, 310-i, green)
	}
}

func DrawXTrain(screen *ebiten.Image, id int, dir int, start int, yPos int) {
	s := start + (dir * trains[id].back)
	e := start + (dir * trains[id].front)
	for i := math.Min(float64(s), float64(e)); i <= math.Max(float64(s), float64(e)); i++ {
		screen.Set(int(i)-dir, yPos-1, red)
		screen.Set(int(i), yPos, red)
		screen.Set(int(i)-dir, yPos+1, red)
	}
}

func DrawYTrain(screen *ebiten.Image, id int, dir int, start int, xPos int) {
	s := start + (dir * trains[id].back)
	e := start + (dir * trains[id].front)
	for i := math.Min(float64(s), float64(e)); i <= math.Max(float64(s), float64(e)); i++ {
		screen.Set(xPos-1, int(i)-dir, red)
		screen.Set(xPos, int(i), red)
		screen.Set(xPos+1, int(i)-dir, red)
	}
}
