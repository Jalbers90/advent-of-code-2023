package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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
	maxColorMap := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}
	contentBytes, _ := os.ReadFile(fname)
	contentStr := string(contentBytes)
	lines := strings.Split(contentStr, "\n")
	total := 0
	for i, line := range lines {
		id := i + 1
		gameStart := strings.Index(line, ":") + 1
		gameStr := line[gameStart:]
		game := strings.Split(gameStr, ";")
		possible := true
		for _, set := range game { // each game (line) splits into sets
			cubes := strings.Split(set, ",") // each set has 3 cubes
			for _, cube := range cubes {     // each cube has score and color
				each := strings.Split(cube, " ")
				num, _ := strconv.Atoi(each[1]) // score of cube
				color := each[2]                // color of cube
				if num > maxColorMap[color] {
					possible = false
				}
			}
			if !possible {
				break
			}
		}
		if possible {
			total += id
		}
	}
	fmt.Printf("PART 1 SOLUTION ::: ID TOTAL ::: %d\n", total)
}

func part2(fname string) {
	contentBytes, _ := os.ReadFile(fname)
	contentStr := string(contentBytes)
	lines := strings.Split(contentStr, "\n")
	total := 0
	for _, line := range lines {
		// id := i + 1
		gameStart := strings.Index(line, ":") + 1
		gameStr := line[gameStart:]
		game := strings.Split(gameStr, ";")
		maxMap := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		for _, set := range game { // each game (line) splits into sets
			cubes := strings.Split(set, ",") // each set has 3 cubes
			for _, cube := range cubes {     // each cube has score and color
				each := strings.Split(cube, " ")
				num, _ := strconv.Atoi(each[1]) // score of cube
				color := each[2]
				if num > maxMap[color] {
					maxMap[color] = num
				}
			}
		}

		// range max map
		product := 1
		for _, v := range maxMap {
			product *= v
		}
		total += product
	}
	fmt.Printf("PART 2 SOLUTION ::: SUM OF POWER ::: %d\n", total)
}
