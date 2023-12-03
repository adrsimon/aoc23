package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getNeighbors(xStart, xEnd, y int, input [][]string) (neighbors []string) {
	for i := y - 1; i <= y+1; i++ {
		for j := xStart - 1; j <= xEnd; j++ {
			if i == y && (j >= xStart && j <= xEnd-1) {
				continue
			}
			if i >= 0 && i < len(input) && j >= 0 && j < len(input[i]) {
				neighbors = append(neighbors, input[i][j])
			}
		}
	}
	return
}

func convert(s *bufio.Scanner) (slice [][]string) {
	for s.Scan() {
		temp := make([]string, 0)
		for _, v := range strings.Split(s.Text(), "") {
			temp = append(temp, v)
		}
		slice = append(slice, temp)
	}
	return slice
}

func part1(s *bufio.Scanner) (total int) {
	input := convert(s)

	for i := 0; i < len(input); i++ {
		num := ""
		numStart := 0
		for j := 0; j < len(input[i]); j++ {
			isDigit := unicode.IsDigit(rune(input[i][j][0]))
			if isDigit {
				if num == "" {
					numStart = j
				}
				num += input[i][j]
			}

			if !isDigit || j == len(input[i])-1 {
				if num != "" {
					neighbors := getNeighbors(numStart, j, i, input)
					for _, neighbor := range neighbors {
						if neighbor != "." {
							n, _ := strconv.Atoi(num)
							total += n
							break
						}
					}
				}
				num = ""
			}
		}
	}
	return
}

func part2(s *bufio.Scanner) (total int) {
	input := convert(s)
	dvals := []int{-1, 0, 1}

	gears := make([]int, 0)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] != "*" {
				continue
			}

			gears = gears[:0]
			for _, dx := range dvals {
				for _, dy := range dvals {
					x := j + dx
					y := i + dy

					if x < 0 || x >= len(input[i]) || y < 0 || y >= len(input) {
						continue
					}

					_, err := strconv.Atoi(input[y][x])
					if err == nil {
						for x > 0 && unicode.IsDigit(rune(input[y][x-1][0])) {
							x--
						}

						n := 0
						for x < len(input[y]) && unicode.IsDigit(rune(input[y][x][0])) {
							n = n*10 + int(input[y][x][0]-'0')
							x++
						}

						valid := true
						for _, gear := range gears {
							if gear == n {
								valid = false
								break
							}
						}

						if valid {
							gears = append(gears, n)
						}
					}
				}
			}

			if len(gears) == 2 {
				total += gears[0] * gears[1]
			}
		}
	}

	return
}

func main() {
	f, err := os.Open("day3/input.txt")

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
