package die

import "game_of_life/game/counter"

// * Muere: Una celula viva puede morir por uno de los 2 casos:
// overpopulation or isolation
func IsDie(board [][]bool, i, j int) bool {
	// Si ya esta muerta no puede morir!
	if !board[i][j] {
		return false
	}

	// Sobrepoblación
	if overpopulation(board, i, j) {
		return true
	}
	// Aislamiento
	if isolation(board, i, j) {
		return true
	}

	return false
}

// Sobrepoblación: Si tiene más de 3 vecinos vivos alrededor.
func overpopulation(board [][]bool, i, j int) bool {
	return counter.CountLiveNeighbors(board, i, j) >= 3
}

//   - Aislamiento: Si tiene sólo un vecino vivo alrededor o ninguno.
func isolation(board [][]bool, i, j int) bool {
	return counter.CountLiveNeighbors(board, i, j) <= 1 
}

