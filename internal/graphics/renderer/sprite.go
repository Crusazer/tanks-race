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

func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Точка поворота (центр изображения)
	point := s.Image.Bounds().Size()
	op.GeoM.Translate(-float64(point.X)/2, -float64(point.Y)/2)

	op.GeoM.Rotate(s.Rotation)
	op.GeoM.Scale(s.Scale.X, s.Scale.Y)

	// Перемещаем в нужную позицию
	op.GeoM.Translate(s.Position.X, s.Position.Y)

	screen.DrawImage(s.Image, op)
}
