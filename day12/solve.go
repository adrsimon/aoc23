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

func count(springs string, arr []int) (total int) {
	var repl []string
	cnt := strings.Count(springs, "?")
	sum := 0
	for _, v := range arr {
		sum += v
	}

	for i := 0; i < (1 << cnt); i++ {
		replaced := springs
		for j := 0; j < len(springs); j++ {
			if (i & (1 << j)) != 0 {
				replaced = strings.Replace(replaced, "?", "#", 1)
			} else {
				replaced = strings.Replace(replaced, "?", ".", 1)
			}
		}
		repl = append(repl, replaced)
	}

	for _, r := range repl {
		if strings.Count(r, "#") != sum {
			continue
		}
		var groups []int
		ongoing := 0
		for _, v := range r {
			if v == '#' {
				if ongoing == 0 {
					ongoing++
					groups = append(groups, ongoing)
				} else {
					ongoing++
					groups[len(groups)-1] = ongoing
				}
			} else {
				ongoing = 0
			}
		}

		if slices.Compare(groups, arr) == 0 {
			total += 1
		}
	}
	return total
}

func part1(s *bufio.Scanner) (total int) {
	for s.Scan() {
		line := strings.Fields(s.Text())
		springs := line[0]
		arrStr := strings.Split(line[1], ",")
		arr := make([]int, len(arrStr))
		for i, v := range arrStr {
			arr[i], _ = strconv.Atoi(v)
		}

		total += count(springs, arr)
	}
	return total
}

func count2(i, j int, springs string, arr []int, cache [][]int) int {
	if i >= len(springs) {
		if j < len(arr) {
			return 0
		}
		return 1
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	part := 0

	if springs[i] == '.' {
		part = count2(i+1, j, springs, arr, cache)
	} else {
		if springs[i] == '?' {
			part += count2(i+1, j, springs, arr, cache)
		}
		if j < len(arr) {
			cnt := 0
			for k := i; k < len(springs); k++ {
				if cnt > arr[j] || springs[k] == '.' || cnt == arr[j] && springs[k] == '?' {
					break
				}
				cnt += 1
			}

			if cnt == arr[j] {
				if i+cnt < len(springs) && springs[i+cnt] != '#' {
					part += count2(i+cnt+1, j+1, springs, arr, cache)
				} else {
					part += count2(i+cnt, j+1, springs, arr, cache)
				}
			}
		}
	}

	cache[i][j] = part
	return part
}

func part2(s *bufio.Scanner) (total int) {
	for s.Scan() {
		line := strings.Fields(s.Text())
		springs := line[0]
		arrStr := strings.Split(line[1], ",")
		arrT := make([]int, len(arrStr))
		arr := arrT
		for i, v := range arrStr {
			arrT[i], _ = strconv.Atoi(v)
		}

		for i := 0; i < 4; i++ {
			springs += "?" + line[0]
			arr = append(arr, arrT...)
		}

		var cache [][]int
		for i := 0; i < len(springs); i++ {
			cache = append(cache, make([]int, len(arr)+1))
			for j := 0; j < len(arr)+1; j++ {
				cache[i][j] = -1
			}
		}

		total += count2(0, 0, springs, arr, cache)
	}
	return total
}

func main() {
	f, err := os.Open("day12/input.txt")

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
