package main

import (
	"fmt"
	"strings"
)

/*
use math on the runes to calculate scores for each of the redundant items in the pack
*/

func main() {
	input := "vJrwpWtwJgWrhcsFMMfFFhFpjqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSLPmmdzqPrVvPwwTWBwgwMqvLMZHhHMvwLHjbvcjnnSBnvTQFnttgJtRGJQctTZtZTCrZsJsPPZsGzwwsLwLmpwMDw"
	fmt.Printf("priority item sum: %d", ruck(input))
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
	var pockets []string
	for end <= length {
		pockets = append(pockets, input[start:end])
		start = end
		end += pocketSize
	}
	var rItems []rune
	for i := 0; i < len(pockets); i += 2 {
		if i+1 > len(pockets) {
			break
		}
		if !strings.ContainsAny(pockets[i], pockets[i+1]) {
			continue
		}
		for j := range pockets[i] {
			if strings.ContainsRune(pockets[i+1], rune(pockets[i][j])) {
				rItems = append(rItems, string(pockets[i][j]))
			}
		}
	}
	fmt.Printf("%+v\n", rItems)
	fmt.Printf("%v, %v\n", 'a'-96, 'A'-38)
	return 0
}
