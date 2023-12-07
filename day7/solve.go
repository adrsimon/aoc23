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

var cards1 = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var cards2 = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

const (
	NOTHING = iota
	PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

func evaluate1(hand []string) (score int) {
	score = NOTHING
	handCards := make(map[string]int)
	for _, card := range hand {
		handCards[card]++
	}

	groups := make([]int, 0)

	for _, card := range cards1 {
		if handCards[card] == 2 {
			groups = append(groups, PAIR)
		}
		if handCards[card] == 3 {
			groups = append(groups, THREE_OF_A_KIND)
		}
		if handCards[card] == 4 {
			groups = append(groups, FOUR_OF_A_KIND)
		}
		if handCards[card] == 5 {
			groups = append(groups, FIVE_OF_A_KIND)
		}
	}

	if len(groups) == 2 {
		if groups[0] == PAIR && groups[1] == THREE_OF_A_KIND {
			score = FULL_HOUSE
		}
		if groups[0] == THREE_OF_A_KIND && groups[1] == PAIR {
			score = FULL_HOUSE
		}
		if groups[0] == PAIR && groups[1] == PAIR {
			score = TWO_PAIR
		}
	} else if len(groups) == 1 {
		score = groups[0]
	}

	return score
}

func evaluate2(hand []string) (score int) {
	score = NOTHING
	handCards := make(map[string]int)
	for _, card := range hand {
		handCards[card]++
	}

	for _, hand := range distributeJacks(handCards) {
		test := make([]string, 0)
		for card, count := range hand {
			for i := 0; i < count; i++ {
				test = append(test, card)
			}
		}
		if evaluate1(test) > score {
			score = evaluate1(test)
		}
	}

	return
}

func distributeJacks(handCards map[string]int) (hands []map[string]int) {
	jacks := handCards["J"]
	hands = make([]map[string]int, 0)

	hands = append(hands, handCards)
	for cards := range handCards {
		hand := make(map[string]int)
		for k, v := range handCards {
			hand[k] = v
		}

		if cards == "J" {
			continue
		}
		hand[cards] += jacks
		hand["J"] = 0
		hands = append(hands, hand)
	}

	return
}

func part1(s *bufio.Scanner) (total int) {
	hands := make([][]string, 0)
	bets := make([]int, 0)
	for s.Scan() {
		line := s.Text()
		hand, strbet := strings.Fields(line)[0], strings.Fields(line)[1]
		hands = append(hands, strings.Split(hand, ""))
		bet, _ := strconv.Atoi(strbet)
		bets = append(bets, bet)
	}

	/** Sorting **/
	for i := 0; i < len(hands)-1; i++ {
		for j := 0; j < len(hands)-i-1; j++ {
			if evaluate1(hands[j]) > evaluate1(hands[j+1]) {
				hands[j], hands[j+1] = hands[j+1], hands[j]
				bets[j], bets[j+1] = bets[j+1], bets[j]
			} else if evaluate1(hands[j]) == evaluate1(hands[j+1]) {
				decided := false
				for !decided {
					for k := 0; k < len(cards1); k++ {
						if slices.Index(cards1, hands[j][k]) > slices.Index(cards1, hands[j+1][k]) {
							hands[j], hands[j+1] = hands[j+1], hands[j]
							bets[j], bets[j+1] = bets[j+1], bets[j]
							decided = true
							break
						} else if slices.Index(cards1, hands[j][k]) < slices.Index(cards1, hands[j+1][k]) {
							decided = true
							break
						}
					}
				}
			}
		}
	}

	for i, bet := range bets {
		total += bet * (i + 1)
	}

	return total
}

func part2(s *bufio.Scanner) (total int) {
	hands := make([][]string, 0)
	bets := make([]int, 0)
	for s.Scan() {
		line := s.Text()
		hand, strbet := strings.Fields(line)[0], strings.Fields(line)[1]
		hands = append(hands, strings.Split(hand, ""))
		bet, _ := strconv.Atoi(strbet)
		bets = append(bets, bet)
	}

	evaluate2(hands[0])

	/** Sorting **/
	for i := 0; i < len(hands)-1; i++ {
		for j := 0; j < len(hands)-i-1; j++ {
			if evaluate2(hands[j]) > evaluate2(hands[j+1]) {
				hands[j], hands[j+1] = hands[j+1], hands[j]
				bets[j], bets[j+1] = bets[j+1], bets[j]
			} else if evaluate2(hands[j]) == evaluate2(hands[j+1]) {
				decided := false
				for !decided {
					for k := 0; k < len(cards2); k++ {
						if slices.Index(cards2, hands[j][k]) > slices.Index(cards2, hands[j+1][k]) {
							hands[j], hands[j+1] = hands[j+1], hands[j]
							bets[j], bets[j+1] = bets[j+1], bets[j]
							decided = true
							break
						} else if slices.Index(cards2, hands[j][k]) < slices.Index(cards2, hands[j+1][k]) {
							decided = true
							break
						}
					}
				}
			}
		}
	}

	for i, bet := range bets {
		total += bet * (i + 1)
	}

	return total
}

func main() {
	f, err := os.Open("day7/input.txt")

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
