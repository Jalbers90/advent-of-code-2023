package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("::: START OF MAIN :::")
	test := flag.Bool("test", false, "if true, use test.txt")
	flag.Parse()
	fname := "input_1.txt"
	if *test {
		fname = "test.txt"
	}
	// part1(fname)

	if *test {
		fname = "test2.txt"
	}
	part2(fname)
}

// find first and last digit in each line
func part1(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		first := "" // zero value for rune
		second := ""
		for _, c := range line {
			if unicode.IsDigit(c) {
				if first == "" {
					first = string(c)
				} else {
					second = string(c)
				}
			}
		}
		if second == "" {
			second = first
		}
		num, err := strconv.Atoi(first + second)
		if err != nil {
			log.Fatal("fail str conversion :::: ", err)
		}
		// fmt.Printf("LINE %s :::: NUM ::: %d\n", line, num)
		total += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PART 1 - Sum of all calibrations ::: %d\n", total)
}

func part2(fname string) {
	words := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	nums := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
		"4": "4",
		"5": "5",
		"6": "6",
		"7": "7",
		"8": "8",
		"9": "9",
	}

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		fwi, lwi, fdi, ldi := math.MaxInt, -1, math.MaxInt, -1 // first/last word/digit index
		fw, lw, fd, ld := "", "", "", ""
		for k, _ := range words {
			first := strings.Index(line, k)
			last := strings.LastIndex(line, k)
			if first != -1 && first < fwi {
				fwi = first
				fw = k
			}
			if last != -1 && last > lwi {
				lwi = last
				lw = k
			}
		}

		for k, _ := range nums {
			first := strings.Index(line, k)
			last := strings.LastIndex(line, k)
			if first != -1 && first < fdi {
				fdi = first
				fd = k
			}
			if last != -1 && last > ldi {
				ldi = last
				ld = k
			}
		}
		str := ""
		if fdi < fwi {
			str += fd
		} else {
			str += words[fw]
		}

		if ldi > lwi {
			str += ld
		} else {
			str += words[lw]
		}

		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal("failed strconv ::: ", err)
		}
		fmt.Printf("LINE %s :::: NUM ::: %d\n", line, num)
		total += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("PART 2 - Sum of all calibrations ::: %d\n", total)
}
