package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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
	if *p1 {
		part1(fname)
	} else if *p2 {
		part2(fname)
	}
}

func part1(fname string) {
	f, _ := os.ReadFile(fname)
	fStr := string(f)
	lines := strings.Split(fStr, "\n")
	total := 0
	for i, line := range lines {
		prefix := fmt.Sprintf("Card %d: ", i+1)
		trimmed := strings.TrimPrefix(line, prefix)
		card := strings.Split(trimmed, " | ")
		winningStr, myNumsStr := card[0], card[1]
		winning := strings.Split(winningStr, " ")
		myNums := strings.Split(myNumsStr, " ")
		score := 0
		for _, num := range myNums {
			if num == "" || num == " " {
				continue
			}
			if containsElement(winning, num) {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		total += score
	}
	fmt.Println("PART 1 TOTAL ::: ", total)
}

func part2(fname string) {
	f, _ := os.ReadFile(fname)
	fStr := string(f)
	lines := strings.Split(fStr, "\n")
	cardCount := make([]int, len(lines))
	for i := range cardCount {
		cardCount[i] = 1
	}

	for i, line := range lines {
		prefix := fmt.Sprintf("Card %d: ", i+1)
		trimmed := strings.TrimPrefix(line, prefix)
		card := strings.Split(trimmed, " | ")
		winningStr, myNumsStr := card[0], card[1]
		winning := strings.Split(winningStr, " ")
		myNums := strings.Split(myNumsStr, " ")
		score := 0
		for _, num := range myNums {
			if num == "" || num == " " {
				continue
			}
			if containsElement(winning, num) {
				score += 1
			}
		}
		multiplier := cardCount[i]
		for j := 1; j <= score; j++ {
			cardCount[i+j] = cardCount[i+j] + multiplier
		}
	}

	total := 0
	for _, num := range cardCount {
		// fmt.Println("CardCount Num ", num)
		total += num
	}

	fmt.Println("PART 2 TOTAL ::: ", total)
}

func containsElement(slice []string, target string) bool {
	for _, element := range slice {
		if element == target {
			return true
		}
	}
	return false
}
