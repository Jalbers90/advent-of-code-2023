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

type Part map[string]int

type WorkFlow struct {
	partID string
	op     string
	num    int
	to     string
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
	workflows := map[string][]WorkFlow{}
	parts := []Part{}
	for _, line := range lines {
		if line == "" {
			continue
		} else if line[0] == '{' { // part
			partStr := line[1 : len(line)-1]
			partSplit := strings.Split(partStr, ",")
			part := Part{}
			for _, p := range partSplit {
				id := string(p[0])
				num, _ := strconv.Atoi(p[2:])
				part[id] = num
			}
			parts = append(parts, part)
		} else { // workflow
			curlyIndex := strings.Index(line, "{")
			name := line[0:curlyIndex]
			wfStrs := strings.Split(line[curlyIndex+1:len(line)-1], ",")
			workflows[name] = []WorkFlow{}
			for _, wfStr := range wfStrs {
				split := strings.Split(wfStr, ":")
				if len(split) > 1 {
					id, op := split[0][0], split[0][1]
					num, _ := strconv.Atoi(split[0][2:])
					to := split[1]
					workflows[name] = append(workflows[name], WorkFlow{partID: string(id), op: string(op), num: num, to: to})
				} else {
					workflows[name] = append(workflows[name], WorkFlow{to: split[0]})
				}

			}
		}
	}

	total := 0
	for _, p := range parts {
		cur := "in"
		for cur != "A" && cur != "R" {
			wfs := workflows[cur]
			for _, wf := range wfs {
				dlog("PART: %+v ::: CUR: %s ::: WORKFLOW %+v \n", p, cur, wf)
				if wf.partID == "" {
					cur = wf.to
					break
				} else {
					if compare(p[wf.partID], wf.num, wf.op) {
						cur = wf.to
						break
					}
				}
			}
		}
		if cur == "A" {
			total += sumPart(p)
		}
	}

	fmt.Println("PART 1 TOTAL :: ", total)
}

func part2(lines []string) {
	workflows := map[string][]WorkFlow{}
	for _, line := range lines {
		if line == "" {
			break
		}
		curlyIndex := strings.Index(line, "{")
		name := line[0:curlyIndex]
		wfStrs := strings.Split(line[curlyIndex+1:len(line)-1], ",")
		workflows[name] = []WorkFlow{}
		for _, wfStr := range wfStrs {
			split := strings.Split(wfStr, ":")
			if len(split) > 1 {
				id, op := split[0][0], split[0][1]
				num, _ := strconv.Atoi(split[0][2:])
				to := split[1]
				workflows[name] = append(workflows[name], WorkFlow{partID: string(id), op: string(op), num: num, to: to})
			} else {
				workflows[name] = append(workflows[name], WorkFlow{to: split[0]})
			}
		}
	}

	total := combos(workflows)

	// printWF(workflows)
	fmt.Println("PART 2 TOTAL ::: ", total)

}

func combos(workflows map[string][]WorkFlow) int {
	var dp func(part map[string][2]int, cur string) int
	dp = func(part map[string][2]int, cur string) int {
		if cur == "A" {
			dlog("CUR %s ::: PART: %+v\n", cur, part)
			return mulRanges(part) // get ranges and multiply together
		}
		if cur == "R" {
			return 0
		}
		dlog("PART: %+v ::: CUR %s\n", part, cur)
		total := 0
		workflow := workflows[cur]
		for _, wf := range workflow {
			total += dp(modifyPart(part, wf), wf.to)
			if wf.op == "<" {
				wf.op = ">"
				wf.num = wf.num - 1
			} else if wf.op == ">" {
				wf.op = "<"
				wf.num = wf.num + 1
			}
			part = modifyPart(part, wf)
		}
		return total
	}
	return dp(map[string][2]int{"x": {1, 4000}, "m": {1, 4000}, "a": {1, 4000}, "s": {1, 4000}}, "in")
}

func modifyPart(part map[string][2]int, wf WorkFlow) map[string][2]int {
	newPart := map[string][2]int{
		"x": {part["x"][0], part["x"][1]},
		"m": {part["m"][0], part["m"][1]},
		"a": {part["a"][0], part["a"][1]},
		"s": {part["s"][0], part["s"][1]},
	}
	if wf.partID == "" {
		return newPart
	}
	id := wf.partID
	num := wf.num
	op := wf.op
	if op == "<" {
		arr := newPart[id]
		if arr[1] >= num {
			arr[1] = num - 1
		}
		newPart[id] = arr
	} else if op == ">" {
		arr := newPart[id]
		if arr[0] <= num {
			arr[0] = num + 1
		}
		newPart[id] = arr
	}
	return newPart
}

func mulRanges(part map[string][2]int) int {
	prod := 1
	for _, v := range part {
		diff := v[1] - v[0] + 1 // high - low + 1
		if diff <= 0 {
			diff = 0
		}
		prod *= diff
	}
	return prod
}

func sumPart(part Part) int {
	sum := 0
	for _, v := range part {
		sum += v
	}
	return sum
}

func compare(a, b int, op string) bool {
	ops := map[string]func(int, int) bool{
		"<": func(a, b int) bool { return a < b },
		">": func(a, b int) bool { return a > b },
	}
	opFunc, ok := ops[op]
	if !ok {
		fmt.Printf("You forgot to add operator %s\n", op)
		return false
	}
	return opFunc(a, b)
}

func printWF(wofkflows map[string][]WorkFlow) {
	for k, v := range wofkflows {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func dlog(format string, args ...interface{}) {
	if DEBUG == 1 {
		log.Printf(format, args...)
	}
}
