package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type Point struct {
	r int // row
	c int // col
}

var grid [][]rune

func main() {
	test := flag.Bool("test", false, "set true to run test input")
	p1 := flag.Bool("p1", false, "set true to run part 1")
	p2 := flag.Bool("p2", false, "set true to run part 2")
	flag.Parse()

	fname := "input.txt"
	if *test {
		fname = "test.txt"
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
	grid = make([][]rune, len(lines))
	usedRows := map[int]interface{}{}
	usedCols := map[int]interface{}{}
	galaxies := map[int]Point{}
	gCount := 0
	for i, line := range lines {
		runes := []rune(line)
		grid[i] = runes
		for j, r := range runes {
			if r == '#' {
				galaxies[gCount] = Point{r: i, c: j}
				gCount += 1
				usedRows[i] = nil
				usedCols[j] = nil
			}
		}
	}
	l := len(grid)
	offset := 0
	for i := 0; i < l; i++ { // add cols
		if _, ok := usedCols[i]; ok {
			continue
		}
		for k, g := range galaxies {
			if i+offset < g.c {
				g.c += 1
				galaxies[k] = g
			}
		}
		offset += 1
	}
	offset = 0
	for i := 0; i < l; i++ { // add rows
		if _, ok := usedRows[i]; ok {
			continue
		}
		for k, g := range galaxies {
			if i+offset < g.r {
				g.r += 1
				galaxies[k] = g
			}
		}
		offset += 1
	}
	total := 0
	for i := 0; i < len(galaxies); i++ {
		g1 := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			g2 := galaxies[j]
			total += shortestDistance(g1, g2)
			// if i == 0 {
			// 	fmt.Println(shortestDistance(g1, g2))
			// }
		}
	}
	fmt.Println("PART 1 TOTAL :: ", total)
}

func part2(fname string) {
	f, _ := os.ReadFile(fname)
	contents := string(f)
	lines := strings.Split(contents, "\n")
	grid = make([][]rune, len(lines))
	usedRows := map[int]interface{}{}
	usedCols := map[int]interface{}{}
	galaxies := map[int]Point{}
	gCount := 0
	for i, line := range lines {
		runes := []rune(line)
		grid[i] = runes
		for j, r := range runes {
			if r == '#' {
				galaxies[gCount] = Point{r: i, c: j}
				gCount += 1
				usedRows[i] = nil
				usedCols[j] = nil
			}
		}
	}

	l := len(grid[0])
	offsetAmount := 999999 //  *2 means add 1, *10 means add 9, *100 means add 99
	offset := 0
	for i := 0; i < l; i++ { // add cols
		if _, ok := usedCols[i]; ok {
			continue
		}
		for k, g := range galaxies {
			if i+offset < g.c {
				g.c += offsetAmount
				galaxies[k] = g
			}
		}
		offset += offsetAmount
	}

	l = len(grid)
	offset = 0
	for i := 0; i < l; i++ { // add rows
		if _, ok := usedRows[i]; ok {
			continue
		}
		for k, g := range galaxies {
			if i+offset < g.r {
				g.r += offsetAmount
				galaxies[k] = g
			}
		}
		offset += offsetAmount
	}
	total := 0
	for i := 0; i < len(galaxies); i++ {
		g1 := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			g2 := galaxies[j]
			total += shortestDistance(g1, g2)
		}
	}
	fmt.Println("PART 2 TOTAL :: ", total)
}

func shortestDistance(g1, g2 Point) int {
	r := math.Abs(float64(g2.r - g1.r))
	c := math.Abs(float64(g2.c - g1.c))
	return int(r + c)
}

func printGrid() {
	fmt.Println("n:", len(grid))
	fmt.Println("m:", len(grid[0]))
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%c ", col)
		}
		fmt.Printf("\n")
	}
}
