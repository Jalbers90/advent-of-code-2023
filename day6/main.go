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
	timesCluttered := (strings.Split(strings.TrimPrefix(lines[0], "Time:      "), " "))
	distsCluttered := strings.Split(strings.TrimPrefix(lines[1], "Distance:  "), " ")
	times := []int{}
	dists := []int{}
	for i := 0; i < len(timesCluttered); i++ {
		t, err := strconv.Atoi(timesCluttered[i])
		if err != nil {
			continue
		} else {
			times = append(times, t)
		}
	}
	for i := 0; i < len(distsCluttered); i++ {
		d, err := strconv.Atoi(distsCluttered[i])
		if err != nil {
			continue
		} else {
			dists = append(dists, d)
		}
	}

	l := len(times)
	wins := make([]int, l)
	for i := 0; i < l; i++ {
		time := times[i]
		dist := dists[i] // dist to beat
		for t := 1; t < time+1; t++ {
			if t*(time-t) > dist {
				wins[i] += 1
			}
		}
	}

	total := wins[0]
	// multiply wins up
	for i := 1; i < len(wins); i++ {
		total *= wins[i]
	}

	fmt.Println("PART 1 TOTAL ::: ", total)
}

func part2(fname string) {
	f, _ := os.ReadFile(fname)
	contents := string(f)
	lines := strings.Split(contents, "\n")
	timeStr := strings.ReplaceAll(strings.TrimPrefix(lines[0], "Time:      "), " ", "")
	distStr := strings.ReplaceAll(strings.TrimPrefix(lines[1], "Distance:  "), " ", "")

	time, err := strconv.Atoi(timeStr)
	if err != nil {
		log.Fatal("Failed at time strconv.Atoi() ", err)
	}
	dist, _ := strconv.Atoi(distStr)

	wins := 0
	for t := 1; t < time+1; t++ {
		if t*(time-t) > dist {
			wins += 1
		}
	}

	fmt.Println("PART 2 TOTAL ::: ", wins)
}
