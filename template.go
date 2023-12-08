package main

import (
	"flag"
	"fmt"
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
	// f, _ := os.ReadFile(fname)
	// contents := string(f)
	// lines := strings.Split(contents, "\n")
}

func part2(fname string) {
	// f, _ := os.ReadFile(fname)
	// contents := string(f)
	// lines := strings.Split(contents, "\n")
}
