package main

import (
	"fmt"
	"go_aoc/input"
	"strconv"
	"strings"
)

func main() {
	in := input.GetInputData("day5.txt")
	stacks, moves := parse(in)
	stacks2 := stacks.copy()
	for i := range moves {
		stacks.exec9k(moves[i])
		stacks2.exec9k1(moves[i])
	}
	fmt.Printf("9000 top boxes: %s\n", stacks.top())
	fmt.Printf("9001 top boxes: %s\n", stacks2.top())
}

func parse(in []string) (stack, [][]string) {
	var (
		moveRaw, stackRaw []string
		delimiter         bool
	)
	for i := range in {
		if !delimiter && strings.TrimSpace(in[i]) == "" {
			delimiter = true
			continue
		}
		if !delimiter {
			stackRaw = append(stackRaw, in[i])
			continue
		}
		moveRaw = append(moveRaw, in[i])
	}
	return parseStack(stackRaw), parseMoves(moveRaw)
}

func parseStack(in []string) stack {
	var indices []int
	// get the last line to find the column count
	for i := range in[len(in)-1] {
		if in[len(in)-1][i] == ' ' {
			continue
		}
		indices = append(indices, i)
	}
	// loop over the input in reverse and grab the crates
	st := make(stack)
	for i := len(in) - 2; i >= 0; i-- {
		// use the indices to find the crates for each column
		for j, n := range indices {
			if len(in[i]) < n || in[i][n] == ' ' {
				continue
			}
			key := strconv.Itoa(j + 1)
			st[key] = append(st[key], string(in[i][n]))
		}
	}
	for i := range indices {
		st["keys"] = append(st["keys"], strconv.Itoa(i+1))
	}
	return st
}

func parseMoves(in []string) [][]string {
	var moves = make([][]string, len(in))
	for i := range in {
		moves[i] = strings.Split(in[i], " ")
	}
	return moves
}

type stack map[string][]string

func (s stack) exec9k(m []string) {
	// move 7 from 3 to 9
	// [0] [1][2][3][4][5]
	count, err := strconv.Atoi(m[1])
	if err != nil {
		panic("could not parse instruction count")
	}
	srcLen := len(s[m[3]])
	// if the amount of crates to remove is greater than crates at source then just take all of what is there
	if count > srcLen {
		s[m[5]] = append(s[m[5]], s[m[3]]...)
		s[m[3]] = []string{} // clear our source crates
		return
	}
	// otherwise start taking the desired number of crates
	var step = 1
	for i := count; i > 0; i-- {
		s[m[5]] = append(s[m[5]], s[m[3]][srcLen-step])
		step++
	}
	// set source to no longer contain the moved crates
	s[m[3]] = s[m[3]][:srcLen-count]
}

func (s stack) exec9k1(m []string) {
	// move 7 from 3 to 9
	// [0] [1][2][3][4][5]
	count, err := strconv.Atoi(m[1])
	if err != nil {
		panic("could not parse instruction count")
	}
	srcLen := len(s[m[3]])
	// if the amount of crates to remove is greater than crates at source then just take all of what is there
	if count > srcLen {
		s[m[5]] = append(s[m[5]], s[m[3]]...)
		s[m[3]] = []string{} // clear our source crates
		return
	}
	// otherwise start taking the desired number of crates
	s[m[5]] = append(s[m[5]], s[m[3]][srcLen-count:]...)
	// set source to no longer contain the moved crates
	s[m[3]] = s[m[3]][:srcLen-count]
}

func (s stack) top() string {
	var output string
	for _, key := range s["keys"] {
		output += s[key][len(s[key])-1]
	}
	return output
}

func (s stack) copy() stack {
	var newStack = stack{}
	for k, v := range s {
		newStack[k] = v
	}
	return newStack
}
