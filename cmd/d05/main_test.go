package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

const data = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestParseStack(t *testing.T) {
	in := strings.Split(data, "\n")
	expect := stack{
		"1":    []string{"Z", "N"},
		"2":    []string{"M", "C", "D"},
		"3":    []string{"P"},
		"keys": []string{"1", "2", "3"},
	}
	st, _ := parse(in)
	fmt.Printf("%+v\n", st)
	assert.Equal(t, expect, st)
}

func TestStackExec(t *testing.T) {
	m := []string{"move", "2", "from", "2", "to", "3"}
	in := strings.Split(data, "\n")
	st, _ := parse(in)
	fmt.Printf("before move: %+v\n", st)
	st.exec(m)
	fmt.Printf("after move: %+v\n", st)
}

func TestParseMoves(t *testing.T) {
	in := strings.Split(data, "\n")
	_, mvs := parse(in)
	require.Len(t, mvs, 4)
	assert.Equal(t, mvs[0], []string{"move", "1", "from", "2", "to", "1"})
	assert.Equal(t, mvs[1], []string{"move", "3", "from", "1", "to", "3"})
	assert.Equal(t, mvs[2], []string{"move", "2", "from", "2", "to", "1"})
	assert.Equal(t, mvs[3], []string{"move", "1", "from", "1", "to", "2"})
}

func TestStackTop(t *testing.T) {
	in := strings.Split(data, "\n")
	st, _ := parse(in)
	assert.Equal(t, "NDP", st.top())
}

func TestStackKeys(t *testing.T) {
	in := strings.Split(data, "\n")
	st, _ := parse(in)
	fmt.Printf("%+v\n", st)
	assert.Equal(t, []string{"1", "2", "3"}, st["keys"])
}
