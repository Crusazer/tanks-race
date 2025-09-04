package core

// import (
// 	"image/color"
// 	"time"

// 	"github.com/Crusazer/tanks-race/internal/input"
// 	m "github.com/Crusazer/tanks-race/pkg/math"
// 	"github.com/Crusazer/tanks-race/pkg/resources"
// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/text/v2"
// )

// const cursorBlinkRate = time.Millisecond * 500

// type TextInput struct {
// 	BaseWidget
// 	Text             string
// 	Placeholder      string
// 	IsFocused        bool
// 	BorderColor      color.Color
// 	BgColor          color.Color
// 	TextColor        color.Color
// 	PlaceholderColor color.Color
// 	Font             *text.GoTextFace
// 	theme            *Theme

// 	cursorTick int // Для мигания курсора
// 	lastUpdate time.Time
// }

// func NewTextInput(width, height int, placeholder string, theme *Theme) *TextInput {
// 	if theme == nil {
// 		theme = DefaultTheme
// 	}
// 	return &TextInput{
// 		BaseWidget: BaseWidget{
// 			Width:     width,
// 			Height:    height,
// 			Visible:   true,
// 			IsFocused: false,
// 		},
// 		Placeholder:      placeholder,
// 		BorderColor:      theme.TextInputBorderColor,
// 		BgColor:          theme.TextInputBgColor,
// 		TextColor:        theme.TextInputTextColor,
// 		PlaceholderColor: theme.TextInputPlaceholderColor,
// 		Font:             resources.UIFont,
// 		theme:            theme,
// 		lastUpdate:       time.Now(),
// 	}
// }

// func (t *TextInput) PreferredSize() (w, h int) {
// 	return t.Width, t.Height
// }

// func (t *TextInput) Update() {
// 	if t.IsFocused {
// 		if time.Since(t.lastUpdate) >= cursorBlinkRate {
// 			t.cursorTick++
// 			t.lastUpdate = time.Now()
// 		}
// 	} else {
// 		t.cursorTick = 0 // Сбрасываем курсор, если не в фокусе
// 	}
// }

// func (t *TextInput) HandleMouseEvent(mousePos m.Vector2, isPressed bool, isReleased bool) {
// 	// Фокус устанавливается родительским компонентом (Screen)
// 	// здесь просто обновляем состояние, если нужно (например, подсветка при hover)
// 	// TextInput обычно не меняет состояние на основе hover, только focus
// }

// func (t *TextInput) HandleAction(action input.UIAction) {
// 	if !t.IsFocused || !t.Visible {
// 		return
// 	}
// 	if action == input.UIActionBackspace {
// 		if len(t.Text) > 0 {
// 			// Это простая реализация, для UTF-8 лучше использовать руны
// 			runes := []rune(t.Text)
// 			t.Text = string(runes[:len(runes)-1])
// 		}
// 	}
// 	// Можно добавить реакцию на UIActionConfirm, например, вызвать коллбэк OnConfirm
// }

// func (t *TextInput) HandleChars(chars []rune) {
// 	if !t.IsFocused || !t.Visible {
// 		return
// 	}
// 	t.Text += string(chars)
// }

// func (t *TextInput) SetFocused(focused bool) {
// 	t.BaseWidget.SetFocused(focused)
// 	if !focused {
// 		t.cursorTick = 0
// 	}
// }

// func (t *TextInput) IsFocusable() bool {
// 	return true
// }

// func (t *TextInput) Draw(dst *ebiten.Image) {
// 	if !t.Visible {
// 		return
// 	}

// 	// Фон
// 	drawRect(dst, t.Pos, float64(t.Width), float64(t.Height), t.BgColor)

// 	// Рамка
// 	borderColor := t.BorderColor
// 	if t.IsFocused {
// 		borderColor = t.theme.FocusBorderColor
// 	}
// 	drawBorder(dst, t.Pos, t.Width, t.Height, borderColor, 2)

// 	// Текст
// 	drawOptions := &text.DrawOptions{}
// 	drawOptions.LineSpacing = t.Font.LineHeight

// 	displayTxt := t.Text
// 	txtColor := t.TextColor
// 	if t.Text == "" && !t.IsFocused {
// 		displayTxt = t.Placeholder
// 		txtColor = t.PlaceholderColor
// 	}

// 	// Отступ текста от края
// 	textPadding := float64(t.Height) / 4 // Примерный отступ
// 	textX := t.Pos.X + textPadding
// 	textY := t.Pos.Y + float64(t.Height)/2 - text.Measure(displayTxt, t.Font, drawOptions.LineSpacing)/2 // Центрирование по вертикали

// 	drawOptions.GeoM.Translate(textX, textY)
// 	drawOptions.ColorScale.SetColor(txtColor)
// 	text.Draw(dst, displayTxt, t.Font, drawOptions)

// 	// Курсор
// 	if t.IsFocused && t.cursorTick%2 == 0 { // Мигает
// 		cursorWidth := 2.0
// 		cursorHeight := float64(t.Height) - textPadding*2

// 		// Вычисляем позицию курсора после текста
// 		textWidth, _ := text.Measure(displayTxt, t.Font, drawOptions.LineSpacing)
// 		cursorX := textX + textWidth
// 		cursorY := t.Pos.Y + textPadding

// 		drawRect(dst, m.Vector2{X: cursorX, Y: cursorY}, cursorWidth, cursorHeight, t.TextColor)
// 	}
// }

// // drawBorder - вспомогательная функция для рисования рамки
// func drawBorder(dst *ebiten.Image, position m.Vector2, width, height int, clr color.Color, thickness int) {
// 	// Top border
// 	drawRect(dst, position, float64(width), float64(thickness), clr)
// 	// Bottom border
// 	drawRect(dst, m.Vector2{X: position.X, Y: position.Y + float64(height) - float64(thickness)}, float64(width), float64(thickness), clr)
// 	// Left border
// 	drawRect(dst, position, float64(thickness), float64(height), clr)
// 	// Right border
// 	drawRect(dst, m.Vector2{X: position.X + float64(width) - float64(thickness), Y: position.Y}, float64(thickness), float64(height), clr)
// }
