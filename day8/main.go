package main

import (
	"flag"
	"fmt"
	"os"
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
		fname = "test3.txt"
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
	nodesStrs := lines[2:]
	nodes := map[string][2]string{}

	for _, line := range nodesStrs {
		// AAA = (BBB, CCC)
		// 01234567
		node := line[:3]
		left := line[7:10]
		right := line[12:15]
		nodes[node] = [2]string{left, right}
	}

	dirs := lines[0]
	m := len(dirs)
	i := 0
	steps := 0
	cur := "AAA"
	for {
		steps += 1
		RorL := dirs[i]
		dir := 0
		if RorL == 'L' {
			dir = 0 // left
		} else {
			dir = 1
		}
		cur = nodes[cur][dir]
		if cur == "ZZZ" {
			break
		}

		i = (i + 1) % m
	}
	fmt.Println("PART 1 STEPS ::: ", steps)
}

func part2(fname string) {
	f, _ := os.ReadFile(fname)
	contents := string(f)
	lines := strings.Split(contents, "\n")
	nodesStrs := lines[2:]
	nodes := map[string][2]string{}
	curs := []string{} // current nodes
	for _, line := range nodesStrs {
		node := line[:3]
		left := line[7:10]
		right := line[12:15]
		nodes[node] = [2]string{left, right}
		if node[2] == 'A' {
			curs = append(curs, node)
		}
	}

	dirs := lines[0]
	m := len(dirs)
	i := 0
	stepArr := make([]int, len(curs))
	steps := 0
	for j, cur := range curs {
		for {
			steps += 1
			RorL := dirs[i]
			dir := 0
			if RorL == 'L' {
				dir = 0 // left
			} else {
				dir = 1
			}
			cur = nodes[cur][dir]
			// fmt.Println("CURRENT ", cur)
			if cur[2] == 'Z' {
				break
			}
			// if steps > 10 {
			// 	break
			// }
			i = (i + 1) % m
		}
		stepArr[j] = steps
		steps = 0
		i = 0
	}
	total := lcm(stepArr)
	fmt.Println("PART 1 TOTAL STEPS ::: ", total)
}

func lcm(nums []int) int {
	result := 1

	for _, num := range nums {
		// least common multiple(a, b) = abs(a * b) / gcd(a, b)
		product := (result * num) / gcd(result, num)
		result = product
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
