package screens

// import (
// 	"fmt"
// 	"image/color"

// 	"github.com/Crusazer/tanks-race/internal/input"
// 	"github.com/Crusazer/tanks-race/internal/ui/core" // Обновленный core

// 	"github.com/hajimehoshi/ebiten/v2"
// )

// type Lobby struct {
// 	layout        core.Layout
// 	inputSystem   *input.InputSystem
// 	focusedWidget core.Widget // Управление фокусом
// 	ipInput       *core.TextInput
// 	portInput     *core.TextInput
// 	connectButton *core.Button
// 	statusLabel   *core.Label
// }

// func NewLobby() *Lobby {
// 	layout := core.NewVerticalFlowLayout(30, 10)

// 	ipInput := core.NewTextInput(240, 40, "IP адрес")
// 	ipInput.Text = "127.0.0.1" // Значение по умолчанию
// 	layout.Add(ipInput)

// 	portInput := core.NewTextInput(240, 40, "Порт")
// 	portInput.Text = "7777" // Значение по умолчанию
// 	layout.Add(portInput)

// 	connectButton := core.NewButton("Подключиться", 240, 40, func() {
// 		fmt.Printf("Подключение к %s:%s...\n", ipInput.Text, portInput.Text)
// 		// Здесь будет логика подключения
// 	})
// 	layout.Add(connectButton)

// 	statusLabel := core.NewLabel("Статус: Ожидание ввода", color.White)
// 	layout.Add(statusLabel)

// 	return &Lobby{
// 		layout:        layout,
// 		inputSystem:   input.NewInputSystem(),
// 		ipInput:       ipInput,
// 		portInput:     portInput,
// 		connectButton: connectButton,
// 		statusLabel:   statusLabel,
// 	}
// }

// func (l *Lobby) Update() error {
// 	l.inputSystem.UpdateUI()

// 	mousePos := l.inputSystem.GetMousePosition()
// 	isMousePressed := l.inputSystem.IsUIActionJustPressed(input.UIActionSelect)
// 	isMouseHeld := l.inputSystem.IsUIActionPressed(input.UIActionSelect)
// 	isMouseReleased := l.inputSystem.IsUIActionJustReleased(input.UIActionSelect)

// 	// Обновление и обработка событий мыши для всех виджетов
// 	hoveredWidget := core.Widget(nil)
// 	for _, w := range l.layout.Widgets() {
// 		if w.IsVisible() {
// 			w.HandleMouseEvent(mousePos, isMouseHeld, isMouseReleased)

// 			if isMousePressed && w.Contains(mousePos) && w.IsFocusable() {
// 				if l.focusedWidget != nil && l.focusedWidget != w {
// 					l.focusedWidget.SetFocused(false)
// 				}
// 				l.focusedWidget = w
// 				l.focusedWidget.SetFocused(true)
// 			}
// 			if w.Contains(mousePos) {
// 				hoveredWidget = w
// 			}
// 		}
// 	}

// 	if isMousePressed && hoveredWidget == nil && l.focusedWidget != nil {
// 		l.focusedWidget.SetFocused(false)
// 		l.focusedWidget = nil
// 	}

// 	// Обработка клавиатуры для сфокусированного виджета
// 	if l.focusedWidget != nil {
// 		keys := ebiten.AppendPressedKeys([]ebiten.Key{})
// 		inputChars := ebiten.InputChars()

// 		for _, key := range keys {
// 			if ebiten.IsKeyJustPressed(key) {
// 				l.focusedWidget.HandleKeyboardEvent(key, 0)
// 			}
// 		}
// 		for _, r := range inputChars {
// 			l.focusedWidget.HandleKeyboardEvent(ebiten.Key(0), r)
// 		}
// 	}

// 	// Обновление внутренних состояний виджетов (например, мигающий курсор в TextInput)
// 	for _, w := range l.layout.Widgets() {
// 		if w.IsVisible() {
// 			w.Update()
// 		}
// 	}

// 	return nil
// }

// func (l *Lobby) Draw(screen *ebiten.Image) {
// 	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
// 	l.layout.ComputeBounds(w, h)

// 	for _, w := range l.layout.Widgets() {
// 		if w.IsVisible() {
// 			w.Draw(screen)
// 		}
// 	}
// }
