package renderer

import m "github.com/Crusazer/tanks-race/pkg/math"

type Camera struct {
	Position       m.Vector2
	Zoom           float64
	ViewportWidth  float64
	ViewportHeight float64
}

func (c *Camera) ViewMatrix() m.Matrix2 {
	s := c.Zoom
	return m.Matrix2{
		{s, 0}, // масштаб X
		{0, s}, // масштаб Y
	}
}

func (c *Camera) WorldToScreen(world m.Vector2) m.Vector2 {
	// только сдвиг, без Zoom (Zoom уже в Sprite.Draw)
	return world.Sub(c.Position).Add(
		m.Vector2{X: c.ViewportWidth / 2, Y: c.ViewportHeight / 2},
	)
}

func (c *Camera) ScreenToWorld(screen m.Vector2) m.Vector2 {
	return screen.Sub(m.Vector2{
		X: c.ViewportWidth / 2,
		Y: c.ViewportHeight / 2,
	}).Div(c.Zoom).Add(c.Position)
}
