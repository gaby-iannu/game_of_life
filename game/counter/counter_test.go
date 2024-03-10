package counter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestCountLiveNeighbors(t *testing.T) {

	cases := []struct{
		name string
		board[][]bool
		expected int
		row int
		col int
	}{
		{
			name: "down",
			board: [][]bool{
				{false, false, false},
				{false, true, false},
				{false, false, false},
			},
			expected: 1,
			row: 0,
			col: 1,
		},
		{
			name: "down left vertix",
			board: [][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, true},
			},
			expected: 1,
			row: 1,
			col: 1,
		},
		{
			name: "down right vertix",
			board: [][]bool{
				{false, false, false},
				{false, false, true},
				{false, false, false},
			},
			expected: 1,
			row: 0,
			col: 1,
		},
		{
			name: "up",
			board: [][]bool{
				{false, false, false},
				{false, false, true},
				{false, false, false},
			},
			expected: 1,
			row: 2,
			col: 2,
		},
		{
			name: "up left vertix",
			board: [][]bool{
				{false, false, false},
				{false, false, true},
				{false, false, false},
			},
			expected: 1,
			row: 2,
			col: 2,
		},
		{
			name: "up right vertix",
			board: [][]bool{
				{true, false, false},
				{false, false, false},
				{false, false, false},
			},
			expected: 1,
			row: 1,
			col: 1,
		},
		{
			name: "same right",
			board: [][]bool{
				{false, false, false},
				{true, false, false},
				{false, false, false},
			},
			expected: 1,
			row: 1,
			col: 1,
		},
		{
			name: "same left",
			board: [][]bool{
				{false, false, false},
				{false, false, true},
				{false, false, false},
			},
			expected: 1,
			row: 1,
			col: 1,
		},
	}

	for _,c := range cases {
		t.Run(c.name, func(t *testing.T){
			cont := CountLiveNeighbors(c.board, c.row, c.col)
			assert.Equal(t, c.expected, cont, "they aren't the same")
		})
	}
}
