package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var DEBUG = 0

type Point struct {
	r int
	c int
}

var dirs = map[string]Point{
	"U": {-1, 0},
	"L": {0, -1},
	"D": {1, 0},
	"R": {0, 1},
}

var numToDir = map[string]string{
	"0": "R",
	"1": "D",
	"2": "L",
	"3": "U",
}

func main() {
	test := flag.Bool("test", false, "set true to run test input")
	d := flag.Bool("d", false, "set true to print dlog's")
	p1 := flag.Bool("p1", false, "set true to run part 1")
	p2 := flag.Bool("p2", false, "set true to run part 2")
	flag.Parse()

	if *d {
		DEBUG = 1
	}

	fname := "input.txt"
	if *test {
		fname = "test.txt"
	}

	start := time.Now()
	f, _ := os.ReadFile(fname)
	contents := string(f)
	lines := strings.Split(contents, "\n")
	if *p1 {
		part1(lines)
	}
	if *p2 {
		part2(lines)
	}
	end := time.Since(start)
	fmt.Println("RUNTIME ::: ", end)
}

func part1(lines []string) {
	path := []Point{{0, 0}}
	cur := Point{0, 0}
	boundary := 1
	for _, line := range lines {
		splitline := strings.Split(line, " ")
		amount, _ := strconv.Atoi(splitline[1])
		d := splitline[0]
		nextPoint := Point{r: cur.r + dirs[d].r*int(amount), c: cur.c + dirs[d].c*int(amount)}
		path = append(path, nextPoint)
		boundary += int(amount)
		cur.r, cur.c = nextPoint.r, nextPoint.c
	}
	path = path[0 : len(path)-1]
	boundary -= 1
	area := polygonArea(path)
	dlog("POINTS :: %+v\n", path)
	dlog("AREA :: %d\n", area)
	dlog("BOUNDARY :: %d\n", boundary)
	fmt.Println("PART 1 AREA :: ", area+(boundary/2)+1)
}

func part2(lines []string) {
	path := []Point{{0, 0}}
	cur := Point{0, 0}
	boundary := 1
	for _, line := range lines {
		splitline := strings.Split(line, " ")
		color := splitline[2]
		amount, _ := strconv.ParseInt(color[2:7], 16, 64)
		d := numToDir[string(color[7])]
		nextPoint := Point{r: cur.r + dirs[d].r*int(amount), c: cur.c + dirs[d].c*int(amount)}
		path = append(path, nextPoint)
		boundary += int(amount)
		cur.r, cur.c = nextPoint.r, nextPoint.c
		// dlog("%s\n", color[2:7])
	}
	path = path[0 : len(path)-1]
	boundary -= 1
	area := polygonArea(path)
	fmt.Println("PART 2 AREA :: ", area+(boundary/2)+1)
}

func polygonArea(path []Point) int {
	// shoelace theorem
	n := len(path)
	if n < 3 { // A polygon must have at least three vertices
		return 0
	}
	area := 0.0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += float64(path[i].c*path[j].r - path[j].c*path[i].r)
	}
	area = math.Abs(area) / 2.0
	return int(area)
}

func dlog(format string, args ...interface{}) {
	if DEBUG == 1 {
		log.Printf(format, args...)
	}
}
