package main

import (
	"fmt"
	"go_aoc/input"
	"strings"
)

/*
A Y - Rock, Paper
B X - Paper, Rock
C Z - Scissors, Scissors
This strategy guide predicts and recommends the following:

In the first round, your opponent will choose Rock (A), and you should choose Paper (Y). This ends in a win for you with a score of 8 (2 because you chose Paper + 6 because you won).
In the second round, your opponent will choose Paper (B), and you should choose Rock (X). This ends in a loss for you with a score of 1 (1 + 0).
The third round is a draw with both players choosing Scissors, giving you a score of 3 + 3 = 6.

rock     = 1 : AX
paper    = 2 : BY
scissors = 3 : CZ
*/

func main() {
	data := input.GetInputData("day2.txt")
	op, ply := runGame(data)
	fmt.Printf("opponent: %d player: %d\n", op, ply)
	op2, ply2 := runGameP2(data)
	fmt.Printf("part two -- opponent: %d player: %d\n", op2, ply2)
}

func runGame(input []string) (int, int) {
	var opScore, pScore int
	for _, line := range input {
		moves := strings.ReplaceAll(strings.TrimSpace(line), " ", "")
		o, p := roundScore(moves)
		opScore += o
		pScore += p
	}
	return opScore, pScore
}

func runGameP2(input []string) (int, int) {
	var opScore, pScore int
	for _, line := range input {
		line = strings.TrimSpace(line)
		moves := strings.Split(line, " ")
		playerChoice := choose(moves)
		o, p := roundScore(moves[0] + playerChoice)
		opScore += o
		pScore += p
	}
	return opScore, pScore
}

func choiceScore(choice string) int {
	switch choice {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	default:
		return 0
	}
}

func roundScore(moves string) (int, int) {
	var opScore, pScore int
	outcome := map[string]int{
		"AX": 0, "AY": 1, "AZ": -1,
		"BX": -1, "BY": 0, "BZ": 1,
		"CX": 1, "CY": -1, "CZ": 0,
	}
	switch outcome[moves] {
	case -1:
		opScore += 6
	case 0:
		opScore += 3
		pScore += 3
	case 1:
		pScore += 6
	}
	opScore += choiceScore(string(moves[0]))
	pScore += choiceScore(string(moves[1]))
	return opScore, pScore
}

func win(op string) string {
	switch op {
	case "A":
		return "Y"
	case "B":
		return "Z"
	case "C":
		return "X"
	default:
		return ""
	}
}

func lose(op string) string {
	switch op {
	case "A":
		return "Z"
	case "B":
		return "X"
	case "C":
		return "Y"
	default:
		return ""
	}
}

func draw(op string) string {
	switch op {
	case "A":
		return "X"
	case "B":
		return "Y"
	case "C":
		return "Z"
	default:
		return ""
	}
}

func choose(result []string) string {
	const (
		strat    = 1
		opponent = 0
	)
	switch result[strat] {
	case "X":
		return lose(result[opponent])
	case "Y":
		return draw(result[opponent])
	case "Z":
		return win(result[opponent])
	default:
		return ""
	}
}
