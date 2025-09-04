package screens

import (
	"fmt"

	"github.com/Crusazer/tanks-race/internal/input"
	"github.com/Crusazer/tanks-race/internal/ui/core"

	"github.com/hajimehoshi/ebiten/v2"
)

type Lobby struct {
	layout        core.Layout
	inputSystem   *input.InputSystem
	focusedWidget core.Widget

	ipInput       *core.TextInput
	portInput     *core.TextInput
	connectButton *core.Button
	statusLabel   *core.Label

	theme *core.Theme
}

func NewLobby(theme *core.Theme) *Lobby {
	if theme == nil {
		theme = core.DefaultTheme
	}

	layout := core.NewVerticalFlowLayout(30, 10)

	ipInput := core.NewTextInput(240, 40, "IP адрес", theme)
	ipInput.Text = "127.0.0.1"
	layout.Add(ipInput)

	portInput := core.NewTextInput(240, 40, "Порт", theme)
	portInput.Text = "7777"
	layout.Add(portInput)

	connectButton := core.NewButton("Подключиться", 240, 40, func() {
		fmt.Printf("Подключение к %s:%s...\n", ipInput.Text, portInput.Text)
		// TODO: Реализовать логику подключения
	}, theme)
	layout.Add(connectButton)
	
	statusLabel := core.NewLabel("Статус: Ожидание ввода", theme)
	layout.Add(statusLabel)

	return &Lobby{
		layout:        layout,
		inputSystem:   input.NewInputSystem(),
		ipInput:       ipInput,
		portInput:     portInput,
		connectButton: connectButton,
		statusLabel:   statusLabel,
		theme:         theme,
	}
}

func (l *Lobby) Update() error {
	l.inputSystem.UpdateUI()

	mousePos := l.inputSystem.GetMousePosition()
	isMousePressed := l.inputSystem.IsUIActionJustPressed(input.UIActionSelect)
	isMouseHeld := l.inputSystem.IsUIActionPressed(input.UIActionSelect)
	isMouseReleased := l.inputSystem.IsUIActionJustReleased(input.UIActionSelect)

	// Обработка мыши и фокуса
	hoveredWidget := core.Widget(nil)
	for _, w := range l.layout.Widgets() {
		if w.IsVisible() {
			w.HandleMouseEvent(mousePos, isMouseHeld, isMouseReleased)

			if isMousePressed && w.Contains(mousePos) && w.IsFocusable() {
				l.setFocus(w)
			}
			if w.Contains(mousePos) {
				hoveredWidget = w
			}
		}
	}

	if isMousePressed && hoveredWidget == nil {
		l.setFocus(nil)
	}

	// Обработка навигации по клавиатуре
	if l.inputSystem.IsUIActionJustPressed(input.UIActionNavigateUp) {
		l.focusPreviousWidget()
	}
	if l.inputSystem.IsUIActionJustPressed(input.UIActionNavigateDown) {
		l.focusNextWidget()
	}

	// Если нет фокуса, но была попытка навигации или подтверждения — ставим фокус на первый виджет
	if l.focusedWidget == nil &&
		(l.inputSystem.IsUIActionJustPressed(input.UIActionNavigateUp) ||
			l.inputSystem.IsUIActionJustPressed(input.UIActionNavigateDown) ||
			l.inputSystem.IsUIActionJustPressed(input.UIActionConfirm)) {
		l.focusFirstWidget()
	}

	// Передача действий и символов сфокусированному виджету
	if l.focusedWidget != nil {
		if l.inputSystem.IsUIActionJustPressed(input.UIActionConfirm) {
			l.focusedWidget.HandleAction(input.UIActionConfirm)
		}
		if l.inputSystem.IsUIActionJustPressed(input.UIActionBackspace) {
			l.focusedWidget.HandleAction(input.UIActionBackspace)
		}

		chars := l.inputSystem.GetInputChars()
		if len(chars) > 0 {
			l.focusedWidget.HandleChars(chars)
		}
	}

	// Обновление виджетов
	for _, w := range l.layout.Widgets() {
		if w.IsVisible() {
			w.Update()
		}
	}

	return nil
}

func (l *Lobby) Draw(screen *ebiten.Image) {
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	l.layout.ComputeBounds(w, h)

	for _, w := range l.layout.Widgets() {
		if w.IsVisible() {
			w.Draw(screen)
		}
	}
}

// Фокус и навигация

func (l *Lobby) setFocus(w core.Widget) {
	if l.focusedWidget == w {
		return
	}
	if l.focusedWidget != nil {
		l.focusedWidget.SetFocused(false)
	}
	l.focusedWidget = w
	if l.focusedWidget != nil {
		l.focusedWidget.SetFocused(true)
	}
}

func (l *Lobby) getFocusableWidgets() []core.Widget {
	var result []core.Widget
	for _, w := range l.layout.Widgets() {
		if w.IsVisible() && w.IsFocusable() {
			result = append(result, w)
		}
	}
	return result
}

func (l *Lobby) focusNextWidget() {
	focusable := l.getFocusableWidgets()
	if len(focusable) == 0 {
		return
	}

	var currentIndex int = -1
	for i, w := range focusable {
		if w == l.focusedWidget {
			currentIndex = i
			break
		}
	}

	nextIndex := 0
	if currentIndex != -1 {
		nextIndex = (currentIndex + 1) % len(focusable)
	}

	l.setFocus(focusable[nextIndex])
}

func (l *Lobby) focusPreviousWidget() {
	focusable := l.getFocusableWidgets()
	if len(focusable) == 0 {
		return
	}

	var currentIndex int = -1
	for i, w := range focusable {
		if w == l.focusedWidget {
			currentIndex = i
			break
		}
	}

	prevIndex := len(focusable) - 1
	if currentIndex != -1 {
		prevIndex = (currentIndex - 1 + len(focusable)) % len(focusable)
	}

	l.setFocus(focusable[prevIndex])
}

func (l *Lobby) focusFirstWidget() {
	focusable := l.getFocusableWidgets()
	if len(focusable) > 0 {
		l.setFocus(focusable[0])
	}
}