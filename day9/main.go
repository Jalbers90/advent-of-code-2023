package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

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
	total := 0
	for _, line := range lines {
		nums := strings.Split(line, " ")
		row1 := []int{}
		for _, s := range nums {
			n, _ := strconv.Atoi(s)
			row1 = append(row1, n)
		}
		last := row1[len(row1)-1]
		for i := 0; i < len(nums); i++ {
			l := len(row1)
			row2 := []int{}
			for j := 1; j < l; j++ {
				row2 = append(row2, row1[j]-row1[j-1])
			}
			last += row2[len(row2)-1]
			row1 = make([]int, len(row2))
			copy(row1, row2)
			if allEqual(row2) {
				break
			}

		}
		total += last
	}

	fmt.Println("PART 1 TOTAL ::: ", total)
}

func part2(fname string) {
	f, _ := os.ReadFile(fname)
	contents := string(f)
	lines := strings.Split(contents, "\n")
	total := 0
	for _, line := range lines {
		nums := strings.Split(line, " ")
		row1 := []int{}
		for _, s := range nums {
			n, _ := strconv.Atoi(s)
			row1 = append(row1, n)
		}
		reverseSlice(row1)
		last := row1[len(row1)-1]
		for i := 0; i < len(nums); i++ {
			l := len(row1)
			row2 := []int{}
			for j := 1; j < l; j++ {
				row2 = append(row2, row1[j]-row1[j-1])
			}
			last += row2[len(row2)-1]
			row1 = make([]int, len(row2))
			copy(row1, row2)
			if allEqual(row2) {
				break
			}

		}
		// fmt.Println(last)
		total += last
	}

	fmt.Println("PART 1 TOTAL ::: ", total)
}

func allEqual(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[i-1] {
			return false
		}
	}
	return true
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
