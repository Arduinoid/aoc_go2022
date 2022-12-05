package main

import (
	"fmt"
	"strings"

	"go_aoc/input"
)

/*
use math on the runes to calculate scores for each of the redundant items in the pack
*/

func main() {
	data := input.GetInputData("day3.txt")
	fmt.Printf("priority item sum: %d\n", ruck(data))
	var start = 'A'
	for i := 0; i < +26; i++ {
		fmt.Printf("%s=%d\n", string(start+rune(i)), start+rune(i)-'A'+27)
	}
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
