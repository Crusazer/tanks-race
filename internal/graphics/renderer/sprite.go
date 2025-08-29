package renderer

import (
	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image    *ebiten.Image
	Position m.Vector2
	Rotation float64 // radians
	Scale    m.Vector2
	Origin   m.Vector2
}

// internal/graphics/renderer/sprite.go
func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-s.Origin.X, -s.Origin.Y)
	op.GeoM.Rotate(s.Rotation)
	op.GeoM.Scale(s.Scale.X, s.Scale.Y)

	// Переносим сам Origin в нужное место мира
	op.GeoM.Translate(s.Position.X, s.Position.Y)

	screen.DrawImage(s.Image, op)
}
