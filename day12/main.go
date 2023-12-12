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
	total := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		recordStr, listStr := split[0], split[1]
		listArr := strings.Split(listStr, ",")
		list := make([]int, len(listArr)) // contiguous group of damaged springs
		for j, c := range listArr {
			n, _ := strconv.Atoi(c)
			list[j] = n
		}
		ways := arrangements(recordStr, list)
		total += ways
	}
	fmt.Println("PART 1 TOTAL :: ", total)
}

func part2(lines []string) {
	total := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		recordStr, listStr := split[0], split[1]
		listArr := strings.Split(listStr, ",")
		list := make([]int, len(listArr)) // contiguous group of damaged springs
		for j, c := range listArr {
			n, _ := strconv.Atoi(c)
			list[j] = n
		}
		unfoldRecord, unfoldList := unfold(recordStr, list)
		ways := arrangements(unfoldRecord, unfoldList)
		total += ways
	}
	fmt.Println("PART 2 TOTAL :: ", total)
}

func unfold(record string, list []int) (string, []int) {
	unfoldStr := record
	for i := 0; i < 4; i++ {
		unfoldStr += "?" + record
	}

	unfoldList := make([]int, len(list)*5)
	for i := 0; i < 5; i++ {
		for j, n := range list {
			unfoldList[j+len(list)*i] = n
		}
	}
	return unfoldStr, unfoldList
}

func arrangements(record string, list []int) int {
	memo := map[string]int{}
	var dp func(i, j int) int

	dp = func(i, j int) int {
		if i >= len(record) {
			if j == len(list) {
				return 1
			}
			return 0
		}
		key := fmt.Sprintf("%d,%d", i, j)
		if v, ok := memo[key]; ok {
			return v
		}
		ways := 0
		char := record[i]
		if char == '.' {
			ways += dp(i+1, j)
		} else {
			// either '?' or '#'
			if char == '?' {
				ways += dp(i+1, j) // place '.'
			}

			if j < len(list) {
				count := list[j]
				newI := i
				for newI < len(record) && count > 0 {
					char = record[newI]
					if char == '?' || char == '#' {
						count -= 1
					} else if char == '.' {
						break
					}
					newI += 1
				}

				if count == 0 {
					if newI == len(record) {
						ways += dp(newI, j+1)
					} else if record[newI] == '.' || record[newI] == '?' {
						ways += dp(newI+1, j+1) // +1 because after placing '#' we must place a '.'
					}
				}
			}
		}
		memo[key] = ways
		return memo[key]
	}

	return dp(0, 0)
}
