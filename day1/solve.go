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
	for s.Scan() {
		str := s.Text()
		chars := strings.Split(str, "")
		nums := make([]int, 0)
		for _, char := range chars {
			if n, err := strconv.Atoi(char); err != nil {
				continue
			} else {
				nums = append(nums, n)
			}
		}
		if len(nums) == 0 {
			continue
		}
		part, err := strconv.Atoi(strconv.Itoa(nums[0]) + strconv.Itoa(nums[len(nums)-1]))
		if err != nil {
			log.Fatal(err)
		}
		total += part
	}
	return total
}

func part2(s *bufio.Scanner) (total int) {
	assoc := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for s.Scan() {
		str := s.Text()
		chars := strings.Split(str, "")
		nums := make([]int, 0)

		for index, char := range chars {
			if n, err := strconv.Atoi(char); err == nil {
				nums = append(nums, n)
				continue
			}

			for key, value := range assoc {
				if len(key)+index > len(chars) {
					continue
				} else if strings.Join(chars[index:len(key)+index], "") == key {
					nums = append(nums, value)
				}
			}
		}
		part, err := strconv.Atoi(strconv.Itoa(nums[0]) + strconv.Itoa(nums[len(nums)-1]))
		if err != nil {
			log.Fatal(err)
		}
		total += part
	}
	return total
}

func main() {
	f, err := os.Open("day1/input.txt")

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
