package app

import (
	"image/color"
	"log"

	"github.com/Crusazer/tanks-race/internal/game/prefabs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	tank *prefabs.Tank
}

func New() *Game {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Tanks Race")

	tank, err := prefabs.NewTank()
	if err != nil {
		log.Fatal(err)
	}

	return &Game{tank: tank}
}

func (g *Game) Update() error {
	dt := 1.0 / 60.0
	g.tank.Move(dt)
	g.tank.UpdateAiming()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 50, 50, 255}) // Фон
	g.tank.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 800, 600
}
