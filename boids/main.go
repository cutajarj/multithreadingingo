package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
	"sync"
)

/*
TODO:
1. Put in proper project and remove all particles references
2. Put on github
3. Create readme.md with instructions on how to install and run
4. Mention dependencies

For lesson videos:
1. Create vector class, one boid as a thread, with borderBounce
2. Show inter thread communication, by shared memory using multiple threads/boids for average velocity
3. Use normal mutex to show proximity average
4. Use Read-write mutex to show other averages
*/

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
