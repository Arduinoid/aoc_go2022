package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dir := os.Getenv("INPUTDATA")
	bs, err := os.ReadFile(dir + "/day1-1.txt")
	if err != nil {
		panic("couldn't read in file: " + dir + "/day1-1.txt")
	}
	highest := highestCount(string(bs))
	fmt.Printf("highest calorie inventory: %d\n", highest)
}

func highestCount(input string) int {
	lines := strings.Split(input, "\n")
	var nums = make([]int, len(lines))
	var elfIndx int

	for _, n := range lines {
		if n == "" {
			elfIndx++
			continue
		}
		if num, err := strconv.Atoi(n); err == nil {
			nums[elfIndx] = nums[elfIndx] + num
		}
	}
	sort.Ints(nums)
	return nums[len(nums)-1]
}
