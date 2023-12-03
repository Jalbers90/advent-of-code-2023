package main

import "flag"

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

}

func part2(fname string) {

}
