package core

import (
	"github.com/Crusazer/tanks-race/internal/input"
	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/hajimehoshi/ebiten/v2"
)

// BaseWidget предоставляет общие поля и базовые методы для всех виджетов.
// Другие виджеты могут встраивать эту структуру.
type BaseWidget struct {
	Pos       m.Vector2
	Width     int
	Height    int
	Visible   bool
	IsFocused bool
}

func (b *BaseWidget) SetBounds(x, y, w, h int) {
	b.Pos.X, b.Pos.Y = float64(x), float64(y)
	b.Width, b.Height = w, h
}

func (b *BaseWidget) Bounds() (x, y, w, h int) {
	return int(b.Pos.X), int(b.Pos.Y), b.Width, b.Height
}

func (b *BaseWidget) Contains(p m.Vector2) bool {
	return b.Visible &&
		p.X >= b.Pos.X && p.X <= b.Pos.X+float64(b.Width) &&
		p.Y >= b.Pos.Y && p.Y <= b.Pos.Y+float64(b.Height)
}

func (b *BaseWidget) IsVisible() bool {
	return b.Visible
}

func (b *BaseWidget) SetVisible(visible bool) {
	b.Visible = visible
}

func (b *BaseWidget) SetFocused(focused bool) {
	b.IsFocused = focused
}

func (b *BaseWidget) IsFocusable() bool {
	return true // По умолчанию, но переопределяется в виджетах (например, Label - false)
}

// Widget - основной интерфейс для всех элементов UI.
type Widget interface {
	SetBounds(x, y, w, h int)
	Bounds() (x, y, w, h int)
	Contains(p m.Vector2) bool
	PreferredSize() (w, h int)

	Update()
	Draw(dst *ebiten.Image)

	IsVisible() bool
	SetVisible(visible bool)

	// === НОВЫЕ МЕТОДЫ ДЛЯ ВВОДА ===
	HandleMouseEvent(mousePos m.Vector2, isPressed bool, isReleased bool)
	HandleAction(action input.UIAction) // Обработка дискретных действий (Enter, Backspace)
	HandleChars(chars []rune)           // Обработка ввода текста

	SetFocused(focused bool)
	IsFocusable() bool
}
