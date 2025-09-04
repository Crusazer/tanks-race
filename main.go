package main

import (
    "log"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/Crusazer/tanks-race/internal/app"
)

func main() {
    ebiten.SetWindowSize(800, 600)
    ebiten.SetWindowTitle("Tanks Race")
    ebiten.SetWindowResizable(true)

    game := app.New()
    
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}