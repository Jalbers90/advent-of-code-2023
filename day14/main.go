package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

var DEBUG = 0

type Point struct {
	row int
	col int
}

var Dirs = []Point{ // N, W, S, E
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
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
	rounds := []Point{}
	for i, line := range lines {
		runes := []rune(line)
		grid = append(grid, runes)
		for j, r := range runes {
			if r == 'O' {
				rounds = append(rounds, Point{row: i, col: j})
			}
		}
	}

	total := 0
	for i, rock := range rounds {
		start := Point{row: rock.row, col: rock.col}
		newPoint := tilt(rock, Point{-1, 0})
		rounds[i] = newPoint
		if start.row != newPoint.row {
			grid[start.row][start.col] = '.'
			grid[newPoint.row][newPoint.col] = 'O'
		}
		total += len(grid) - newPoint.row
	}
	fmt.Println("PART 1 TOTAL :: ", total)
}

func part2(lines []string) {
	rounds := []Point{}
	for i, line := range lines {
		runes := []rune(line)
		grid = append(grid, runes)
		for j, r := range runes {
			if r == 'O' {
				rounds = append(rounds, Point{row: i, col: j})
			}
		}
	}

	nRounds := len(rounds)
	totalMap := map[int][]int{}
	// (1000000000 - number of junk states) % cycle_length = state number answer lands on in cycle
	for n := 0; n < 100; n++ { // 1000000000
		for i, dir := range Dirs {
			if i < 2 { // loop forward North and West
				for j, rock := range rounds {
					start := Point{row: rock.row, col: rock.col}
					newPoint := tilt(rock, dir)
					rounds[j] = newPoint
					// dlog("Start:%+v, End:%+v, rounds[j]: %+v", rock, newPoint, rounds[j])
					if (start.row != newPoint.row) || (start.col != newPoint.col) {
						grid[start.row][start.col] = '.'
						grid[newPoint.row][newPoint.col] = 'O'
					}
				}

			} else { // loop backward South and East
				for j := nRounds - 1; j > -1; j-- {
					rock := rounds[j]
					start := Point{row: rock.row, col: rock.col}
					newPoint := tilt(rock, dir)
					rounds[j] = newPoint
					if (start.row != newPoint.row) || (start.col != newPoint.col) {
						grid[start.row][start.col] = '.'
						grid[newPoint.row][newPoint.col] = 'O'
					}
				}
			}
			sortRounds(rounds)
		}
		// printGrid()
		total := 0
		for _, rock := range rounds {
			total += len(grid) - rock.row
		}
		if _, ok := totalMap[total]; ok {
			totalMap[total] = append(totalMap[total], n+1)
			dlog("EQUAL CYCLE FOUND AT CYCLE %d", n+1)
		} else {
			totalMap[total] = append(totalMap[total], n+1)
		}
		dlog("Cycles: %d ::: TOTAL: %d", n+1, total)
	}
	dlog("TOTAL MAP ::: %+v\n", totalMap)
}

func tilt(rock, dir Point) Point {
	newPoint := Point{row: rock.row + dir.row, col: rock.col + dir.col}
	if outOfBounds(newPoint) {
		return rock
	}
	if grid[newPoint.row][newPoint.col] == '#' || grid[newPoint.row][newPoint.col] == 'O' {
		return rock
	}

	newPoint = tilt(newPoint, dir)
	return newPoint

}

func outOfBounds(p Point) bool {
	return p.row < 0 || p.row >= len(grid) || p.col < 0 || p.col >= len(grid[0])
}

func dlog(format string, args ...interface{}) {
	if DEBUG == 1 {
		log.Printf(format, args...)
	}
}

func sortRounds(rounds []Point) {
	sort.SliceStable(rounds, func(i, j int) bool {
		a := rounds[i]
		b := rounds[j]
		if a.row != b.row {
			return a.row < b.row
		}
		return a.col < b.col
	})
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
