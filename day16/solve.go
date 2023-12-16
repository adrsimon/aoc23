package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

type tbeam struct {
	pos []int
	dir string
}

func part1(field []string, dir string, pos []int) (total int) {
	beams := make([]tbeam, 0)
	originalBeam := tbeam{dir: dir, pos: pos}
	beams = append(beams, originalBeam)
	var vis []string
	var energized []string

	for len(beams) > 0 {
		beam := beams[0]
		beams = beams[1:]
		for !slices.Contains(vis, beam.dir+"-"+strconv.Itoa(beam.pos[0])+"-"+strconv.Itoa(beam.pos[1])) {
			vis = append(vis, beam.dir+"-"+strconv.Itoa(beam.pos[0])+"-"+strconv.Itoa(beam.pos[1]))

			if !slices.Contains(energized, strconv.Itoa(beam.pos[0])+"-"+strconv.Itoa(beam.pos[1])) {
				energized = append(energized, strconv.Itoa(beam.pos[0])+"-"+strconv.Itoa(beam.pos[1]))
			}

			if field[beam.pos[0]][beam.pos[1]] == '/' {
				if beam.dir == "right" {
					beam.dir = "up"
				} else if beam.dir == "left" {
					beam.dir = "down"
				} else if beam.dir == "up" {
					beam.dir = "right"
				} else if beam.dir == "down" {
					beam.dir = "left"
				}
			} else if field[beam.pos[0]][beam.pos[1]] == '\\' {
				if beam.dir == "right" {
					beam.dir = "down"
				} else if beam.dir == "left" {
					beam.dir = "up"
				} else if beam.dir == "up" {
					beam.dir = "left"
				} else if beam.dir == "down" {
					beam.dir = "right"
				}
			} else if field[beam.pos[0]][beam.pos[1]] == '|' && slices.Contains([]string{"left", "right"}, beam.dir) {
				beams = append(beams, tbeam{dir: "up", pos: []int{beam.pos[0], beam.pos[1]}})
				beams = append(beams, tbeam{dir: "down", pos: []int{beam.pos[0], beam.pos[1]}})
				break
			} else if field[beam.pos[0]][beam.pos[1]] == '-' && slices.Contains([]string{"up", "down"}, beam.dir) {
				beams = append(beams, tbeam{dir: "left", pos: []int{beam.pos[0], beam.pos[1]}})
				beams = append(beams, tbeam{dir: "right", pos: []int{beam.pos[0], beam.pos[1]}})
				break
			}

			switch beam.dir {
			case "right":
				if beam.pos[1] != len(field[beam.pos[0]])-1 {
					beam.pos[1]++
				}
			case "left":
				if beam.pos[1] != 0 {
					beam.pos[1]--
				}
			case "up":
				if beam.pos[0] != 0 {
					beam.pos[0]--
				}
			case "down":
				if beam.pos[0] != len(field)-1 {
					beam.pos[0]++
				}
			}
		}
	}

	return len(energized)
}

func part2(s *bufio.Scanner) (total int) {
	field := make([]string, 0)
	for s.Scan() {
		field = append(field, s.Text())
	}
	i := len(field)
	j := len(field[0])

	maxi := 0
	for x := 0; x < i; x++ {
		fmt.Println(x)
		t := part1(field, "right", []int{x, 0})
		if t > maxi {
			maxi = t
		}
		t = part1(field, "left", []int{x, j - 1})
		if t > maxi {
			maxi = t
		}
	}
	for y := 0; y < j; y++ {
		fmt.Println(y)
		t := part1(field, "down", []int{0, y})
		if t > maxi {
			maxi = t
		}
		t = part1(field, "up", []int{i - 1, y})
		if t > maxi {
			maxi = t
		}
	}

	return maxi
}

func main() {
	f, err := os.Open("day16/input.txt")

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
	field := make([]string, 0)
	for s.Scan() {
		field = append(field, s.Text())
	}
	fmt.Println(part1(field, "right", []int{0, 0}))

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
