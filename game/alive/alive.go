package alive

import "game_of_life/game/counter"

// Vive: una célula se mantiene viva si tiene 2 o 3 vecinos a su alrededor.
func IsAlive(board [][]bool, i, j int) bool {
	count := counter.CountLiveNeighbors(board, i, j)

	return count==2 || count==3
}
