package app

import (
	"image/color"
	"log"

	"github.com/Crusazer/tanks-race/internal/game/prefabs"
	"github.com/Crusazer/tanks-race/internal/graphics/renderer"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	tank             *prefabs.Tank
	camera           *renderer.Camera
	screenW, screenH int
}

func New() *Game {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Tanks Race")

	tank, err := prefabs.NewTank()
	if err != nil {
		log.Fatal(err)
	}

	return &Game{
		tank: tank,
		camera: &renderer.Camera{
			Position: tank.Body.Position,
			Zoom:     0.6,
		},
	}
}

func (g *Game) Update() error {
	dt := 1.0 / 60.0
	g.tank.Move(dt)
	g.tank.UpdateAiming(g.camera)

	// камера следит за танком
	g.camera.Position = g.tank.Body.Position
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{150, 150, 150, 255}) // Фон
	g.tank.Draw(screen, g.camera)
}

func (g *Game) Layout(w, h int) (int, int) {
	g.screenW, g.screenH = w, h
	g.camera.ViewportWidth = float64(w)
	g.camera.ViewportHeight = float64(h)
	return w, h
}
