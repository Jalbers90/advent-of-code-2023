package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

var DEBUG = 0

type Point struct {
	row int
	col int
	dir rune // 'N', 'W', 'S', 'E'
}

var dirs = map[rune][2]int{
	'N': {-1, 0},
	'W': {0, -1},
	'S': {1, 0},
	'E': {0, 1},
}

var angles = map[rune]map[rune]Point{
	'/':  {'N': {0, 1, 'E'}, 'W': {1, 0, 'S'}, 'S': {0, -1, 'W'}, 'E': {-1, 0, 'N'}},
	'\\': {'N': {0, -1, 'W'}, 'W': {-1, 0, 'N'}, 'S': {0, 1, 'E'}, 'E': {1, 0, 'S'}},
}

var flats = map[rune]map[rune][]Point{
	'-': {'N': {{0, 1, 'E'}, {0, -1, 'W'}}, 'W': {{0, -1, 'W'}}, 'S': {{0, 1, 'E'}, {0, -1, 'W'}}, 'E': {{0, 1, 'E'}}}, // double for going N or S
	'|': {'N': {{-1, 0, 'N'}}, 'W': {{1, 0, 'S'}, {-1, 0, 'N'}}, 'S': {{1, 0, 'S'}}, 'E': {{1, 0, 'S'}, {-1, 0, 'N'}}}, // double for going W or E

}

var grid [][]rune

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
	for _, line := range lines {
		runes := []rune(line)
		grid = append(grid, runes)
	}

	if *p1 {
		part1(Point{0, 0, 'E'})
	}
	if *p2 {
		part2()
	}
	end := time.Since(start)
	fmt.Println("RUNTIME ::: ", end)
}

func part1(start Point) int {
	q := []Point{start}
	visited := map[Point]interface{}{}
	startKey := fmt.Sprintf("%d,%d", start.row, start.col)
	energized := map[string]interface{}{
		startKey: nil,
	}
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			point := q[0]
			q = q[1:]
			visited[point] = nil
			pointType := grid[point.row][point.col]
			newPoints := []Point{}
			if pointType == '.' {
				dr, dc := dirs[point.dir][0], dirs[point.dir][1]
				np := Point{row: point.row + dr, col: point.col + dc, dir: point.dir}
				if !outOfBounds(np.row, np.col) {
					newPoints = append(newPoints, np)
				}
			} else if pointType == '/' || pointType == '\\' {
				delta := angles[pointType][point.dir]
				dr, dc, d := delta.row, delta.col, delta.dir
				np := Point{row: point.row + dr, col: point.col + dc, dir: d}
				if !outOfBounds(np.row, np.col) {
					newPoints = append(newPoints, np)
				}
			} else {
				deltas := flats[pointType][point.dir]
				for _, delta := range deltas {
					dr, dc, d := delta.row, delta.col, delta.dir
					np := Point{row: point.row + dr, col: point.col + dc, dir: d}
					if !outOfBounds(np.row, np.col) {
						newPoints = append(newPoints, np)
					}
				}
			}

			for _, p := range newPoints {
				key := fmt.Sprintf("%d,%d", p.row, p.col)
				if _, ok := energized[key]; !ok {
					energized[key] = nil
				}

				if _, ok := visited[p]; !ok {
					q = append(q, p)
				}
			}
			dlog("%+v\n", q)
		}
	}
	// fmt.Println("PART 1 TOTAL :: ", len(energized))
	return len(energized)
}

func part2() {
	maxEnergy := float64(0)

	for i := 0; i < len(grid[0]); i++ {
		maxEnergy = math.Max(maxEnergy, float64(part1(Point{0, i, 'S'})))
		maxEnergy = math.Max(maxEnergy, float64(part1(Point{len(grid) - 1, i, 'N'})))
	}

	for i := 0; i < len(grid); i++ {
		maxEnergy = math.Max(maxEnergy, float64(part1(Point{i, 0, 'E'})))
		maxEnergy = math.Max(maxEnergy, float64(part1(Point{i, len(grid[0]) - 1, 'W'})))
	}

	fmt.Println("PART 2 TOTAL :: ", maxEnergy)
}

func outOfBounds(r, c int) bool {
	return r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0])
}

func printGrid() {
	fmt.Printf("\n")
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%c ", col)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func dlog(format string, args ...interface{}) {
	if DEBUG == 1 {
		log.Printf(format, args...)
	}
}
