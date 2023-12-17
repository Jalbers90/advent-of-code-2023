package main

import (
	"container/heap"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var DEBUG = 0

var grid [][]int

var dirs = []string{"N", "W", "S", "E"}
var reverseDirs = map[string]string{
	"N": "S",
	"W": "E",
	"S": "N",
	"E": "W",
}
var dirMap = map[string][2]int{
	"N": {-1, 0},
	"W": {0, -1},
	"S": {1, 0},
	"E": {0, 1},
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
		fname = "test2.txt"
	}

	start := time.Now()
	f, _ := os.ReadFile(fname)
	contents := string(f)
	lines := strings.Split(contents, "\n")
	for i, line := range lines {
		runes := []rune(line)
		l := len(runes)
		grid = append(grid, make([]int, l))
		for j, r := range runes {
			grid[i][j], _ = strconv.Atoi(string(r))
		}
	}
	if *p1 {
		part1()
	}
	if *p2 {
		part2()
	}
	end := time.Since(start)
	fmt.Println("RUNTIME ::: ", end)
}

func part1() {
	q := &PointHeap{{0, 0, 0, 0, ""}} // row, col, heatloss, steps, dir
	heap.Init(q)
	seen := map[string]interface{}{}
	end := Point{r: len(grid) - 1, c: len(grid[0]) - 1}
	minHeatloss := 0
	for q.Len() > 0 {
		// dlog("%+v\n", q)
		cur := heap.Pop(q).(Point)
		seenstr := fmt.Sprintf("%d,%d,%s,%d", cur.r, cur.c, cur.dir, cur.steps)
		if _, ok := seen[seenstr]; ok {
			// fmt.Println("HIII ::: ", cur)
			continue
		}
		seen[seenstr] = nil
		if cur.r == end.r && cur.c == end.c {
			minHeatloss = cur.heatloss
			break
		}
		for _, d := range dirs {
			if (d == cur.dir && cur.steps >= 3) || d == reverseDirs[cur.dir] {
				continue
			}
			nextRow, nextCol := cur.r+dirMap[d][0], cur.c+dirMap[d][1]
			if !outOfBounds(nextRow, nextCol) {
				nextheat := cur.heatloss + grid[nextRow][nextCol]
				newSteps := 1
				if d == cur.dir {
					newSteps += cur.steps
				}
				newPoint := Point{r: nextRow, c: nextCol, heatloss: nextheat, steps: newSteps, dir: d}
				heap.Push(q, newPoint)
			}
		}
	}
	fmt.Println("PART 1 HEAT LOSS :: ", minHeatloss)
}

func part2() {
	q := &PointHeap{{0, 0, 0, 0, "S"}, {0, 0, 0, 0, "E"}} // row, col, heatloss, steps, dir
	heap.Init(q)
	seen := map[string]interface{}{}
	end := Point{r: len(grid) - 1, c: len(grid[0]) - 1}
	minHeatloss := 0
	for q.Len() > 0 {
		cur := heap.Pop(q).(Point)
		seenstr := fmt.Sprintf("%d,%d,%s,%d", cur.r, cur.c, cur.dir, cur.steps)
		if _, ok := seen[seenstr]; ok {
			continue
		}
		seen[seenstr] = nil
		if cur.r == end.r && cur.c == end.c && cur.steps >= 4 {
			minHeatloss = cur.heatloss
			break
		}
		for _, d := range dirs {
			if (d == cur.dir && cur.steps == 10) || (cur.steps < 4 && d != cur.dir) || d == reverseDirs[cur.dir] {
				continue
			}
			nextRow, nextCol := cur.r+dirMap[d][0], cur.c+dirMap[d][1]
			if !outOfBounds(nextRow, nextCol) {
				nextheat := cur.heatloss + grid[nextRow][nextCol]
				newSteps := 1
				if d == cur.dir {
					newSteps += cur.steps
				}
				newPoint := Point{r: nextRow, c: nextCol, heatloss: nextheat, steps: newSteps, dir: d}
				heap.Push(q, newPoint)
			}
		}
	}
	fmt.Println("PART 2 HEAT LOSS :: ", minHeatloss)
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
		fmt.Printf(format, args...)
	}
}
