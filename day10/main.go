package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

var grid [][]rune
var n int                                              // rows
var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left, right (row, col)

type Point struct {
	coords [2]int
	route  rune
}

var pipeMap = map[rune][][2]int{ // each pipe can go two directions
	'|': {{1, 0}, {-1, 0}}, // row, col
	'-': {{0, 1}, {0, -1}},
	'L': {{-1, 0}, {0, 1}},
	'J': {{-1, 0}, {0, -1}},
	'7': {{1, 0}, {0, -1}},
	'F': {{1, 0}, {0, 1}},
	// 'S': {{1, 0}, {0, 1}}, // test p1
	// 'S': {{1, 0}, {0, -1}}, // test2 p2 ... 7
	// 'S': {{1, 0}, {0, 1}}, // test3 or test4 p2 ... F
	'S': {{-1, 0}, {0, 1}}, // input
	'.': {},
}

func main() {
	test := flag.Bool("test", false, "set true to run test input")
	p1 := flag.Bool("p1", false, "set true to run part 1")
	p2 := flag.Bool("p2", false, "set true to run part 2")
	flag.Parse()

	fname := "input.txt"
	if *test {
		fname = "test3.txt"
	}

	start := time.Now()
	if *p1 {
		part1(fname)
	}
	if *p2 {
		part2(fname)
	}
	end := time.Since(start)
	fmt.Println("RUNTIME ::: ", end)
}

func part1(fname string) {
	f, _ := os.ReadFile(fname)
	contents := string(f)
	lines := strings.Split(contents, "\n")
	n = len(lines)
	grid = make([][]rune, n)
	q := []Point{}
	for i, line := range lines {
		runes := []rune(line)
		grid[i] = runes
		for j, r := range runes {
			if r == 'S' {
				letter := pipeMap['S']
				for _, coord := range letter {
					r, c := coord[0]+i, coord[1]+j
					q = append(q, Point{coords: [2]int{r, c}})
				}
				q[0].route = 'A'
				q[1].route = 'B'
				grid[i][j] = 'X'
			}
		}
	}

	stepMap := map[rune]int{
		'A': 1,
		'B': 1,
	}
	// S = 'L'
	for len(q) > 0 {
		lvlLen := len(q)
		for i := 0; i < lvlLen; i++ {
			point := q[0] // pop from front
			q = q[1:]     // pop from front
			row, col := point.coords[0], point.coords[1]
			letter := grid[row][col]
			next := pipeMap[letter]
			// fmt.Printf("CUR POINT ::: %+v LETTER: %c\n", point, letter)
			grid[row][col] = 'X' // so we don't go backwards
			for _, p := range next {
				r, c := row+p[0], col+p[1]
				if grid[r][c] == 'X' {
					continue
				}
				newPoint := Point{coords: [2]int{r, c}, route: point.route}
				q = append(q, newPoint)
				stepMap[point.route] += 1
			}
		}
	}

	steps := int(math.Max(float64(stepMap['A']), float64(stepMap['B'])))
	fmt.Println("PART 1 STEPS ::: ", steps)
}

func part2(fname string) {
	f, _ := os.ReadFile(fname)
	contents := string(f)
	lines := strings.Split(contents, "\n")

	grid = make([][]rune, len(lines))
	q := []Point{}     // for our bfs's
	path := [][2]int{} // slice containing all coords of the loop

	for i, line := range lines {
		runes := []rune(line)
		grid[i] = runes
		for j, r := range runes {
			if r == 'S' {
				letter := pipeMap['S']
				r, c := letter[0][0]+i, letter[0][1]+j
				q = append(q, Point{coords: [2]int{r, c}})
				path = append(path, [2]int{i, j})
				grid[i][j] = 'X'
			}
		}
	}

	for len(q) > 0 {
		point := q[0] // pop from front
		q = q[1:]     // pop from front
		row, col := point.coords[0], point.coords[1]
		letter := grid[row][col]
		next := pipeMap[letter]
		// fmt.Printf("CUR POINT ::: %+v LETTER: %c\n", point, letter)
		grid[row][col] = 'X' // so we don't go backwards
		path = append(path, [2]int{row, col})
		for _, p := range next {
			r, c := row+p[0], col+p[1]
			if grid[r][c] == 'X' {
				continue
			}
			newCoords := [2]int{r, c}
			newPoint := Point{coords: newCoords, route: point.route}
			q = append(q, newPoint)
		}
	}

	// range path and do a dfs for any connected points that aren't part of the path or already visited
	// simple flood fill
	// visited '.' now turn into 'O' just in case
	// ray cast algo to check if point (and rest of area) is inside polygon or not
	enclosed := 0
	// printGrid()
	for _, X := range path {
		row, col := X[0], X[1]
		for _, dir := range dirs {
			r, c := row+dir[0], col+dir[1]

			if outOfBounds(r, c) {
				continue
			} else if grid[r][c] != 'X' && grid[r][c] != 'O' {
				tiles, escaped := countTiles(r, c)
				// fmt.Printf("COUNT-TILES :: (%d, %d) :: TILES %d :: ESCAPED :: %t\n", r, c, tiles, escaped)
				if !escaped && isInsidePolygon(r, c, path) {
					enclosed += tiles
				}
			}
		}
	}

	fmt.Println("PART 2 ENCLOSED TILES ::: ", enclosed)
}

func outOfBounds(r, c int) bool {
	return r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0])
}

func countTiles(row, col int) (int, bool) {
	if outOfBounds(row, col) {
		return 0, true
	} else if grid[row][col] == 'X' || grid[row][col] == 'O' {
		return 0, false
	}

	grid[row][col] = 'O'
	tiles := 1
	escaped := false
	// fmt.Println("START COUNT-TILES :: ", tiles, row, col)
	for _, d := range dirs {
		r, c := d[0]+row, d[1]+col
		newTiles, e := countTiles(r, c)
		tiles += newTiles
		if e {
			escaped = true
		}
	}
	return tiles, escaped
}

func printGrid() {
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%c ", col)
		}
		fmt.Printf("\n")
	}
}

func isInsidePolygon(r, c int, polygon [][2]int) bool {
	intersections := 0
	l := len(polygon)
	testPoint := [2]int{r, c} // row, col
	for i := 0; i < l; i++ {
		next := (i + 1) % l
		if intersectsEdge(testPoint, polygon[i], polygon[next]) {
			intersections += 1
		}
	}
	return intersections%2 == 1
}

func intersectsEdge(test, start, end [2]int) bool {
	if (start[0] > test[0]) != (end[0] > test[0]) {
		if test[1] < (end[1]-start[1])*(test[0]-start[0])/(end[0]-start[0])+start[1] {
			return true
		}
	}

	return false
}

// return (edgeStart.Y > testPoint.Y) != (edgeEnd.Y > testPoint.Y) &&
// 		testPoint.X < (edgeEnd.X-edgeStart.X)*(testPoint.Y-edgeStart.Y)/(edgeEnd.Y-edgeStart.Y)+edgeStart.X
