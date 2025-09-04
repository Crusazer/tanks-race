package core

import (
	"image/color"

	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/Crusazer/tanks-race/pkg/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ButtonStatus int

const (
	StatusActive ButtonStatus = iota
	StatusInactive
	StatusHover
	StatusPressed
)

type Button struct {
	Text   string
	Pos    m.Vector2
	Width  int
	Height int
	Status ButtonStatus

	OnClick      func()
	NormalColor  color.Color
	HoverColor   color.Color
	PressedColor color.Color
}

func (b *Button) SetBounds(x, y, w, h int) {
	b.Pos.X, b.Pos.Y = float64(x), float64(y)
	b.Width, b.Height = w, h
}

func (b *Button) Bounds() (x, y, w, h int) {
	return int(b.Pos.X), int(b.Pos.Y), b.Width, b.Height
}

func (b *Button) Contains(p m.Vector2) bool {
	return p.X >= b.Pos.X && p.X <= b.Pos.X+float64(b.Width) &&
		p.Y >= b.Pos.Y && p.Y <= b.Pos.Y+float64(b.Height)
}

func (b *Button) Update(mousePosition m.Vector2, isPressed bool) {
	hover := b.Contains(mousePosition)

	switch {
	case hover && isPressed:
		b.Status = StatusPressed
	case hover && !isPressed:
		if b.Status == StatusPressed && b.OnClick != nil {
			b.OnClick()
		}
		b.Status = StatusHover
	case !hover:
		b.Status = StatusActive
	}
}

func (b *Button) Draw(dst *ebiten.Image) {
	var clr color.Color
	switch b.Status {
	case StatusPressed:
		clr = b.PressedColor
	case StatusHover:
		clr = b.HoverColor
	default:
		clr = b.NormalColor
	}

	drawRect(dst, b.Pos, float64(b.Width), float64(b.Height), clr)
	drawText(dst, b.Text,
		b.Pos.X+float64(b.Width)/2,
		b.Pos.Y+float64(b.Height)/2)
}

func drawRect(dst *ebiten.Image, position m.Vector2, width, height float64, clr color.Color) {
	rect := ebiten.NewImage(int(width), int(height))
	rect.Fill(clr)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(position.X, position.Y)
	dst.DrawImage(rect, op)
}

func drawText(dst *ebiten.Image, str string, x, y float64) {
	op := &text.DrawOptions{}
	w, h := text.Measure(str, resources.UIFont, op.LineSpacing)
	// центрируем:  X = середина кнопки,  Y = середина кнопки минус половина высоты текста
	op.GeoM.Translate(
		x-w/2, // по горизонтали
		y-h/2, // по вертикали
	)

	text.Draw(dst, str, resources.UIFont, op)
}
