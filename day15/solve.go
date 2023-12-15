package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func hash(s string) (total int) {
	for _, c := range s {
		total += int(c)
		total *= 17
		total = total % 256
	}
	return total
}

func part1(s *bufio.Scanner) (total int) {
	input := make([]string, 0)
	for s.Scan() {
		input = append(input, strings.Split(s.Text(), ",")...)
	}

	for _, str := range input {
		total += hash(str)
	}

	return total
}

type lens struct {
	tag   string
	focal int
}

func part2(s *bufio.Scanner) (total int) {
	input := make([]string, 0)
	for s.Scan() {
		input = append(input, strings.Split(s.Text(), ",")...)
	}

	boxes := make(map[int][]lens)

	for _, str := range input {
		if len(strings.Split(str, "=")) == 2 {
			l := strings.Split(str, "=")
			hash := hash(l[0])
			focal, _ := strconv.Atoi(l[1])
			k := len(boxes[hash])
			for i, f := range boxes[hash] {
				if f.tag == l[0] {
					boxes[hash] = append(boxes[hash][:i], boxes[hash][i+1:]...)
					k = i
				}
			}
			boxes[hash] = slices.Insert(boxes[hash], k, lens{tag: l[0], focal: focal})
		} else {
			l := strings.Split(str, "-")
			hash := hash(l[0])
			for i, f := range boxes[hash] {
				if f.tag == l[0] {
					boxes[hash] = append(boxes[hash][:i], boxes[hash][i+1:]...)
				}
			}
		}
	}

	for k, l := range boxes {
		for i, f := range l {
			total += (k + 1) * (i + 1) * f.focal
		}
	}

	return total
}

func main() {
	f, err := os.Open("day15/input.txt")

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
