package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToApply(t *testing.T) {
	cases := []struct{
		name string
		input [][]bool
		expected [][]bool
		row int
		col int
	}{
		{
			name: "To apply is alive",
			input: [][]bool{
				{false, false, false},
				{true, true, false},
				{true, false, false},
			},
			expected: [][]bool {
				{false, false, false},
				{true, true, false},
				{true, false, false},
			},
			row: 2,
			col: 0,

		},
		{
			name: "To apply is born",
			input: [][]bool{
				{false, false, false},
				{true, true, false},
				{true, false, false},
			},
			expected: [][]bool {
				{false, false, false},
				{true, true, false},
				{true, true, false},
			},
			row: 2,
			col: 1,

		},
		{
			name: "To apply is die by overpopulation",
			input: [][]bool{
				{false, false, false},
				{true, true, true},
				{true, true, false},
			},
			expected: [][]bool {
				{false, false, false},
				{true, true, true},
				{true, false, false},
			},
			row: 2,
			col: 1,

		},
		{
			name: "To apply is die by isolation",
			input: [][]bool{
				{false, false, false},
				{false, false, false},
				{true, true, false},
			},
			expected: [][]bool {
				{false, false, false},
				{false, false, false},
				{true, false, false},
			},
			row: 2,
			col: 1,

		},
	}

	for _,c := range cases {
		t.Run(c.name, func(t *testing.T){
			toApply(&c.input, c.row, c.col)
			for i,rows := range c.input {
				for j,v := range rows {
					assert.Equal(t, c.expected[i][j], v, fmt.Errorf("error row %d col %d", i,j))	
				}
			}
		})
	}

}
