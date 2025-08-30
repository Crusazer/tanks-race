package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type MenuScene struct{}

func NewMenuScene() Scene { return &MenuScene{} }

func (m *MenuScene) Enter()  {}
func (m *MenuScene) Exit()   {}
func (m *MenuScene) Update(dt float64) error {
	// TODO: обработка кнопок, переход в PlayingScene
	return nil
}
func (m *MenuScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 40, 60, 255})
	// TODO: рисуем меню
}
func (m *MenuScene) Layout(w, h int) (int, int) { return w, h }