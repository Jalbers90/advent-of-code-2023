package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
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
	content := string(f)
	lines := strings.Split(content, "\n")
	sourceStrs := strings.Split(strings.TrimPrefix(lines[0], "seeds: "), " ")
	sources := make([]int, len(sourceStrs))
	for i := 0; i < len(sourceStrs); i++ {
		sources[i], _ = strconv.Atoi(sourceStrs[i])
	}
	lines = lines[2:]
	ranges := [][3]int{}
	for _, line := range lines {
		if len(line) == 0 {
			// update sources
			// fmt.Println("update sources")
			for i := 0; i < len(sources); i++ {
				s := sources[i]
				for _, mapping := range ranges {
					newSource := checkRange(s, mapping)
					if newSource != -1 {
						sources[i] = newSource
						break
					}
				}
			}
		} else if !unicode.IsDigit(rune(line[0])) { // title line
			// reset on title lines
			ranges = [][3]int{}
		} else {
			// collect ranges
			rangesStr := strings.Split(line, " ")
			ranges = append(ranges, [3]int{})
			l := len(ranges)
			ranges[l-1][0], _ = strconv.Atoi(rangesStr[0]) // dest
			ranges[l-1][1], _ = strconv.Atoi(rangesStr[1]) // source
			ranges[l-1][2], _ = strconv.Atoi(rangesStr[2]) // range length
		}
	}

	// parse last time
	for i := 0; i < len(sources); i++ {
		s := sources[i]
		for _, mapping := range ranges {
			newSource := checkRange(s, mapping)
			if newSource != -1 {
				sources[i] = newSource
				break
			}
		}
	}
	min := math.MaxInt
	for _, s := range sources {
		min = int(math.Min(float64(s), float64(min)))
	}
	fmt.Println("PART 1 MINIMUM LOCATION ::: ", min)
}

func part2(fname string) {
	f, _ := os.ReadFile(fname)
	content := string(f)
	lines := strings.Split(content, "\n")
	sourceStrs := strings.Split(strings.TrimPrefix(lines[0], "seeds: "), " ")

	sources := []int{}
	start := 0
	rangeLen := 0
	for i := 0; i < len(sourceStrs); i++ {
		if i%2 == 0 {
			start, _ = strconv.Atoi(sourceStrs[i])
		} else {
			rangeLen, _ = strconv.Atoi(sourceStrs[i])
			total := start + rangeLen
			for j := start; j < total; j++ {
				sources = append(sources, j)
			}
		}
	}
	// fmt.Println(sources)
	lines = lines[2:]
	ranges := [][3]int{}
	for _, line := range lines {
		if len(line) == 0 {
			// update sources
			// fmt.Println("update sources")
			for i := 0; i < len(sources); i++ {
				s := sources[i]
				for _, mapping := range ranges {
					newSource := checkRange(s, mapping)
					if newSource != -1 {
						sources[i] = newSource
						break
					}
				}
			}
		} else if !unicode.IsDigit(rune(line[0])) { // title line
			// reset on title lines
			ranges = [][3]int{}
		} else {
			// collect ranges
			rangesStr := strings.Split(line, " ")
			ranges = append(ranges, [3]int{})
			l := len(ranges)
			ranges[l-1][0], _ = strconv.Atoi(rangesStr[0]) // dest
			ranges[l-1][1], _ = strconv.Atoi(rangesStr[1]) // source
			ranges[l-1][2], _ = strconv.Atoi(rangesStr[2]) // range length
		}
	}

	for i := 0; i < len(sources); i++ {
		s := sources[i]
		for _, mapping := range ranges {
			newSource := checkRange(s, mapping)
			if newSource != -1 {
				sources[i] = newSource
				break
			}
		}
	}
	min := math.MaxInt
	for _, s := range sources {
		min = int(math.Min(float64(s), float64(min)))
	}
	fmt.Println("PART 2 MINIMUM LOCATION ::: ", min)
}

func checkRange(source int, mapping [3]int) int { // return new source
	destCat, sourceCat, rlen := mapping[0], mapping[1], mapping[2]
	if source < sourceCat || source > sourceCat+rlen-1 {
		return -1
	}

	diff := destCat - sourceCat
	newSource := source + diff
	// fmt.Println("SOURCE ", source, "MAPPING ", mapping, "NEW SOURCE ", newSource)
	return newSource
}
