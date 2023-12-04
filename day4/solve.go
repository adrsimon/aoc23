package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(s *bufio.Scanner) (total int) {
	for s.Scan() {
		line := strings.Split(s.Text(), ":")[1]
		parts := strings.Split(line, "|")
		winners := strings.Fields(parts[0])
		mines := strings.Fields(parts[1])

		minesMap := make(map[string]int)
		for _, num := range mines {
			minesMap[num]++
		}

		matches := 0
		for _, num := range winners {
			if count, exists := minesMap[num]; exists && count > 0 {
				matches++
				minesMap[num]--
			}
		}

		if matches == 0 {
			continue
		}
		points := 1 << (matches - 1)
		total += points
	}

	return total
}

func part2(s *bufio.Scanner) (total int) {
	cnt := map[int]int{}

	for s.Scan() {
		line := strings.Split(s.Text(), ":")
		parts := strings.Split(line[1], "|")
		winners := strings.Fields(parts[0])
		mines := strings.Fields(parts[1])

		id, err := strconv.Atoi(strings.Fields(line[0])[1])
		if err != nil {
			log.Fatal(err)
		}
		cnt[id] += 1

		minesMap := make(map[string]int)
		for _, num := range mines {
			minesMap[num]++
		}

		matches := 0
		for _, num := range winners {
			if count, exists := minesMap[num]; exists && count > 0 {
				matches++
				minesMap[num]--
			}
		}

		for i := 1; i <= matches; i++ {
			cnt[id+i] += cnt[id]
		}
	}

	for _, v := range cnt {
		if v > 0 {
			total += v
		}
	}

	return
}

func main() {
	f, err := os.Open("day4/input.txt")

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
