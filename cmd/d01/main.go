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
	counts := orderedCounts(string(bs))
	fmt.Printf("highest calorie inventory: %d\n", counts[len(counts)-1])
	fmt.Printf("top three summed: %d\n", sum(counts[len(counts)-3:]...)) // 69208 69434 69795]
}

func orderedCounts(input string) []int {
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
	return nums
}

func sum(nums ...int) int {
	var res int
	for _, n := range nums {
		res += n
	}
	return res
}
