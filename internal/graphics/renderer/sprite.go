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

func (s *Sprite) Draw(screen *ebiten.Image, cam *Camera) {
	op := &ebiten.DrawImageOptions{}

	// 1. Origin → 0,0
	op.GeoM.Translate(-s.Origin.X, -s.Origin.Y)
	// 2. Поворот вокруг Origin
	op.GeoM.Rotate(s.Rotation)
	// 3. Масштаб спрайта + масштаб камеры
	op.GeoM.Scale(s.Scale.X*cam.Zoom, s.Scale.Y*cam.Zoom)
	// 4. Перенос в экранный центр
	scr := cam.WorldToScreen(s.Position)
	op.GeoM.Translate(scr.X, scr.Y)

	screen.DrawImage(s.Image, op)
}
