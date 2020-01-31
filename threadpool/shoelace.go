package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Point2D struct {
	x int
	y int
}

var (
	r = regexp.MustCompile(`\((\d*),(\d*)\)`)
)

func findArea(pointsStr string) float64 {
	var points []Point2D
	for _, p := range r.FindAllStringSubmatch(pointsStr, -1) {
		x, _ := strconv.Atoi(p[1])
		y, _ := strconv.Atoi(p[2])
		points = append(points, Point2D{x, y})
	}

	area := 0.0
	for i := 0; i < len(points); i++ {
		a, b := points[i], points[(i+1)%len(points)]
		area += float64(a.x*b.y) - float64(a.y*b.x)
	}
	return math.Abs(area) / 2.0
}

func main() {
	absPath, _ := filepath.Abs("./multithreadingingo/threadpool/")
	dat, err := ioutil.ReadFile(filepath.Join(absPath, "polygons.txt"))
	if err != nil {
		panic(err)
	}
	text := string(dat)

	start := time.Now()
	for _, line := range strings.Split(text, "\n") {
		//(3,4),(5,11),(12,8),(9,5),(5,6)
		findArea(line)
		//fmt.Println("Area is ", ans)
	}
	elapsed := time.Since(start)
	fmt.Printf("Processing took %s\n", elapsed)
}
