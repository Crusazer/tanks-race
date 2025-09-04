package core

import (
	"image/color"

	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Slider struct {
	BaseWidget
	Value         float64
	Min           float64
	Max           float64
	Step          float64
	TrackColor    color.Color
	HandleColor   color.Color
	IsDragging    bool
	OnValueChange func(float64)
}

func NewSlider(width, height int, min, max, initialValue, step float64, onValueChange func(float64)) *Slider {
	return &Slider{
		BaseWidget: BaseWidget{
			Width:   width,
			Height:  height,
			Visible: true,
		},
		Min:           min,
		Max:           max,
		Step:          step,
		Value:         initialValue,
		TrackColor:    color.RGBA{100, 100, 100, 255},
		HandleColor:   color.RGBA{200, 200, 200, 255},
		OnValueChange: onValueChange,
	}
}

func (s *Slider) PreferredSize() (w, h int) {
	return s.Width, s.Height
}

func (s *Slider) Update() {
	// Slider не имеет динамического состояния для Update()
}

func (s *Slider) HandleMouseEvent(mousePos m.Vector2, isPressed bool, isReleased bool) {
	if !s.Visible {
		return
	}

	if s.IsDragging {
		if isReleased {
			s.IsDragging = false
		} else {
			// Обновляем значение слайдера при перетаскивании
			normalizedPos := (mousePos.X - s.Pos.X) / float64(s.Width)
			newValue := s.Min + normalizedPos*(s.Max-s.Min)

			// Округляем до шага
			if s.Step > 0 {
				newValue = float64(int(newValue/s.Step)) * s.Step
			}

			// Ограничиваем значение
			if newValue < s.Min {
				newValue = s.Min
			}
			if newValue > s.Max {
				newValue = s.Max
			}

			if s.Value != newValue {
				s.Value = newValue
				if s.OnValueChange != nil {
					s.OnValueChange(s.Value)
				}
			}
		}
	} else if isPressed && s.Contains(mousePos) {
		s.IsDragging = true
		// Сразу же устанавливаем значение при нажатии
		s.HandleMouseEvent(mousePos, true, false) // Рекурсивный вызов для установки значения
	}
}

func (s *Slider) SetFocused(focused bool) {
	// Слайдер не имеет специального состояния фокуса для клавиатурного ввода
}

func (s *Slider) IsFocusable() bool {
	return false // Или true, если вы хотите управлять слайдером с клавиатуры (стрелками)
}

func (s *Slider) HandleKeyboardEvent(key ebiten.Key, r rune) {
	// Если IsFocusable() == true, здесь можно обрабатывать стрелки влево/вправо
}

func (s *Slider) Draw(dst *ebiten.Image) {
	if !s.Visible {
		return
	}

	trackHeight := float64(s.Height) / 4
	trackY := s.Pos.Y + float64(s.Height)/2 - trackHeight/2

	// Рисуем трек
	drawRect(dst, m.Vector2{X: s.Pos.X, Y: trackY}, float64(s.Width), trackHeight, s.TrackColor)

	// Вычисляем позицию ручки
	normalizedValue := (s.Value - s.Min) / (s.Max - s.Min)
	handleWidth := float64(s.Width) / 10
	handleHeight := float64(s.Height)
	handleX := s.Pos.X + normalizedValue*float64(s.Width) - handleWidth/2
	handleY := s.Pos.Y // Центрируем ручку по высоте слайдера

	// Ограничиваем ручку краями трека
	if handleX < s.Pos.X {
		handleX = s.Pos.X
	}
	if handleX+handleWidth > s.Pos.X+float64(s.Width) {
		handleX = s.Pos.X + float64(s.Width) - handleWidth
	}

	// Рисуем ручку
	drawRect(dst, m.Vector2{X: handleX, Y: handleY}, handleWidth, handleHeight, s.HandleColor)
}
