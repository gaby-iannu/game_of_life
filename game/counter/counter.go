package counter

// Cuenta vecinos vivos
func CountLiveNeighbors(board [][]bool, i, j int) int {
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
	if left < cols && board[i][left] {
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
