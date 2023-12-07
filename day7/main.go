package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards    []rune
	bid      int
	strength int
}

var CardMap = map[rune]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

var HandType = map[string]int{
	"High":  1,
	"One":   2,
	"Two":   3,
	"Three": 4,
	"Full":  5,
	"Four":  6,
	"Five":  7,
}

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
	contents := string(f)
	lines := strings.Split(contents, "\n")
	hands := []Hand{}
	for _, line := range lines {
		splitHand := strings.Split(line, " ")
		bid, _ := strconv.Atoi(splitHand[1])
		h := Hand{
			cards:    []rune(splitHand[0]),
			bid:      bid,
			strength: 0,
		}
		cMap := map[rune]int{}
		for _, c := range h.cards {
			if _, ok := cMap[c]; ok {
				cMap[c] += 1
			} else {
				cMap[c] = 1
			}
		}
		l := len(cMap)
		if l == 1 { // five of a kind
			h.strength = 7
		} else if l == 2 { // four of a kind or full house
			for _, v := range cMap {
				if v == 4 {
					h.strength = 6
					break
				} else {
					h.strength = 5
				}
			}

		} else if l == 3 { // three of a kind or two pair
			for _, v := range cMap {
				if v == 3 {
					h.strength = 4
					break
				} else {
					h.strength = 3
				}
			}

		} else if l == 4 { // one pair
			h.strength = 2
		} else if l == 5 { // high card
			h.strength = 1
		}

		hands = append(hands, h)
	}

	sortHands(hands)
	total := 0
	for i, h := range hands {
		// fmt.Println(h)
		score := (i + 1) * h.bid
		total += score
	}

	fmt.Println("PART 1 TOTAL ::: ", total)
}

func part2(fname string) {
	CardMap['J'] = 0
	f, _ := os.ReadFile(fname)
	contents := string(f)
	lines := strings.Split(contents, "\n")
	hands := []Hand{}
	for _, line := range lines {
		splitHand := strings.Split(line, " ")
		bid, _ := strconv.Atoi(splitHand[1])
		h := Hand{
			cards:    []rune(splitHand[0]),
			bid:      bid,
			strength: 0,
		}
		cMap := map[rune]int{}
		for _, c := range h.cards {
			// fmt.Printf("%c", c)
			if _, ok := cMap[c]; ok {
				cMap[c] += 1
			} else {
				cMap[c] = 1
			}
		}
		jCount := 0
		three, four, five := 0, 0, 0
		pairs := 0
		for letter, count := range cMap {
			if letter == 'J' {
				jCount = count
			} else if count == 3 {
				three += 1
			} else if count == 2 {
				pairs += 1
			} else if count == 4 {
				four += 1
			} else if count == 5 {
				five += 1
			}
		}
		if five == 1 {
			h.strength = HandType["Five"]
		} else if four == 1 {
			if jCount == 1 {
				h.strength = HandType["Five"]
			} else {
				h.strength = HandType["Four"]
			}
		} else if three == 1 {
			if jCount == 1 {
				h.strength = HandType["Four"]
			} else if jCount == 2 {
				h.strength = HandType["Five"]
			} else if pairs == 1 {
				h.strength = HandType["Full"]
			} else {
				h.strength = HandType["Three"]
			}
		} else if pairs == 2 {
			if jCount == 1 {
				h.strength = HandType["Full"]
			} else {
				h.strength = HandType["Two"]
			}
		} else if pairs == 1 {
			if jCount == 0 {
				h.strength = HandType["One"]
			} else if jCount == 1 {
				h.strength = HandType["Three"]
			} else if jCount == 2 {
				h.strength = HandType["Four"]
			} else if jCount == 3 {
				h.strength = HandType["Five"]
			}
		} else {
			if jCount == 0 {
				h.strength = HandType["High"]
			} else if jCount == 1 {
				h.strength = HandType["One"]
			} else if jCount == 2 {
				h.strength = HandType["Three"]
			} else if jCount == 3 {
				h.strength = HandType["Four"]
			} else if jCount == 4 {
				h.strength = HandType["Five"]
			} else if jCount == 5 {
				h.strength = HandType["Five"]
			}
		}

		hands = append(hands, h)
	}

	sortHands(hands)
	total := 0
	for i, h := range hands {
		// fmt.Println(string(h.cards), h.bid, h.strength)
		score := (i + 1) * h.bid
		total += score
	}

	fmt.Println("PART 2 TOTAL ::: ", total)
}

func sortHands(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool {
		// Compare by strength
		if hands[i].strength != hands[j].strength {
			return hands[i].strength < hands[j].strength
		}

		// If strength is equal, compare by cards
		for k, c1 := range hands[i].cards {
			c2 := hands[j].cards[k]
			if CardMap[c1] == CardMap[c2] {
				continue
			}
			if CardMap[c1] < CardMap[c2] {
				return true
			} else {
				return false
			}
		}
		return true
	})
}
