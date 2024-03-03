package main

import (
	"context"
	"image/color"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenWidth = int32(1280)
const screenHeight = int32(720)
const width = 100
const height = 100

func show(c chan *[screenWidth][screenHeight]bool)  {
	var col color.RGBA

	rl.InitWindow(screenWidth, screenHeight, "Conway's Game of Life")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		board := *(<-c)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for i := 0; i < int(screenWidth); i+=width {
			for j := 0; j < int(screenHeight); j+=height {
				col = rl.White
				if board[i][j] {
					col = rl.Black
				}
				rl.DrawRectangle(int32(i), int32(j), width, height, col)
			}
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()

}

func trueOrFalse() bool {
	if rand.Intn((1-0+1)+0) == 1 {
		return true
	}
	return false
}

func updateValues(c chan *[screenWidth][screenHeight]bool, board [screenWidth][screenHeight]bool) {
	c <- &board
	for i := 0; i < int(screenWidth); i++ {
		for j := 0; j < int(screenHeight); j++ {
			board[i][j] = trueOrFalse()	
		}
	}
	time.Sleep(500 * time.Millisecond)
}

func update(c chan *[screenWidth][screenHeight]bool, ctx context.Context) {
	var board [screenWidth][screenHeight]bool
	start(&board)
	for {
		select{
		case <- ctx.Done():
			return
		default:
			updateValues(c, board)
		}
	}
}

func start(board *[screenWidth][screenHeight]bool) {
	for i := 0; i < int(screenWidth); i++ {
		for j := 0; j < int(screenHeight); j++ {
			board[i][j] = trueOrFalse()	
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan *[screenWidth][screenHeight]bool)
	go update(c, ctx)
	// wg.Add(2)
	show(c)
	cancel()
}
