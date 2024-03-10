package born

import "game_of_life/game/counter"

// * Nace: Si una celula muerta tiene exactamente 3 celulas vecinas vivas "nace" (siguiente iteración estará viva)
func IsBorn(board[][]bool, i, j int) bool {
	// Si esta viva no puede volver a nacer
	if board[i][j] {
		return false
	}

	// Nace si tiene exactamente 3 celulas vecinas vivas
	if counter.CountLiveNeighbors(board, i, j) == 3 {
		return true
	}
	return false
} 
