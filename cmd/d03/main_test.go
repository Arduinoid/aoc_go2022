package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	input       = `vJrwpWtwJgWrhcsFMMfFFhFpjqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSLPmmdzqPrVvPwwTWBwgwMqvLMZHhHMvwLHjbvcjnnSBnvTQFnttgJtRGJQctTZtZTCrZsJsPPZsGzwwsLwLmpwMDw`
	expectedSum = 157
)

func TestRuck(t *testing.T) {
	assert.Equal(t, expectedSum, ruck(input))
}
