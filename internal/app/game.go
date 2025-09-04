package app

import (
	"github.com/Crusazer/tanks-race/internal/app/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sceneManager *scenes.SceneManager
}

func New() *Game {
	g := &Game{}
	g.sceneManager = &scenes.SceneManager{}
	g.sceneManager.ChangeScene(scenes.NewMenuScene(g.sceneManager))
	return g
}

func (g *Game) Update() error {
	return g.sceneManager.Update(1.0 / 60)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return g.sceneManager.Layout(w, h)
}
