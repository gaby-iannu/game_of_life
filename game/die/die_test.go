package die

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverPopulation(t *testing.T) {
	cases := []struct{
		name string
		board[][] bool
		overpopulation bool
		row int
		col int
	}{
		{
			name: "Not overpopulation",
			board:[][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, false},
			},
			overpopulation: false,
			row: 1,
			col: 1,
		},
		{
			name: "Not overpopulation up right vertix",
			board:[][]bool{
				{true, false, false},
				{false, false, false},
				{false, false, false},
			},
			overpopulation: false,
			row: 0,
			col: 0,
		},
		{
			name: "Not overpopulation down left vertix",
			board:[][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, true},
			},
			overpopulation: false,
			row: 2,
			col: 2,
		},
		{
			name: "Not overpopulation down right vertix",
			board:[][]bool{
				{false, false, false},
				{false, false, false},
				{true, false, false},
			},
			overpopulation: false,
			row: 2,
			col: 0,
		},
		{
			name: "Not overpopulation up left vertix",
			board:[][]bool{
				{false, false, true},
				{false, false, false},
				{false, false, false},
			},
			overpopulation: false,
			row: 0,
			col: 2,
		},
		{
			name: "Overpopulation row up",
			board:[][]bool{
				{true, true, true},
				{false, true, false},
				{false, false, false},
			},
			overpopulation: true,
			row: 1,
			col: 1,
		},
		{
			name: "Overpopulation row down",
			board:[][]bool{
				{false, false, false},
				{false, true, false},
				{true, true, true},
			},
			overpopulation: true,
			row: 1,
			col: 1,
		},
		{
			name: "Overpopulation col right",
			board:[][]bool{
				{true, false, false},
				{true, true, false},
				{true, false, false},
			},
			overpopulation: true,
			row: 1,
			col: 1,
		},
		{
			name: "Overpopulation col left",
			board:[][]bool{
				{false, false, true},
				{false, true, true},
				{false, false, true},
			},
			overpopulation: true,
			row: 1,
			col: 1,
		},
	}

	for _,c := range cases {
		t.Run(c.name, func(t *testing.T){
			result := overpopulation(c.board, c.row, c.col)
			assert.Equal(t, c.overpopulation, result, "result isn't equal!")
		})	
	}
}

func TestIsolation(t *testing.T) {
	cases := []struct{
		name string
		board [][]bool
		isolation bool
		row int
		col int
	}{
		{
			name: "isolation",
			board:[][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, false},

			},
			isolation: true,
			row: 0,
			col: 0,
		},
		{
			name: "isolation only 1 cell live",
			board:[][]bool{
				{false, false, false},
				{false, true, false},
				{false, false, false},

			},
			isolation: true,
			row: 0,
			col: 0,
		},
		{
			name: "isolation only 1 cell live on vertix down",
			board:[][]bool{
				{false, false, false},
				{false, true, false},
				{false, false, false},

			},
			isolation: true,
			row: 1,
			col: 1,
		},
		{
			name: "isolation only 1 cell live on vertix up",
			board:[][]bool{
				{false, false, false},
				{false, true, false},
				{false, false, false},

			},
			isolation: true,
			row: 0,
			col: 1,
		},
		{
			name: "no isolation",
			board:[][]bool{
				{false, true, false},
				{false, true, false},
				{false, false, false},

			},
			isolation: false,
			row: 0,
			col: 2,
		},
	}

	for _,c := range cases {
		t.Run(c.name, func(t *testing.T){
			result := isolation(c.board, c.row, c.col)
			assert.Equal(t, c.isolation, result, "it is not same")
		})		
	}
}

func TestDie(t *testing.T) {
	cases := []struct{
		name string
		board [][]bool
		expected bool
		row int
		col int
	}{
		{
			name: "cell is not alive",
			board: [][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, false},
			},
			expected: false,
			row: 0,
			col: 0,
		},
		{
			name: "cell die by overpopulation",
			board: [][]bool{
				{false, false, false},
				{false, true, false},
				{true, true, true},
			},
			expected: true,
			row: 1,
			col: 1,
		},
		{
			name: "cell die by isolation",
			board: [][]bool{
				{false, false, false},
				{false, true, false},
				{false, false, false},
			},
			expected: true,
			row: 1,
			col: 1,
		},
	}

	for _,c := range cases {
		t.Run(c.name, func(t *testing.T){
			result := Die(c.board, c.row, c.col)
			assert.Equal(t, c.expected, result, "they aren't same")
		})
	}
}
