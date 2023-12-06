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
	total = 1

	s.Scan()
	times := strings.Fields(strings.Split(s.Text(), ":")[1])
	s.Scan()
	distances := strings.Fields(strings.Split(s.Text(), ":")[1])

	for i := 0; i < len(times); i++ {
		sub := 0
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		for wait := 0; wait < time; wait++ {
			if wait*(time-wait) > distance {
				sub += 1
			}
		}
		total *= sub
	}

	return total
}

func part2(s *bufio.Scanner) (total int) {
	s.Scan()
	times := strings.Fields(strings.Split(s.Text(), ":")[1])
	s.Scan()
	distances := strings.Fields(strings.Split(s.Text(), ":")[1])
	var strTime, strDistance string

	for i := 0; i < len(times); i++ {
		strTime += times[i]
		strDistance += distances[i]
	}
	time, _ := strconv.Atoi(strTime)
	distance, _ := strconv.Atoi(strDistance)

	for wait := 0; wait < time; wait++ {
		if wait*(time-wait) > distance {
			total += 1
		}
	}

	return total
}

func main() {
	f, err := os.Open("day6/input.txt")

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
