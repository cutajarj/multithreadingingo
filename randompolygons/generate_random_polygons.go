package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
)

type Point2D struct {
	x float64
	y float64
}

/**
 * Translated into Go lang from https://cglab.ca/~sander/misc/ConvexGeneration/convex.html
 * Thank you Sander Verdonschot!
 */
func generateRandomConvexPolygon(n int) []Point2D {
	xPool := make([]float64, n)
	yPool := make([]float64, n)
	for i := 0; i < n; i++ {
		xPool[i] = rand.Float64()
		yPool[i] = rand.Float64()
	}

	sort.Float64s(xPool)
	sort.Float64s(yPool)

	minX, maxX, minY, maxY := xPool[0], xPool[n-1], yPool[0], yPool[n-1]

	xVec := make([]float64, n)
	yVec := make([]float64, n)

	lastTop, lastBot := minX, minX

	for i := 1; i < n-1; i++ {
		x := xPool[i]
		if rand.Float32() < 0.5 {
			xVec[i-1] = x - lastTop
			lastTop = x
		} else {
			xVec[i-1] = lastBot - x
			lastBot = x
		}
	}

	xVec[n-2] = maxX - lastTop
	xVec[n-1] = lastBot - maxX

	lastLeft, lastRight := minY, minY

	for i := 1; i < n-1; i++ {
		y := yPool[i]
		if rand.Float32() < 0.5 {
			yVec[i] = y - lastLeft
			lastLeft = y
		} else {
			yVec[i] = lastRight - y
			lastRight = y
		}
	}

	yVec[n-2] = maxY - lastLeft
	yVec[n-1] = lastRight - maxY

	rand.Shuffle(len(yVec), func(i, j int) { yVec[i], yVec[j] = yVec[j], yVec[i] })

	vec := make([]Point2D, n)

	for i := 0; i < n; i++ {
		vec[i] = Point2D{xVec[i], yVec[i]}
	}

	sort.SliceStable(vec, func(i, j int) bool {
		return math.Atan2(vec[i].y, vec[i].x) > math.Atan2(vec[j].y, vec[j].x)
	})

	x, y, minPolygonX, minPolygonY := 0.0, 0.0, 0.0, 0.0
	points := make([]Point2D, n)

	for i := 0; i < n; i++ {
		points[i] = Point2D{x, y}
		x += vec[i].x
		y += vec[i].y

		minPolygonX = math.Min(minPolygonX, x)
		minPolygonY = math.Min(minPolygonY, y)
	}

	xShift := minX - minPolygonX
	yShift := minY - minPolygonY

	for i := 0; i < n; i++ {
		p := points[i]
		points[i] = Point2D{(p.x + xShift) * 300.0, (p.y + yShift) * 300}
	}
	return points
}

func main() {
	file, err := os.Create("polygons.txt")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	for i := 0; i < 100000; i++ {
		points := generateRandomConvexPolygon(rand.Intn(100) + 4)
		for n, p := range points {
			if n == 0 {
				fmt.Fprintf(file, "(%.0f,%.0f)", math.Round(p.x), math.Round(p.y))
			} else {
				fmt.Fprintf(file, ",(%.0f,%.0f)", math.Round(p.x), math.Round(p.y))
			}
		}
		fmt.Fprintf(file, "\n")
	}
	file.Close()

}
