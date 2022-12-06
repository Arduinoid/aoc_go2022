package main

import (
	"fmt"
	"go_aoc/input"
	"strconv"
	"strings"
)

func main() {
	data := input.GetInputData("day4.txt")
	fmt.Printf("containing sets: %d\n", contains(data))
	fmt.Printf("overlapping sets: %d\n", overlaps(data))
}

func contains(input []string) int {
	var result int
	for i := range input {
		r := parsePairs(input[i])
		if inRange(r[0], r[1], r[2], r[3]) {
			result++
		}
	}
	return result
}

func overlaps(input []string) int {
	var result int
	for i := range input {
		r := parsePairs(input[i])
		if intersect(r[0], r[1], r[2], r[3]) {
			result++
		}
	}
	return result
}

func intersect(a1, a2, b1, b2 int) bool {
	if a1 >= b1 && a1 <= b2 || a2 <= b2 && a2 >= b1 {
		return true
	}
	return false
}

func inRange(a1, a2, b1, b2 int) bool {
	// b contains a
	if a1 >= b1 && a1 <= b2 && a2 >= b1 && a2 <= b2 {
		return true
	}
	if b1 >= a1 && b1 <= a2 && b2 >= a1 && b2 <= a2 {
		return true
	}
	return false
}

func parsePairs(in string) []int {
	pairs := strings.Split(in, ",")
	ranges := strings.Split(pairs[0], "-")
	ranges = append(ranges, strings.Split(pairs[1], "-")...)
	result := make([]int, len(ranges))
	for i := range ranges {
		n, err := strconv.Atoi(ranges[i])
		if err != nil {
			panic("could not convert number")
		}
		result[i] = n
	}
	return result
}
