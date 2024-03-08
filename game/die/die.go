package die

// * Muere: Una celula viva puede morir por uno de los 2 casos:
// overpopulation or isolation
func Die(board [][]bool, i, j int) bool {
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
	return countNeighbors(board, i, j) >= 3
}

//   - Aislamiento: Si tiene sólo un vecino vivo alrededor o ninguno.
func isolation(board [][]bool, i, j int) bool {
	return countNeighbors(board, i, j) <= 1 
}

func countNeighbors(board [][]bool, i, j int) int {
	cont := 0
	rows, cols := len(board), len(board[0]) 
	up, right := i - 1, j - 1
	down, left := i + 1, j + 1

	// up right vertix
	if up >= 0 && right >= 0 && board[up][right]{
		cont++
	}
	// up 
	if up >= 0 && board[up][j] {
		cont++
	}
	// up left vertix
	if up >= 0 && left < cols && board[up][left] {
		cont++
	}
	// same left
	if left < cols && up >= 0 && board[up][left] {
		cont++
	}
	// down left vertix
	if down < rows && left < cols && board[down][left]{
		cont++
	}
	// down
	if down < rows && board[down][j] {
		cont++
	}
	// down right vertix
	if down < rows && right >= 0 && board[down][right] {
		cont++
	}
	// same right
	if right >= 0 && board[i][right] {
		cont++
	}

	return cont
}
