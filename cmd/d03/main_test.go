package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const (
	ruckData = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	expectedSum = 157

	ruckBadgeData = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	expectedBadgeSum = 70
)

func TestRuck(t *testing.T) {
	in := strings.Split(ruckData, "\n")
	assert.Equal(t, expectedSum, ruck(in))
}

func TestScore(t *testing.T) {
	A := 'A'
	a := 'a'
	m := 'm'
	Z := 'Z'
	z := 'z'
	M := 'M'
	assert.Equal(t, 27, score(A))
	assert.Equal(t, 1, score(a))
	assert.Equal(t, 13, score(m))
	assert.Equal(t, 52, score(Z))
	assert.Equal(t, 26, score(z))
	assert.Equal(t, 39, score(M))
}

func TestFindBadges(t *testing.T) {
	in := strings.Split(ruckBadgeData, "\n")
	assert.Equal(t, expectedBadgeSum, badges(in))
}
