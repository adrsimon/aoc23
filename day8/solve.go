package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func pgcd(a, b int) int {
	if b == 0 {
		return a
	}
	return pgcd(b, a%b)
}

func ppcm(a, b int) int {
	return a * b / pgcd(a, b)
}

func ppcmSlice(s []int) int {
	if len(s) == 1 {
		return s[0]
	}
	return ppcm(s[0], ppcmSlice(s[1:]))
}

func part1(s *bufio.Scanner) (total int) {
	s.Scan()
	moves := strings.Split(s.Text(), "")
	s.Scan()

	tree := make(map[string][]string)
	for s.Scan() {
		line := strings.Split(s.Text(), " = ")
		dests := line[1][1 : len(line[1])-1]
		tree[line[0]] = strings.Split(dests, ", ")
	}

	curr := "AAA"

	for curr != "ZZZ" {
		for _, move := range moves {
			if move == "L" {
				curr = tree[curr][0]
			} else {
				curr = tree[curr][1]
			}
			total += 1

			if curr == "ZZZ" {
				break
			}
		}
	}

	return total
}

func part2(s *bufio.Scanner) (total int) {
	s.Scan()
	moves := strings.Split(s.Text(), "")
	s.Scan()

	tree := make(map[string][]string)
	var anodes, znodes []string

	for s.Scan() {
		line := strings.Split(s.Text(), " = ")
		dests := line[1][1 : len(line[1])-1]
		tree[line[0]] = strings.Split(dests, ", ")
		if line[0][2] == 'A' {
			anodes = append(anodes, line[0])
		}
		if line[0][2] == 'Z' {
			znodes = append(znodes, line[0])
		}
	}

	var cnts []int
	for _, anode := range anodes {
		cnt := 0
		curr := anode
		for !slices.Contains(znodes, curr) {
			for _, move := range moves {
				if move == "L" {
					curr = tree[curr][0]
				} else {
					curr = tree[curr][1]
				}
				cnt += 1

				if slices.Contains(znodes, curr) {
					break
				}
			}
		}
		cnts = append(cnts, cnt)
	}

	total = ppcmSlice(cnts)

	return total
}

func main() {
	f, err := os.Open("day8/input.txt")

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
