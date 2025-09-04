package core

import (
	"image/color"

	"github.com/Crusazer/tanks-race/internal/input"
	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/Crusazer/tanks-race/pkg/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Label struct {
	BaseWidget               // встраиваем базовый виджет (позиция, размер, видимость)
	Text        string       // отображаемый текст
	Color       color.Color  // цвет текста
	Font        text.Face    // шрифт (используем общий UI‑шрифт)
	theme       *Theme       // ссылка на тему (чтобы можно было менять цвет через неё)
}

// NewLabel создаёт метку. Тема передаётся явно, но если её нет – берём DefaultTheme.
func NewLabel(text string, theme *Theme) *Label {
	if theme == nil {
		theme = DefaultTheme
	}
	return &Label{
		BaseWidget: BaseWidget{
			Visible: true,
		},
		Text:  text,
		Color: theme.LabelTextColor,
		Font:  resources.UIFont, // общий UI‑шрифт
		theme: theme,
	}
}

func (l *Label) PreferredSize() (w, h int) {
	if l.Font == nil {
		l.Font = resources.UIFont
	}
	op := &text.DrawOptions{}
	wf, hf := text.Measure(l.Text, l.Font, op.LineSpacing)
	return int(wf), int(hf)
}

// Update – у метки нет динамического состояния, но метод обязателен.
func (l *Label) Update() { /* no‑op */ }

// HandleMouseEvent – метка не реагирует на мышь.
func (l *Label) HandleMouseEvent(mousePos m.Vector2, isPressed bool, isReleased bool) {
	/* no‑op */
}

// HandleAction – метка не получает действий (Enter, Backspace и т.п.).
func (l *Label) HandleAction(action input.UIAction) {
	/* no‑op */
}

// HandleChars – метка не принимает ввод символов.
func (l *Label) HandleChars(chars []rune) {
	/* no‑op */
}

// SetFocused – метка не может быть в фокусе, но реализуем метод, чтобы соответствовать интерфейсу.
func (l *Label) SetFocused(focused bool) {
	// Мы просто игнорируем, но сохраняем состояние в BaseWidget
	l.BaseWidget.SetFocused(focused)
}

// IsFocusable – метка не интерактивна.
func (l *Label) IsFocusable() bool { return false }

func (l *Label) Draw(dst *ebiten.Image) {
	if !l.Visible {
		return
	}

	// Если в будущем захотим подсвечивать метку при фокусе,
	// можно добавить рамку аналогично Button/TextInput.
	// Сейчас просто рисуем текст.

	// Подготовка опций
	op := &text.DrawOptions{}
	op.LineSpacing = l.Font.Metrics().HLineGap // используем высоту строки шрифта

	// Центрируем текст внутри текущих границ виджета
	w, _ := text.Measure(l.Text, l.Font, op.LineSpacing)
	x := l.Pos.X + float64(l.Width)/2 - w/2
	y := l.Pos.Y + float64(l.Height)/2
	op.GeoM.Translate(x, y)

	// Применяем цвет (ColorScale – умножаем на нужный цвет)
	op.ColorScale.ScaleWithColor(l.Color)

	text.Draw(dst, l.Text, l.Font, op)
}