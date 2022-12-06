package main

import (
	"fmt"
	"strings"

	"go_aoc/input"
)

func main() {
	data := input.GetInputData("day3.txt")
	fmt.Printf("priority item sum: %d\n", ruck(data))
	fmt.Printf("badge sum: %d\n", badges(data))
}

func ruck(input []string) int {
	var result int
	for i := range input {
		left := input[i][0 : len(input[i])/2]
		right := input[i][len(input[i])/2:]
		if !strings.ContainsAny(left, right) {
			continue
		}
		var tempPocket string
		for j := range left {
			if !strings.ContainsRune(tempPocket, rune(left[j])) && strings.ContainsRune(right, rune(left[j])) {
				tempPocket += string(left[j])
				out := score(rune(left[j]))
				result += out
			}
		}
	}
	fmt.Printf("a=%v, A=%v\n", 'a', 'A')
	return result
}

func badges(input []string) int {
	/*
		[1, 2, 3] [2] 0/3=0
		[2, 4, 5]     1/3=0.333
		[2, 5, 6]     2/3=0.666
		[1, 2, 3] [2] 3/3=1
		[2, 4, 5]
		[2, 5, 6]
	*/
	var iset string
	var result int
	for i := 0; i < len(input); i += 3 {
		// initialize and collect into intermediate array
		iset = ""
		firstElf := input[i]
		secondElf := input[i+1]
		thirdElf := input[i+2]
		for j := range firstElf {
			if !strings.ContainsRune(iset, rune(firstElf[j])) && strings.ContainsRune(secondElf, rune(firstElf[j])) {
				iset += string(firstElf[j])
			}
		}
		for k := range iset {
			if strings.ContainsRune(thirdElf, rune(iset[k])) {
				result += score(rune(iset[k]))
			}
		}
	}
	return result
}

func score(in rune) int {
	const (
		lowA     = 'a'
		upperA   = 'A'
		alphaLen = 26
	)
	if in >= lowA && in <= lowA+alphaLen {
		return (int(in) - lowA) + 1
	} else if in >= upperA && in <= upperA+alphaLen {
		return (int(in) - upperA) + 27
	}
	return 0
}
