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
	thresholds := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for s.Scan() {
		game := strings.Split(s.Text(), ":")
		stringId := strings.Split(game[0], " ")[1]
		id, err := strconv.Atoi(stringId)
		if err != nil {
			log.Fatal(err)
		}
		sets := strings.Split(game[1], ";")
		valid := true

		for set := range sets {
			if !valid {
				break
			}

			sets[set] = strings.TrimSpace(sets[set])
			cols := strings.Split(sets[set], ",")
			for col := range cols {
				cols[col] = strings.TrimSpace(cols[col])
				split := strings.Split(cols[col], " ")
				val, err := strconv.Atoi(split[0])
				if err != nil {
					log.Fatal(err)
				}
				if thresholds[split[1]] < val {
					valid = false
				}
			}
		}
		if valid {
			total += id
		}
	}
	return
}

func part2(s *bufio.Scanner) (total int) {
	pows := make([]int, 0)

	for s.Scan() {
		game := strings.Split(s.Text(), ":")

		maxs := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		sets := strings.Split(game[1], ";")
		for set := range sets {
			sets[set] = strings.TrimSpace(sets[set])
			cols := strings.Split(sets[set], ",")

			for col := range cols {
				cols[col] = strings.TrimSpace(cols[col])
				split := strings.Split(cols[col], " ")
				val, err := strconv.Atoi(split[0])
				if err != nil {
					log.Fatal(err)
				}
				if maxs[split[1]] < val {
					maxs[split[1]] = val
				}
			}
		}

		pows = append(pows, maxs["red"]*maxs["green"]*maxs["blue"])
	}

	for _, pow := range pows {
		total += pow
	}
	return
}

func main() {
	f, err := os.Open("day2/input.txt")

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

	_, err = f.Seek(0, 0)
	if err != nil {
		return
	}
	s = bufio.NewScanner(f)

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
