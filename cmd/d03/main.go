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
	in := strings.Join(data, "")
	fmt.Printf("priority item sum: %d\n", ruck(in))
}

func ruck(input string) int {
	const (
		sackSize   = 12
		pocketSize = sackSize / 2
	)
	var (
		start  = 0
		end    = 6
		length = len(input)
	)
	// slice input into pocket sized chuncks
	var pockets []string
	for end <= length {
		pockets = append(pockets, input[start:end])
		start = end
		end += pocketSize
	}
	// iterate over pocket sets
	var pocketCount = len(pockets)
	var rItems []rune
	for i := 0; i < pocketCount; i += 2 {
		if i+1 > pocketCount {
			break
		}
		left := pockets[i]
		right := pockets[i+1]
		if !strings.ContainsAny(left, right) {
			continue
		}
		// iterate over a pocket
		for j := range left {
			if strings.ContainsRune(right, rune(left[j])) {
				rItems = append(rItems, rune(left[j]))
			}
		}
	}
	//fmt.Printf("%+v\n", pockets)
	//fmt.Printf("%+v\n", rItems)
	fmt.Printf("a=%v, A=%v\n", 'a', 'A')
	var result int
	for i := range rItems {
		result += score(rItems[i])
	}
	return result
}

func score(in rune) int {
	const (
		lowCaseOffset   = 96
		upperCaseOffset = 64
	)
	if in > lowCaseOffset && in < lowCaseOffset+26 {
		return int(in) - lowCaseOffset
	} else if in > upperCaseOffset && in < upperCaseOffset+26 {
		return int(in) - upperCaseOffset + 26
	}
	return 0
}
