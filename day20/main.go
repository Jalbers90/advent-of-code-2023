package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var DEBUG = 0

// Flip-flop modules (prefix %) are either on or off;
// they are initially off.
// If a flip-flop module receives a high pulse, it is ignored and nothing happens.
// However, if a flip-flop module receives a low pulse, it flips between on and off.
// If it's on it sends a high pulse. If it's off it sends a low pulse

// Conjunction modules (prefix &) remember the type of the most recent pulse received from each of their connected input modules;
// they initially default to remembering a low pulse for each input.
// When a pulse is received, the conjunction module first updates its memory for that input.
// Then, if it remembers high pulses for all inputs, it sends a low pulse; otherwise, it sends a high pulse.

// Broadcast module.
// When it receives a pulse, it sends the same pulse to all of its destination modules.

type Module struct {
	Type         byte
	on           bool            // false off true on
	recent       map[string]byte // "L", "H"
	destinations []string
	id           string
}

type Configuration map[string]Module

// type FFSignals map[byte]int {
// 	'L':
// }

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
		fname = "test2.txt"
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
	config := Configuration{}
	for _, line := range lines {
		split := strings.Split(line, " -> ")
		moduleStr, destinationStr := split[0], split[1]
		module := Module{}
		destinations := strings.Split(destinationStr, ", ")
		module.destinations = destinations
		if moduleStr[0] == '%' {
			module.Type = '%' // flip-flop
			module.id = moduleStr[1:]
			config[moduleStr[1:]] = module
		} else if moduleStr[0] == '&' {
			module.Type = '&' // conjunction
			module.recent = map[string]byte{}
			// for _, d := range destinations {
			// 	module.recent[d] = 'L'
			// }
			module.id = moduleStr[1:]
			config[moduleStr[1:]] = module
		} else {
			module.Type = 'B' // broadcaster
			module.id = "broadcaster"
			config["broadcaster"] = module
		}
	}

	for id, mod := range config {
		for _, d := range mod.destinations {
			dest := config[d]
			if dest.Type == '&' {
				dest.recent[id] = 'L'
			}
		}
	}

	lows := 0
	highs := 0
	for i := 0; i < 1000; i++ {
		lows += 1
		sig := byte('L')
		q := []Module{config["broadcaster"]}
		for len(q) > 0 {
			dlog("%+v\n", config)
			cur := q[0]
			q = q[1:]
			sig = setSignal(cur, sig)
			dlog("CURRENT MODULE: %+v ::: SIGNAL: %c\n", cur, sig)
			for _, d := range cur.destinations {
				lows, highs = incrementTotal(lows, highs, sig)
				dest, ok := config[d]
				if !ok {
					continue
				}
				dlog("DESTINATION MODULE: %s ::: SIGNAL: %c ::: LOWS: %d ::: HIGHS: %d\n", dest.id, sig, lows, highs)
				if dest.Type == '%' {
					if sig == 'L' {
						dest.on = !dest.on
						config[d] = dest
						q = append(q, config[d])
					}
				} else if dest.Type == '&' {
					dest.recent[cur.id] = sig
					config[d] = dest
					q = append(q, config[d])
				}
			}
		}
	}

	fmt.Printf("1 BUTTON PRESS ::: LOWS: %d ::: HIGHS: %d\n", lows, highs)
	fmt.Println("PART 1 TOTAL ::::", lows*highs)
}

func part2(lines []string) {
	config := Configuration{}
	for _, line := range lines {
		split := strings.Split(line, " -> ")
		moduleStr, destinationStr := split[0], split[1]
		module := Module{}
		destinations := strings.Split(destinationStr, ", ")
		module.destinations = destinations
		if moduleStr[0] == '%' {
			module.Type = '%' // flip-flop
			module.id = moduleStr[1:]
			config[moduleStr[1:]] = module
		} else if moduleStr[0] == '&' {
			module.Type = '&' // conjunction
			module.recent = map[string]byte{}
			// for _, d := range destinations {
			// 	module.recent[d] = 'L'
			// }
			module.id = moduleStr[1:]
			config[moduleStr[1:]] = module
		} else {
			module.Type = 'B' // broadcaster
			module.id = "broadcaster"
			config["broadcaster"] = module
		}
	}

	for id, mod := range config {
		for _, d := range mod.destinations {
			dest := config[d]
			if dest.Type == '&' {
				dest.recent[id] = 'L'
			}
		}
	}

	sources := map[string]int{} // sources that feed into ls
	found := false
	i := 0
	for {
		i++
		sig := byte('L')
		q := []Module{config["broadcaster"]}
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			sig = setSignal(cur, sig)
			for _, d := range cur.destinations {
				dest, ok := config[d]
				if !ok {
					continue
				}
				if d == "ls" && sig == 'H' {
					if _, ok := sources[cur.id]; ok {
						found = true
						break
					}
					sources[cur.id] = i
				}
				if dest.Type == '%' {
					if sig == 'L' {
						dest.on = !dest.on
						config[d] = dest
						q = append(q, config[d])
					}
				} else if dest.Type == '&' {
					dest.recent[cur.id] = sig
					config[d] = dest
					q = append(q, config[d])
				}
			}
		}
		if found {
			break
		}
	}

	presses := 1
	for _, v := range sources {
		presses *= v
	}
	fmt.Println("PART 2 BTN PRESSES ::::", presses)
}

func incrementTotal(lows, highs int, sig byte) (int, int) {
	if sig == 'L' {
		lows += 1
	} else {
		highs += 1
	}
	return lows, highs
}

func setSignal(cur Module, sig byte) byte {
	if cur.Type == '%' {
		if cur.on == true {
			return 'H'
		} else {
			return 'L'
		}
	} else if cur.Type == '&' {
		for _, r := range cur.recent {
			if r == 'L' {
				return 'H'
			}
		}
		return 'L'
	} else {
		return sig
	}
}

func dlog(format string, args ...interface{}) {
	if DEBUG == 1 {
		log.Printf(format, args...)
	}
}
