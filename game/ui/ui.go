// Package ui provides ...
package ui

import (
	"game_of_life/game/utils"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Show(c chan [][]bool)  {
	var col color.RGBA

	rl.InitWindow(int32(utils.ScreenWidth), int32(utils.ScreenHeight), "Conway's Game of Life")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		board := <-c
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for i := 0; i < utils.ScreenWidth + 5; i+=utils.Width {
			for j := 0; j < utils.ScreenHeight + 5; j+=utils.Height {
				col = rl.White
				if board[i][j] {
					col = rl.Black
				}
				rl.DrawRectangle(int32(i), int32(j), utils.Width, utils.Height, col)
			}
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()

}
