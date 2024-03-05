package die

import "game_of_life/game/utils"

// * Muere: Una celula viva puede morir por uno de los 2 casos:
//   - Aislamiento: Si tiene sólo un vecino vivo alrededor o ninguno.
func Die(board [][]bool, i, j int) bool {
	// Si ya esta muerta no puede morir!
	if isDead(board, i, j) {
		return false
	}

	// Sobrepoblación
	if overpopulation(board, i, j) {
		return true
	}
	// Aislamiento

	return false
}

func isDead(board [][]bool, i, j int) bool {
	return board[i][j]
}

//   - Sobrepoblación: Si tiene más de 3 vecinos vivos alrededor.
func overpopulation(board [][]bool, i, j int) bool {
	
	if evaluateRowUp(board, i, j) >= 3 {
		return true
	}

	if evaluateRowDown(board, i, j) >= 3 {
		return true
	}
	return false
}

func evaluateRowUp(board[][]bool, i, j int) int {
	cont := 0
	up := i - 1
	if up >= 0 {
		// Chequeo fila superior
		right := j - 1
		left := j + 1
		// Chequeo si la celula de arriba esta viva
		if !isDead(board, up, j) {
			cont++
		}
		// Chequeo si la celula de arriba a la derecha esta viva
		if right >=0 && !isDead(board, up, right) {
			cont++
		}
		// Chequeo si la celula de arriba a la izquierda esta viva
		if left < utils.Width && !isDead(board, up, left) {
			cont++
		} 
	}

	return cont
}

func evaluateRowDown(board[][]bool, i, j int) int {
	cont := 0
	down := i + 1
	if down < utils.ScreenHeight && down > 0 {
		// Chequeo fila inferior
		right := j - 1
		left := j + 1
		// Chequeo si la celula de abajo esta viva
		if !isDead(board, down, j) {
			cont++
		}
		// Chequeo si la celula de abajo a la derecha esta viva
		if !isDead(board, down, right) {
			cont++
		}
		// Chequeo si la celula de abajo a la izquierda esta viva
		if !isDead(board, down, left) {
			cont++
		}
	}

	return cont
}

