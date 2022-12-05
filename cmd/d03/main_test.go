package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	data        = `vJrwpWtwJgWrhcsFMMfFFhFpjqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSLPmmdzqPrVvPwwTWBwgwMqvLMZHhHMvwLHjbvcjnnSBnvTQFnttgJtRGJQctTZtZTCrZsJsPPZsGzwwsLwLmpwMDw`
	expectedSum = 157
)

func TestRuck(t *testing.T) {
	assert.Equal(t, expectedSum, ruck(data))
}

func TestScore(t *testing.T) {
	A := 'A'
	a := 'a'
	Z := 'Z'
	z := 'z'
	M := 'M'
	assert.Equal(t, 27, score(A))
	assert.Equal(t, 1, score(a))
	assert.Equal(t, 27+25, score(Z))
	assert.Equal(t, 26, score(z))
	assert.Equal(t, 26, score(M))
}
