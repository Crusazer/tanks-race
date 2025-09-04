package core

import (
	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Widget interface {
	SetBounds(x, y, w, h int)
	Bounds() (x, y, w, h int)
	Update(mausPos m.Vector2, isPressed bool)
	Draw(dst *ebiten.Image)
}
