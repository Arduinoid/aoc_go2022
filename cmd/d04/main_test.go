package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const (
	dataSet = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	expectedSets = 2
)

/*
indexes where one contains the other
3,4
*/

func TestContains(t *testing.T) {
	in := strings.Split(dataSet, "\n")
	assert.Equal(t, expectedSets, contains(in))
}

func TestInRange(t *testing.T) {
	assert.True(t, inRange(2, 2, 1, 3))
	assert.False(t, inRange(1, 2, 3, 4))
	assert.True(t, inRange(1, 10, 3, 4))
}

func TestParsePairs(t *testing.T) {
	in := "1-2,3-4"
	expect := []int{1, 2, 3, 4}
	assert.Equal(t, expect, parsePairs(in))
}

func TestIntersets(t *testing.T) {
	assert.True(t, intersect(1, 2, 2, 3))
	assert.False(t, intersect(1, 2, 3, 4))
	assert.False(t, intersect(1, 4, 3, 3))
	assert.True(t, intersect(4, 10, 3, 11))
}
