package app

import (
	"github.com/Crusazer/tanks-race/internal/app/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sceneMgr *scenes.SceneManager
}

func New() *Game {
	g := &Game{}
	g.sceneMgr = scenes.NewSceneManager(scenes.NewPlayingScene()) // Старт с игры
	return g
}

func (g *Game) Update() error {
	return g.sceneMgr.Update(1.0 / 60)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneMgr.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return g.sceneMgr.Layout(w, h)
}
