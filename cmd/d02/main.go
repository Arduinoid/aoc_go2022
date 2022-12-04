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

const (
	win  = 1
	lose = -1
	draw = 0
)

func main() {
	data := input.GetInputData("day2.txt")
	lines := strings.Split(data, "\n")
	//shapes := map[string]int{"rock": 1, "paper": 2, "scissors": 3}
	//oponent := map[string]string{"A": "rock", "B": "paper", "C": "scissors"}
	//player := map[string]string{"X": "rock", "Y": "paper", "Z": "scissors"}
	winChoice := map[string]string{"A": "Y", "B": "Z", "C": "X"}
	loseChoice := map[string]string{"A": "Z", "B": "X", "C": "Y"}
	drawChoice := map[string]string{"A": "X", "B": "Y", "C": "Z"}
	//strat := []int{win, lose, draw}
	stratChoice := map[int]map[string]string{win: winChoice, lose: loseChoice, draw: drawChoice}
	outcome := map[string]int{
		"AX": 0, "AY": 1, "AZ": -1,
		"BX": -1, "BY": 0, "BZ": 1,
		"CX": 1, "CY": -1, "CZ": 0,
	}
	choiceScore := map[string][]int{
		"AX": {1, 1}, "AY": {1, 2}, "AZ": {1, 3},
		"BX": {2, 1}, "BY": {2, 2}, "BZ": {2, 3},
		"CX": {3, 1}, "CY": {3, 2}, "CZ": {3, 3},
	}
	var opScore, pScore int
	var st = make(state)
	for i, line := range lines {
		play := strings.Split(line, " ")
		moves := play[0] + play[1]
		if _, ok := st[outcome[moves]]; ok {

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
		opScore += choiceScore[moves][0]
		pScore += choiceScore[moves][1]
	}
	fmt.Printf("opponent: %d player: %d", opScore, pScore)
	/*
		outcome if played as input
		opponent: 13069 player: 10595

		if strategy applied from example
			A Y - Rock, Paper
			B X - Paper, Rock
			C Z - Scissors, Scissors
		opponent: 13978 player: 9428

		if win, lose, draw sequence applied
		opponent: 12025 player: 12500
	*/
}

type state map[int]bool
