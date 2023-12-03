package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var n, m int // total rows, cols
var grid [][]rune
var adjacent = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, 1}, {1, 1}, {-1, -1}, {1, -1}}

func main() {
	test := flag.Bool("test", false, "set true to run test input")
	p1 := flag.Bool("p1", false, "set true to run part 1")
	p2 := flag.Bool("p2", false, "set true to run part 2")
	flag.Parse()

	fname := "input.txt"
	if *test {
		// fname = "edge_test.txt"
		fname = "test.txt"
	}
	if *p1 {
		part1(fname)
	} else if *p2 {
		part2(fname)
	}
}

func part1(fname string) {
	f, _ := os.ReadFile(fname)
	content := string(f)
	lines := strings.Split(content, "\n")
	nums := []map[int][][]int{} // slice of {KEY: []int}
	n = len(lines)
	m = len([]rune(lines[0]))
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	// check if each digit is adjacent to a symbol
	for i, runes := range grid {
		cur := "" // current working number
		coords := [][]int{}
		for j, c := range runes {
			if unicode.IsDigit(c) {
				cur += string(c)
				coords = append(coords, []int{i, j})
			} else {
				if cur != "" {
					curInt, _ := strconv.Atoi(cur)
					m := map[int][][]int{}
					m[curInt] = coords
					nums = append(nums, m)
					cur = ""
					coords = [][]int{}
				}
			}
		}
		if cur != "" {
			curInt, _ := strconv.Atoi(cur)
			m := map[int][][]int{}
			m[curInt] = coords
			nums = append(nums, m)
		}
	}

	total := 0
	for _, m := range nums {
		// found[key] = map[string]bool{}
		// foundSymbol := false
		contains := []string{}
		for key, coords := range m {
			for _, point := range coords {
				row, col := point[0], point[1]

				for _, adj := range adjacent {
					r, c := row+adj[0], col+adj[1]
					if checkAdjacent(r, c, contains) {
						total += key
						contains = append(contains, fmt.Sprintf("%d,%d", r, c))
					}
				}
			}
			fmt.Println("KEY ", key, " ::: CONTAINS ", contains)
		}
	}

	fmt.Printf("PART 1 TOTAL :::: %d\n", total)
	// printFound(nums)
}

func part2(fname string) {
	f, _ := os.ReadFile(fname)
	content := string(f)
	lines := strings.Split(content, "\n")
	n = len(lines)
	m = len([]rune(lines[0]))
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	stars := [][]int{}
	allCoords := map[string]int{}
	for i, runes := range grid {
		cur := "" // current working number
		coords := [][]int{}
		for j, c := range runes {
			if c == '*' {
				stars = append(stars, []int{i, j})
			}
			if unicode.IsDigit(c) {
				cur += string(c)
				coords = append(coords, []int{i, j})
			} else {
				if cur != "" {
					curInt, _ := strconv.Atoi(cur)
					for _, p := range coords {
						point := fmt.Sprintf("%d,%d", p[0], p[1])
						allCoords[point] = curInt
					}
					cur = ""
					coords = [][]int{}
				}
			}
		}
		if cur != "" {
			curInt, _ := strconv.Atoi(cur)
			for _, p := range coords {
				point := fmt.Sprintf("%d,%d", p[0], p[1])
				allCoords[point] = curInt
			}
		}
	}

	// fmt.Println(allCoords)
	total := 0
	for _, gear := range stars {
		counted := map[int]bool{}
		row, col := gear[0], gear[1]
		ratio := 1
		found := 0
		for _, adj := range adjacent {
			r, c := row+adj[0], col+adj[1]
			if checkGearAdjacent(r, c) {
				p := fmt.Sprintf("%d,%d", r, c)
				num := allCoords[p]
				_, ok := counted[num]
				if !ok {
					ratio *= num
					found += 1
					counted[num] = true
				}
			}
		}
		if found == 2 {
			total += ratio
		}
	}

	fmt.Printf("PART 2 TOTAL :::: %d\n", total)

}

func checkAdjacent(r int, c int, contains []string) bool {
	// check out of bounds
	if c < 0 || c >= m || r < 0 || r >= n {
		return false
	}

	// check new point
	if grid[r][c] != '.' && !unicode.IsDigit(grid[r][c]) && !containsElement(contains, fmt.Sprintf("%d,%d", r, c)) {
		return true
	}

	return false
}

func checkGearAdjacent(r, c int) bool {
	// check out of bounds
	if c < 0 || c >= m || r < 0 || r >= n {
		return false
	}
	// check new point
	if unicode.IsDigit(grid[r][c]) {
		return true
	}

	return false
}

func containsElement(slice []string, target string) bool {
	for _, element := range slice {
		if element == target {
			return true
		}
	}
	return false
}
