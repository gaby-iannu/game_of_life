package alive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestIsAlive(t *testing.T) {
	cases := []struct{
		name string
		board [][]bool
		expected bool
		row int
		col int
	}{
		{
			name: "is not alive",
			board: [][]bool {
				{false, false, false},
				{false, false, false},
				{false, false, false},
			},
			expected: false,
			row: 0,
			col: 0,
		},
		{
			name: "is not alive count 1",
			board: [][]bool {
				{false, true, false},
				{false, false, true},
				{false, false, false},
			},
			expected: false,
			row: 0,
			col: 1,
		},
		{
			name: "is alive count 2",
			board: [][]bool {
				{false, true, true},
				{false, false, true},
				{false, false, false},
			},
			expected: true,
			row: 0,
			col: 1,
		},
		{
			name: "is alive count 3",
			board: [][]bool {
				{true, true, true},
				{false, false, true},
				{false, false, false},
			},
			expected: true,
			row: 0,
			col: 1,
		},
	}

	for _,c:= range cases {
		t.Run(c.name, func(t *testing.T){
			result := IsAlive(c.board, c.row, c.col)
			assert.Equal(t, c.expected, result, "they aren't equals")
		})	
	}
}
