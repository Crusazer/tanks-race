package core

import (
	"image/color"

	"github.com/Crusazer/tanks-race/internal/input"
	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/Crusazer/tanks-race/pkg/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ButtonStatus int

const (
	StatusActive ButtonStatus = iota
	StatusHover
	StatusPressed
)

type Button struct {
	BaseWidget

	Text   string
	Status ButtonStatus

	OnClick      func()
	NormalColor  color.Color
	HoverColor   color.Color
	PressedColor color.Color

	theme        *Theme 
}

func NewButton(text string, width, height int, onClick func(), theme *Theme) *Button {
	if theme == nil {
		theme = DefaultTheme
	}
	return &Button{
		BaseWidget: BaseWidget{
			Width:     width,
			Height:    height,
			Visible:   true,
			IsFocused: false,
		},
		Text:         text,
		NormalColor:  theme.ButtonNormalColor,
		HoverColor:   theme.ButtonHoverColor,
		PressedColor: theme.ButtonPressedColor,
		OnClick:      onClick,
		theme:        theme,
	}
}

// PreferredSize возвращает предпочтительный размер кнопки.
// В данном случае это просто заданные ширина и высота.
func (b *Button) PreferredSize() (w, h int) {
	return b.Width, b.Height
}

func (b *Button) Update() {
	// No-op for now, button state is handled by HandleMouseEvent
}

func (b *Button) HandleMouseEvent(mousePos m.Vector2, isPressed bool, isReleased bool) {
	if !b.Visible {
		return
	}

	hover := b.Contains(mousePos)

	switch {
	case hover && isPressed:
		b.Status = StatusPressed
	case hover && isReleased: // Клик происходит при отпускании кнопки мыши над виджетом
		if b.Status == StatusPressed && b.OnClick != nil {
			b.OnClick()
		}
		b.Status = StatusHover // После клика, если все еще hover
	case hover: // Просто hover, но не нажата
		b.Status = StatusHover
	default: // Не hover
		b.Status = StatusActive
	}
}

func (b *Button) SetFocused(focused bool) {
	b.BaseWidget.SetFocused(focused)
}

func (b *Button) IsFocusable() bool {
	return true // Кнопки обычно можно "фокусировать" для навигации с клавиатуры
}

func (b *Button) HandleAction(action input.UIAction) {
	if !b.Visible {
		return
	}
	// Кнопка активируется по действию Confirm (Enter) или Select (если он пришел от клавиатуры)
	if action == input.UIActionConfirm {
		if b.OnClick != nil {
			b.OnClick()
		}
	}
}

// HandleChars для Button. Кнопка игнорирует ввод символов.
func (b *Button) HandleChars(chars []rune) {
	// No-op
}

func (b *Button) Draw(dst *ebiten.Image) {
	if !b.Visible {
		return
	}

	var clr color.Color
	switch b.Status {
	case StatusPressed:
		clr = b.PressedColor
	case StatusHover:
		clr = b.HoverColor
	default:
		clr = b.NormalColor
	}

	// Рисуем фон
	drawRect(dst, b.Pos, float64(b.Width), float64(b.Height), clr)

	// Рисуем текст
	drawText(dst, b.Text,
		b.Pos.X+float64(b.Width)/2,
		b.Pos.Y+float64(b.Height)/2)

	// Если в фокусе, рисуем рамку
	if b.IsFocused {
		drawBorder(dst, b.Pos, b.Width, b.Height, b.theme.FocusBorderColor, 2)
	}
}

// Вспомогательная функция для рамки
func drawBorder(dst *ebiten.Image, position m.Vector2, width, height int, clr color.Color, thickness int) {
	// Top
	drawRect(dst, position, float64(width), float64(thickness), clr)
	// Bottom
	drawRect(dst, m.Vector2{X: position.X, Y: position.Y + float64(height) - float64(thickness)}, float64(width), float64(thickness), clr)
	// Left
	drawRect(dst, position, float64(thickness), float64(height), clr)
	// Right
	drawRect(dst, m.Vector2{X: position.X + float64(width) - float64(thickness), Y: position.Y}, float64(thickness), float64(height), clr)
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
	op.GeoM.Translate(
		x-w/2,
		y-h/2,
	)
	text.Draw(dst, str, resources.UIFont, op)
}
