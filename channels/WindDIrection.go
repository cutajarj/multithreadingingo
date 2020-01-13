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

var (
	windRegex = regexp.MustCompile(`\d* METAR.*EGLL \d*Z [A-Z ]*(\d\d\d|VRB).*=`)
	windDist  [8]int
)

func convertToArray(text string) []string {
	lines := strings.Split(text, "\n")
	metarSlice := make([]string, 0, len(lines))
	metarStr := ""
	for _, line := range lines {
		if strings.Contains(line, "TAF") {
			break
		}
		if !strings.HasPrefix(line, "#") {
			metarStr += strings.Trim(line, " ")
		}
		if strings.HasSuffix(line, "=") {
			metarSlice = append(metarSlice, metarStr)
			metarStr = ""
		}
	}
	return metarSlice
}

func extractWindDirection(metars []string) []string {
	winds := make([]string, 0, len(metars))
	for _, metar := range metars {
		if windRegex.MatchString(metar) {
			winds = append(winds, windRegex.FindAllStringSubmatch(metar, -1)[0][1])
		}
	}
	return winds
}

func extractWindDistribution(winds []string) {
	for _, wind := range winds {
		if d, err := strconv.ParseFloat(wind, 64); err == nil {
			dirIndex := int(math.Round(d/45.0)) % 8
			windDist[dirIndex]++
		}
	}
}

func main() {
	absPath, _ := filepath.Abs("./multithreadingingo/metarfiles/")
	files, _ := ioutil.ReadDir(absPath)
	start := time.Now()
	for _, file := range files {
		dat, err := ioutil.ReadFile(filepath.Join(absPath, file.Name()))
		//dat, err := ioutil.ReadFile(filepath.Join(absPath, "apr2013.txt"))
		if err != nil {
			panic(err)
		}
		text := string(dat)
		//1. Change to array, each metar report is a separate item in the array
		metars := convertToArray(text)
		//2. Regex extract wind direction, EGLL 312350Z 07004KT CAVOK 12/09 Q1016 NOSIG= -> 070
		winds := extractWindDirection(metars)
		//3. Assign to N, NE, E, SE, S, SW, W, NW, 070 -> E + 1
		extractWindDistribution(winds)
	}
	elapsed := time.Since(start)
	fmt.Printf("%v\n", windDist)
	fmt.Printf("Processing took %s", elapsed)
}
