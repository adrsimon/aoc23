package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func cosmicExpand(s *bufio.Scanner) (space []string, galaxies [][]int, galCnt int, rows, cols []bool) {
	space = make([]string, 0)
	for s.Scan() {
		space = append(space, s.Text())
	}

	rows = make([]bool, len(space))
	cols = make([]bool, len(space[0]))
	galCnt = 0
	galaxies = make([][]int, len(space))
	for i := range galaxies {
		galaxies[i] = make([]int, len(space[0]))

		for j := range galaxies[i] {
			galaxies[i][j] = -1
		}
	}

	for i := 0; i < len(space); i++ {
		expanded := true
		for j := 0; j < len(space[0]); j++ {
			if space[i][j] != '.' {
				expanded = false
				galaxies[i][j] = galCnt
				galCnt++
			}
		}

		rows[i] = expanded
	}

	for j := 0; j < len(space[0]); j++ {
		expanded := true
		for i := 0; i < len(space); i++ {
			if space[i][j] != '.' {
				expanded = false
				break
			}
		}

		cols[j] = expanded
	}

	return space, galaxies, galCnt, rows, cols
}

func part1(s *bufio.Scanner) (total int) {
	space, galaxies, cnt, rows, cols := cosmicExpand(s)

	for i := range galaxies {
		for j := range galaxies[i] {
			if galaxies[i][j] == -1 {
				continue
			}

			total += bfs(space, Pos{i, j}, galaxies, cnt, rows, cols, 2)
			galaxies[i][j] = -1
		}
	}

	return total
}

func part2(s *bufio.Scanner) (total int) {
	space, galaxies, cnt, expandedRows, expandedCols := cosmicExpand(s)

	for i := range galaxies {
		for j := range galaxies[i] {
			if galaxies[i][j] == -1 {
				continue
			}

			total += bfs(space, Pos{i, j}, galaxies, cnt, expandedRows, expandedCols, 1000000)
			galaxies[i][j] = -1
		}
	}

	return total
}

type Pos struct {
	i, j int
}

type PosAndPath struct {
	Pos
	curLen int
}

var offsets = []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func bfs(space []string, start Pos, ends [][]int, galaxiesCnt int, rows, cols []bool, delta int) int {
	delta--

	var queue = []PosAndPath{{Pos: start, curLen: 0}}
	visited := make([][]bool, len(space))
	for i := range visited {
		visited[i] = make([]bool, len(space[0]))
	}

	distances := make([]int, galaxiesCnt)
	for i := range distances {
		distances[i] = math.MaxInt
	}

	distances[ends[start.i][start.j]] = 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dir := range offsets {
			newC := PosAndPath{Pos: Pos{cur.i + dir.i, cur.j + dir.j}, curLen: cur.curLen + 1}
			if newC.i < 0 || newC.i >= len(space) || newC.j < 0 || newC.j >= len(space[0]) {
				continue
			}

			if visited[newC.i][newC.j] {
				continue
			}

			if rows[newC.i] {
				newC.curLen += delta
			}

			if cols[newC.j] {
				newC.curLen += delta
			}

			if v := ends[newC.i][newC.j]; v != -1 {
				distances[v] = min(distances[v], newC.curLen)
			}

			visited[newC.i][newC.j] = true
			queue = append(queue, newC)
		}
	}

	total := 0
	for _, d := range distances {
		if d == math.MaxInt {
			continue
		}
		total += d
	}

	return total
}

func main() {
	f, err := os.Open("day11/input.txt")

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
