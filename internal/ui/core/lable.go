package core

// import (
// 	"image/color"

// 	m "github.com/Crusazer/tanks-race/pkg/math"
// 	"github.com/Crusazer/tanks-race/pkg/resources"
// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/text/v2"
// )

// type Label struct {
// 	BaseWidget
// 	Text      string
// 	TextColor color.Color
// 	Font      *text.GoTextFace // Используем ваш шрифт
// }

// func NewLabel(text string, textColor color.Color) *Label {
// 	return &Label{
// 		BaseWidget: BaseWidget{
// 			Visible: true,
// 		},
// 		Text:      text,
// 		TextColor: textColor,
// 		Font:      resources.UIFont, // Ваш шрифт
// 	}
// }

// func (l *Label) PreferredSize() (w, h int) {
// 	op := &text.DrawOptions{}
// 	op.LineSpacing = l.Font.LineHeight
// 	width, height := text.Measure(l.Text, l.Font, op.LineSpacing)
// 	return int(width), int(height)
// }

// func (l *Label) Update() {
// 	// Label не имеет динамического состояния
// }

// func (l *Label) HandleMouseEvent(mousePos m.Vector2, isPressed bool, isReleased bool) {
// 	// Label не реагирует на мышь
// }

// func (l *Label) HandleKeyboardEvent(key ebiten.Key, r rune) {
// 	// Label не реагирует на клавиатуру
// }

// func (l *Label) SetFocused(focused bool) {
// 	// Label не может быть в фокусе
// }

// func (l *Label) IsFocusable() bool {
// 	return false // Label не может получать фокус
// }

// func (l *Label) Draw(dst *ebiten.Image) {
// 	if !l.Visible {
// 		return
// 	}

// 	op := &text.DrawOptions{}
// 	w, h := text.Measure(l.Text, l.Font, op.LineSpacing)

// 	// Центрируем текст внутри bounds (или можно задать выравнивание)
// 	op.GeoM.Translate(
// 		l.Pos.X+float64(l.Width)/2-w/2,
// 		l.Pos.Y+float64(l.Height)/2-h/2,
// 	)
// 	op.ColorScale.SetColor(l.TextColor)
// 	text.Draw(dst, l.Text, l.Font, op)
// }
