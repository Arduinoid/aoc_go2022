package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRPSGame(t *testing.T) {
	// get file data
	const input = `A Y
B X
C Z`
	const expect = 15
	lines := strings.Split(input, "\n")
	// run game function
	_, p := runGame(lines)
	// assert results
	assert.Equal(t, expect, p)
}
