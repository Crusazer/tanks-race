package main

import (
	"log"

	"github.com/Crusazer/tanks-race/internal/app"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := app.New()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
