package main

import (
	"context"
	"game_of_life/game"
	"log"
)

func init() {
	// log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}


func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan [][]bool)
	game.Run(c, ctx)
	cancel()
}
