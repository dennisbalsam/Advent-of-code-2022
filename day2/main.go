package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	typeScores = map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}
	types = map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}
	typeWinner = map[string]string{
		"paper":    "rock",
		"scissors": "paper",
		"rock":     "scissors",
	}
	typeLoser = map[string]string{
		"rock":     "paper",
		"paper":    "scissors",
		"scissors": "rock",
	}
	lines = []string{}
)

func main() {
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	part1()
	part2()
}

func part1() {
	score := 0
	for _, line := range lines {
		s := strings.Split(line, " ")
		score += typeScores[types[s[1]]]
		if types[s[0]] == types[s[1]] {
			score += 3
		} else {
			switch fmt.Sprintf("%s,%s", types[s[0]], types[s[1]]) {
			case "rock,paper", "scissors,rock", "paper,scissors":
				score += 6
			}
		}
	}
	fmt.Printf("Total Score Part 1: %v\n", score)
}

func part2() {
	score := 0
	roundOutcomes := map[string]int{"X": 0, "Y": 3, "Z": 6}
	for _, line := range lines {
		s := strings.Split(line, " ")
		score += roundOutcomes[s[1]]
		switch s[1] {
		case "X":
			score += typeScores[typeWinner[types[s[0]]]]
		case "Y":
			score += typeScores[types[s[0]]]
		case "Z":
			score += typeScores[typeLoser[types[s[0]]]]
		}
	}
	fmt.Printf("Total Score Part 2: %v\n", score)
}
