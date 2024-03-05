// Package game provides ...
package game

import (
	"context"
	"game_of_life/game/ui"
	"game_of_life/game/utils"
	"math/rand"
	"time"
)

// Ejecuta el juego de la vida
func Run(c chan [][]bool, ctx context.Context) {
	rand.Seed(time.Now().UnixNano())
	go update(c, ctx)
	ui.Show(c)
}

// Rules
// Cada celula tiene 8 celulas vecinas, que son las celulas que están próximas, incluida la diagonal.
// Las celulas tienen 2 estados vivas o muertas.
// El estado de las celulas evoluciona a lo largo de unidades de tiempo discretas.
// El estado de todas las celulas se tiene en cuenta para calcular el estado de las mismas al turno siguiente.
// Todas las celulas se actulizan simultaneamente en cada turno, siguiendo estas reglas:
// * Nace: Si una celula muerta tiene exactamente 3 celulas vecinas vivas "nace" (siguiente iteración estará viva)

// Estado finales:
// * Extinción: Al cabo de un número finito de generaciones desaparecen todos los miembros de la problación o células vivas.
// * Estabilización: Al cabo de un número finito de generaciones la población queda estabilizada, ya sea de forma
//   rígida o bien de forma oscilante entre dos o más formas
// * Crecimiento constante: La población crece turno tras turno y se mantiene así un número infinito de generaciones.

func update(c chan [][]bool, ctx context.Context) {
	var board [][]bool 
	board = start()
	for {
		select{
		case <- ctx.Done():
			return
		default:
			updateValues(c, &board)
		}
	}
}

func start() [][]bool {
	board := make([][]bool, utils.ScreenWidth)
	for i := 0; i < utils.ScreenWidth; i++ {
		board[i] = make([]bool, utils.ScreenHeight)
		for j := 0; j < utils.ScreenHeight; j++ {
			board[i][j] = trueOrFalse()	
		}
	}

	return board
}

func trueOrFalse() bool {
	if rand.Intn((1-0+1)+0) == 1 {
		return true
	}
	return false
}

func updateValues(c chan [][]bool, board *[][]bool) {
	c <- *(board)
	for i := 0; i < utils.ScreenWidth; i++ {
		for j := 0; j < utils.ScreenHeight; j++ {
			(*board)[i][j] = trueOrFalse()	
		}
	}
	time.Sleep(500 * time.Millisecond)
}

