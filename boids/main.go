package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
	"sync"
)

const (
	screenWidth, screenHeight = 640, 360
	boidCount                 = 500
	viewRadius                = 13
	adjRate                   = 0.015
)

var (
	green   = color.RGBA{10, 255, 50, 255}
	boids   [boidCount]*Boid
	boidMap [screenWidth + 1][screenHeight + 1]int
	rWLock  = sync.RWMutex{}
)

// update is called every frame (1/60 [s]).
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y), green)
		screen.Set(int(boid.position.x), int(boid.position.y-1), green)
		screen.Set(int(boid.position.x), int(boid.position.y+1), green)
	}
	return nil
}

func main() {
	for i := 0; i <= screenWidth; i++ {
		for j := 0; j <= screenHeight; j++ {
			boidMap[i][j] = -1
		}
	}
	rWLock.Lock()
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}
	rWLock.Unlock()
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Boids in a box"); err != nil {
		log.Fatal(err)
	}
}
