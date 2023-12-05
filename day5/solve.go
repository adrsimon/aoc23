package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type parser struct {
	dest int
	orig int
	step int
}

type mapper struct {
	orig string
	dest string
}

func parse(s *bufio.Scanner) (maps map[mapper][]parser) {
	maps = make(map[mapper][]parser)
	var current mapper

	for s.Scan() {
		if s.Text() == "" {
			continue
		}

		if strings.Contains(s.Text(), "map") {
			str := strings.Split(strings.Fields(s.Text())[0], "-to-")
			m := mapper{}
			m.orig = str[0]
			m.dest = str[1]
			current = m
			maps[current] = []parser{}
		} else {
			p := parser{}
			p.dest, _ = strconv.Atoi(strings.Fields(s.Text())[0])
			p.orig, _ = strconv.Atoi(strings.Fields(s.Text())[1])
			p.step, _ = strconv.Atoi(strings.Fields(s.Text())[2])
			maps[current] = append(maps[current], p)
		}
	}

	return maps
}

func calculate(maps map[mapper][]parser, seeds []string) (lowest int) {
	var curr int
	var locations []int

	for _, seed := range seeds {
		state := "seed"
		curr, _ = strconv.Atoi(seed)
		for state != "location" {
			for m := range maps {
				if m.orig == state {
					for _, p := range maps[m] {
						if curr >= p.orig && curr <= p.orig+p.step {
							curr = p.dest + (curr - p.orig)
							break
						}
					}
					state = m.dest
					break
				}
			}
		}
		locations = append(locations, curr)
	}

	lowest = 1 << 31
	for _, loc := range locations {
		if loc < lowest {
			lowest = loc
		}
	}

	return
}

func part1(s *bufio.Scanner) (lowest int) {
	s.Scan()
	seeds := strings.Fields(strings.Split(s.Text(), ":")[1])

	maps := parse(s)
	lowest = calculate(maps, seeds)

	return lowest
}

func part2(s *bufio.Scanner) (total int) {
	s.Scan()
	seeding := strings.Fields(strings.Split(s.Text(), ":")[1])
	var seeds []string

	/**
	 * Very unoptimized, took a bunch of minutes to generate the seeds, might look into it later
	 */
	for i := 0; i < len(seeding); i += 2 {
		start, _ := strconv.Atoi(seeding[i])
		end, _ := strconv.Atoi(seeding[i+1])
		var partSeeds []string
		for j := start; j <= start+end; j++ {
			partSeeds = append(partSeeds, strconv.Itoa(j))
			fmt.Println(j)
		}
		seeds = append(seeds, partSeeds...)
	}

	fmt.Println(len(seeds))

	maps := parse(s)
	total = calculate(maps, seeds)

	return total
}

func main() {
	f, err := os.Open("day5/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	s := bufio.NewScanner(f)
	fmt.Println(part1(s))

	_, err = f.Seek(0, 0)
	if err != nil {
		return
	}
	s = bufio.NewScanner(f)
	fmt.Println(part2(s))

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
