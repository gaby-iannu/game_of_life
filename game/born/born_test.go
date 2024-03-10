package born

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestIsBorn(t *testing.T) {
	cases := []struct{
		name string
		board[][]bool
		expected bool
		row int
		col int
	}{
		{
			name: "it is alive it can't born again",
			board:[][]bool{
				{true, false, false},
				{false, false, false},
				{false, false, false},
			},
			expected: false,
			row: 0,
			col: 0,
		},
		{
			name: "it is born by count live neighbor",
			board:[][]bool{
				{false, false, false},
				{false, true, true},
				{false, true, false},
			},
			expected: true,
			row: 2,
			col: 2,
		},
		{
			name: "it is not born by it has less 3 neighbor live",
			board:[][]bool{
				{false, false, false},
				{false, true, true},
				{false, false, true},
			},
			expected: false,
			row: 2,
			col: 2,
		},
	}

	for _,c := range cases {
		t.Run(c.name, func(t *testing.T){
			result := IsBorn(c.board, c.row, c.col)
			assert.Equal(t, c.expected, result, "values are not equals")
		})		
	}
}
