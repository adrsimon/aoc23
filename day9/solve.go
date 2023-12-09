package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertToSuites(s *bufio.Scanner) (suites [][]int) {
	suites = make([][]int, 0)
	for s.Scan() {
		strs := strings.Split(s.Text(), " ")
		suite := make([]int, 0)
		for _, str := range strs {
			num, _ := strconv.Atoi(str)
			suite = append(suite, num)
		}
		suites = append(suites, suite)
	}
	return suites
}

func extrapolate(suite []int) (extrapolated int) {
	tab := make([][]int, 0)
	tab = append(tab, suite)

	ok := false
	for !ok {
		toExtrapolate := tab[len(tab)-1]
		subSuite := make([]int, 0)
		for i := 0; i < len(toExtrapolate)-1; i++ {
			subSuite = append(subSuite, toExtrapolate[i+1]-toExtrapolate[i])
		}
		tab = append(tab, subSuite)

		ok = true
		for _, num := range subSuite {
			if num != 0 {
				ok = false
				break
			}
		}
	}

	for i := len(tab) - 2; i >= 0; i-- {
		tab[i] = append(tab[i], tab[i+1][len(tab[i+1])-1]+tab[i][len(tab[i])-1])
	}

	extrapolated = tab[0][len(tab[0])-1]

	return extrapolated
}

func part1(s *bufio.Scanner) (total int) {
	suites := convertToSuites(s)

	for _, suite := range suites {
		total += extrapolate(suite)
	}

	return total
}

func part2(s *bufio.Scanner) (total int) {
	suites := convertToSuites(s)

	for _, suite := range suites {
		reverted := make([]int, 0)
		for i := len(suite) - 1; i >= 0; i-- {
			reverted = append(reverted, suite[i])
		}
		total += extrapolate(reverted)
	}

	return total
}

func main() {
	f, err := os.Open("day9/input.txt")

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
