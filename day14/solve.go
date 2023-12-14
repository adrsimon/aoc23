package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func tiltNorth(p []string) []string {
	platform := make([][]string, 0)

	for _, line := range p {
		platform = append(platform, strings.Split(line, ""))
	}

	for i, row := range platform {
		for j, rock := range row {
			if rock == "O" {
				k := i
				for k > 0 && platform[k-1][j] == "." {
					platform[k][j] = "."
					platform[k-1][j] = "O"
					k--
				}
			}
		}
	}

	for i, row := range platform {
		p[i] = strings.Join(row, "")
	}

	return p
}
func tiltSouth(p []string) []string {
	platform := make([][]string, 0)

	for _, line := range p {
		platform = append(platform, strings.Split(line, ""))
	}

	for i := len(platform) - 1; i >= 0; i-- {
		for j, rock := range platform[i] {
			if rock == "O" {
				k := i
				for k < len(platform)-1 && platform[k+1][j] == "." {
					platform[k][j] = "."
					platform[k+1][j] = "O"
					k++
				}
			}
		}
	}

	for i, row := range platform {
		p[i] = strings.Join(row, "")
	}

	return p
}
func tiltWest(p []string) []string {
	platform := make([][]string, 0)

	for _, line := range p {
		platform = append(platform, strings.Split(line, ""))
	}

	for j := 0; j < len(platform[0]); j++ {
		for i, row := range platform {
			if row[j] == "O" {
				k := j
				for k > 0 && platform[i][k-1] == "." {
					platform[i][k] = "."
					platform[i][k-1] = "O"
					k--
				}
			}
		}
	}

	for i, row := range platform {
		p[i] = strings.Join(row, "")
	}

	return p
}
func tiltEast(p []string) []string {
	platform := make([][]string, 0)

	for _, line := range p {
		platform = append(platform, strings.Split(line, ""))
	}

	for j := len(platform[0]) - 1; j >= 0; j-- {
		for i, row := range platform {
			if row[j] == "O" {
				k := j
				for k < len(platform[0])-1 && platform[i][k+1] == "." {
					platform[i][k] = "."
					platform[i][k+1] = "O"
					k++
				}
			}
		}
	}

	for i, row := range platform {
		p[i] = strings.Join(row, "")
	}

	return p
}

func part1(s *bufio.Scanner) (total int) {
	platform := make([]string, 0)
	for s.Scan() {
		platform = append(platform, s.Text())
	}

	platform = tiltSouth(platform)

	for i, row := range platform {
		for _, rock := range row {
			if rock == 'O' {
				total += len(platform) - i
			}
		}
	}

	return total
}

func part2(s *bufio.Scanner) (total int) {
	platform := make([]string, 0)
	for s.Scan() {
		platform = append(platform, s.Text())
	}

	memo := make(map[string]int)

	cnt := 1000000000
	bp := cnt
	for n := 0; n < cnt; n++ {
		fmt.Println("ouais")
		platform = tiltNorth(platform)
		platform = tiltWest(platform)
		platform = tiltSouth(platform)
		platform = tiltEast(platform)
		if v, ok := memo[strings.Join(platform, "")]; ok {
			bp = n + (cnt-v)%(n-v) - 1
		} else {
			memo[strings.Join(platform, "")] = n
		}
		if n == bp {
			break
		}
	}

	for i, row := range platform {
		for _, rock := range row {
			if rock == 'O' {
				total += len(platform) - i
			}
		}
	}

	return total
}

func main() {
	f, err := os.Open("day14/input.txt")

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
