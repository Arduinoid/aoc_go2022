package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const (
	data = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	expectedSum = 157
)

func TestRuck(t *testing.T) {
	in := strings.Split(data, "\n")
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
