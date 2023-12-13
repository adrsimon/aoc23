package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func solve(field []string, p2 bool) int {
	w := len(field[0])
	h := len(field)

	for c := 0; c < h-1; c++ {
		diff := 0
		for col := 0; col < w; col++ {
			for offset := 0; ; offset++ {
				over := c - offset
				under := c + offset + 1
				if over < 0 || under >= h {
					break
				}
				if field[over][col] != field[under][col] {
					diff++
				}
			}
		}
		if (p2 && diff == 1) || (!p2 && diff == 0) {
			return c + 1
		}
	}
	return 0
}

func part1(s *bufio.Scanner) (total int) {
	fields := make([][]string, 0)

	i := 0
	for s.Scan() {
		line := s.Text()
		if line == "" {
			i++
			continue
		}
		if len(fields) <= i {
			fields = append(fields, make([]string, 0))
		}
		fields[i] = append(fields[i], line)
	}

	for _, f := range fields {
		total += 100 * solve(f, false)

		transposed := make([]string, len(f[0]))

		for j := 0; j < len(f[0]); j++ {
			var col string
			for i := 0; i < len(f); i++ {
				col += string(f[i][j])
			}
			transposed[j] = col
		}
		total += solve(transposed, false)
	}

	return total
}

func part2(s *bufio.Scanner) (total int) {
	fields := make([][]string, 0)

	i := 0
	for s.Scan() {
		line := s.Text()
		if line == "" {
			i++
			continue
		}
		if len(fields) <= i {
			fields = append(fields, make([]string, 0))
		}
		fields[i] = append(fields[i], line)
	}

	for _, f := range fields {
		total += 100 * solve(f, true)

		transposed := make([]string, len(f[0]))

		for j := 0; j < len(f[0]); j++ {
			var col string
			for i := 0; i < len(f); i++ {
				col += string(f[i][j])
			}
			transposed[j] = col
		}
		total += solve(transposed, true)
	}

	return total
}

func main() {
	f, err := os.Open("day13/input.txt")

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
