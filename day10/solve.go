package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var offsets = [][]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

var moves = []string{
	"|",
	"-",
	"L",
	"J",
	"7",
	"F",
}

func convertToMap(s *bufio.Scanner) (m [][]string, sx, sy int) {
	i := 0
	for s.Scan() {
		line := s.Text()
		row := strings.Split(line, "")
		m = append(m, row)
		for j, v := range row {
			if v == "S" {
				sy = j
				sx = i
			}
		}
		i++
	}
	return m, sx, sy
}

type Pos struct {
	x, y int
}

func findPath(m [][]string, sx, sy int) (path []Pos) {
	prevPos := []int{sx, sy}
	pos := []int{sx, sy}

	for _, v := range offsets {
		x := sx + v[0]
		y := sy + v[1]

		if x < 0 || x >= len(m) || y < 0 || y >= len(m[0]) {
			continue
		}
		if !slices.Contains(moves, m[x][y]) {
			continue
		}

		if m[x][y] == "-" && (y == sy) {
			continue
		} else if m[x][y] == "|" && (x == sx) {
			continue
		} else if m[x][y] == "L" && (x == sx+1 || y == sy-1) {
			continue
		} else if m[x][y] == "J" && (x == sx-1 || y == sy-1) {
			continue
		} else if m[x][y] == "7" && (x == sx-1 || y == sy+1) {
			continue
		} else if m[x][y] == "F" && (x == sx+1 || y == sy+1) {
			continue
		}

		pos = []int{x, y}
		break
	}

	for pos[0] != sx || pos[1] != sy {
		if !slices.Contains(moves, m[pos[0]][pos[1]]) {
			continue
		}

		move := ""

		if m[pos[0]][pos[1]] == "-" {
			if pos[1] < prevPos[1] {
				move = "left"
			} else {
				move = "right"
			}
		} else if m[pos[0]][pos[1]] == "|" {
			if pos[0] > prevPos[0] {
				move = "down"
			} else {
				move = "up"
			}
		} else if m[pos[0]][pos[1]] == "L" {
			if pos[1] == prevPos[1] {
				move = "right"
			} else {
				move = "up"
			}
		} else if m[pos[0]][pos[1]] == "J" {
			if pos[1] == prevPos[1] {
				move = "left"
			} else {
				move = "up"
			}
		} else if m[pos[0]][pos[1]] == "7" {
			if pos[1] == prevPos[1] {
				move = "left"
			} else {
				move = "down"
			}
		} else if m[pos[0]][pos[1]] == "F" {
			if pos[1] == prevPos[1] {
				move = "right"
			} else {
				move = "down"
			}
		}

		path = append(path, Pos{pos[0], pos[1]})

		if move == "up" {
			prevPos = pos
			pos = []int{pos[0] - 1, pos[1]}
		} else if move == "down" {
			prevPos = pos
			pos = []int{pos[0] + 1, pos[1]}
		} else if move == "left" {
			prevPos = pos
			pos = []int{pos[0], pos[1] - 1}
		} else if move == "right" {
			prevPos = pos
			pos = []int{pos[0], pos[1] + 1}
		}
	}

	return path
}

func part1(s *bufio.Scanner) (total int) {
	m, sx, sy := convertToMap(s)
	return len(findPath(m, sx, sy))/2 + 1
}

func part2(s *bufio.Scanner) (total int) {
	m, sx, sy := convertToMap(s)
	path := findPath(m, sx, sy)

	for i := range m {
		for j := range m[i] {
			if !slices.Contains(path, Pos{i, j}) && m[i][j] != "S" {
				m[i][j] = "."
			}

			if m[i][j] == "S" {
				if j-1 > 0 && j+1 < len(m[i]) && (m[i][j-1] == "7" || m[i][j-1] == "F" || m[i][j-1] == "-") && (m[i][j+1] == "L" || m[i][j+1] == "J" || m[i][j+1] == "-") {
					m[i][j] = "-"
				} else if i-1 > 0 && i+1 < len(m) && (m[i-1][j] == "F" || m[i-1][j] == "L" || m[i-1][j] == "|") && (m[i+1][j] == "7" || m[i+1][j] == "J" || m[i+1][j] == "|") {
					m[i][j] = "|"
				} else if i-1 > 0 && j+1 < len(m[i]) && (m[i-1][j] == "F" || m[i-1][j] == "L" || m[i-1][j] == "|") && (m[i][j+1] == "L" || m[i][j+1] == "J" || m[i][j+1] == "-") {
					m[i][j] = "7"
				} else if i-1 > 0 && j-1 > 0 && (m[i-1][j] == "F" || m[i-1][j] == "L" || m[i-1][j] == "|") && (m[i][j-1] == "7" || m[i][j-1] == "F" || m[i][j-1] == "-") {
					m[i][j] = "J"
				} else if i+1 < len(m) && j+1 < len(m[i]) && (m[i+1][j] == "7" || m[i+1][j] == "J" || m[i+1][j] == "|") && (m[i][j+1] == "L" || m[i][j+1] == "J" || m[i][j+1] == "-") {
					m[i][j] = "F"
				} else if i+1 < len(m) && j-1 > 0 && (m[i+1][j] == "7" || m[i+1][j] == "J" || m[i+1][j] == "|") && (m[i][j-1] == "7" || m[i][j-1] == "F" || m[i][j-1] == "-") {
					m[i][j] = "L"
				}
			}

			fmt.Println(m[sx][sy])

			if m[i][j] == "." {
				cnt := 0
				l7 := 1
				fj := 1

				for a := 0; a < j; a++ {
					if m[i][a] == "|" {
						cnt++
					}
					if m[i][a] == "7" || m[i][a] == "L" {
						l7++
					}
					if m[i][a] == "F" || m[i][a] == "J" {
						fj++
					}
				}

				cnt += l7 / 2
				cnt += fj / 2

				if cnt%2 == 1 {
					fmt.Println(i, j, m[i][j])
					total++
				}
			}
		}
	}

	return
}

func main() {
	f, err := os.Open("day10/input.txt")

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
