package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var DEBUG = 0

type Step struct {
	label string
	focal int
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
	steps := strings.Split(contents, ",")

	value := 0
	total := 0
	for _, step := range steps {
		value = hash(step, value)
		// fmt.Println("STEP :: ", step, "VALUE ::", value)
		total += value
		value = 0
	}
	fmt.Println("PART 1 TOTAL :: ", total)
}

func part2(fname string) {
	f, _ := os.ReadFile(fname)
	contents := string(f)
	steps := strings.Split(contents, ",")
	boxes := map[int][]Step{}

	for _, step := range steps {
		op := "-"
		if strings.Contains(step, "=") {
			op = "="
		}
		opIndex := strings.Index(step, op)
		label := step[:opIndex]
		focal := 0
		if op == "=" {
			focal, _ = strconv.Atoi(step[opIndex+1:])
		}
		s := Step{label: label, focal: focal}
		box := hash(label, 0)

		if op == "=" {
			boxSlice := boxes[box]
			found := false
			for i, el := range boxSlice {
				if el.label == label {
					boxSlice[i] = s
					found = true
					break
				}
			}
			if !found {
				boxSlice = append(boxSlice, s)
			}
			boxes[box] = boxSlice
		} else {
			boxSlice := boxes[box]
			for i, el := range boxSlice {
				if el.label == label {
					boxSlice = append(boxSlice[:i], boxSlice[i+1:]...)
					break
				}
			}
			boxes[box] = boxSlice
		}
	}
	dlog("BOXES :: %+v", boxes)
	total := 0
	for bNo, bSlice := range boxes {
		for i, s := range bSlice {
			total += (bNo + 1) * (i + 1) * s.focal
		}
	}
	fmt.Println("PART 2 TOTAL :: ", total)
}

func dlog(format string, args ...interface{}) {
	if DEBUG == 1 {
		log.Printf(format, args...)
	}
}

func hash(step string, value int) int {
	runes := []rune(step)
	for _, r := range runes {
		value += int(r)
		value *= 17
		value = value % 256
	}
	return value
}
