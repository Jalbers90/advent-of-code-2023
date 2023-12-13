package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var patterns [][][]rune

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
	grid := [][]rune{}
	for i, line := range lines {
		if line == "" {
			patterns = append(patterns, grid)
			grid = [][]rune{}
			continue
		}
		runes := []rune(line)
		grid = append(grid, runes)
		if i == len(lines)-1 {
			patterns = append(patterns, grid)
			grid = [][]rune{}
		}
	}

	total := 0
	for _, pattern := range patterns {
		reflect := colReflection(pattern, 0)
		if reflect == -1 {
			reflect = rowReflection(pattern, 0)
		}
		total += reflect
	}

	fmt.Println("PART 1 TOTAL :: ", total)
}

func part2(lines []string) {
	grid := [][]rune{}
	for i, line := range lines {
		if line == "" {
			patterns = append(patterns, grid)
			grid = [][]rune{}
			continue
		}
		runes := []rune(line)
		grid = append(grid, runes)
		if i == len(lines)-1 {
			patterns = append(patterns, grid)
			grid = [][]rune{}
		}
	}

	total := 0
	for _, pattern := range patterns {
		reflect := colReflection(pattern, 1)
		if reflect == -1 {
			reflect = rowReflection(pattern, 1)
		}
		// fmt.Println("MAIN LOOP ::", "i:", i, "REFLECTION POINTS ::", reflect)
		total += reflect
	}

	fmt.Println("PART 1 TOTAL :: ", total)
}

func rowReflection(p [][]rune, allowedErrors int) int {
	found := -1
	for i := 1; i < len(p); i++ {
		errors := compareRows(i, i-1, p)
		if errors <= allowedErrors {
			errors = 0
			edge, inside := closestEdge(i-1, i, len(p))
			left := 0
			right := 0
			if edge > inside {
				right = edge
				left = inside
			} else {
				right = inside
				left = edge
			}
			for left < right {
				errors += compareRows(left, right, p)
				if errors > allowedErrors {
					break
				}
				left += 1
				right -= 1
			}
			if left > right && errors == allowedErrors {
				found = i * 100
				break
			}
		}
	}
	return found
}

func closestEdge(n1, n2, length int) (int, int) {
	if length-n2 <= n1 {
		return length - 1, n1 - (length - n2) + 1
	} else {
		return 0, n1*2 + 1
	}
}

func colReflection(p [][]rune, allowedErrors int) int {
	found := -1
	for i := 1; i < len(p[0]); i++ {
		errors := compareCols(i, i-1, p)
		if errors <= allowedErrors {
			errors = 0
			edge, inside := closestEdge(i-1, i, len(p[0]))
			left := 0
			right := 0
			if edge > inside {
				right = edge
				left = inside
			} else {
				right = inside
				left = edge
			}
			for left < right {
				errors += compareCols(left, right, p)
				if errors > allowedErrors {
					break
				}
				left += 1
				right -= 1
			}
			if left > right && errors == allowedErrors {
				found = i
				break
			}
		}
	}
	return found
}

func compareRows(r1, r2 int, grid [][]rune) int {
	l := len(grid[r1])
	errors := 0
	for i := 0; i < l; i++ {
		if grid[r1][i] != grid[r2][i] {
			errors += 1
		}
	}
	return errors
}

func compareCols(c1, c2 int, grid [][]rune) int {
	errors := 0
	for i := range grid {
		if grid[i][c1] != grid[i][c2] {
			errors += 1
		}
	}
	return errors
}

func printGrid(grid [][]rune) {
	fmt.Printf("\n")
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%c ", col)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
