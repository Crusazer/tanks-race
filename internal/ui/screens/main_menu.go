package screens

import (
	"fmt"

	"github.com/Crusazer/tanks-race/internal/input"
	"github.com/Crusazer/tanks-race/internal/ui/core"

	"github.com/hajimehoshi/ebiten/v2"
)

type MenuAction string

const (
	ActionSinglePlay  MenuAction = "single_play"
	ActionNetworkPlay MenuAction = "network_play"
	ActionEditor      MenuAction = "editor"
	ActionTankConfig  MenuAction = "tank_config"
	ActionSettings    MenuAction = "settings"
	ActionExit        MenuAction = "exit"
)

type MainMenu struct {
	layout        core.Layout // Используем интерфейс Layout
	onAction      func(MenuAction)
	inputSystem   *input.InputSystem
	focusedWidget core.Widget // Для клавиатурного фокуса
}

func NewMainMenu(onAction func(MenuAction), theme *core.Theme) *MainMenu {
	if theme == nil {
		theme = core.DefaultTheme
	}

	layout := core.NewVerticalFlowLayout(30, 15)

	singlePlayButton := core.NewButton("Одиночная игра", 240, 40, func() { onAction(ActionSinglePlay) }, theme)
	layout.Add(singlePlayButton)

	networkPlayButton := core.NewButton("Игра по сети", 240, 40, func() { onAction(ActionNetworkPlay); fmt.Println("Нажата: Игра по сети") }, theme)
	layout.Add(networkPlayButton)

	editorButton := core.NewButton("Редактор уровней", 240, 40, func() { onAction(ActionEditor); fmt.Println("Нажата: Редактор уровней") }, theme)
	layout.Add(editorButton)

	tankConfigButton := core.NewButton("Настройка танка", 240, 40, func() { onAction(ActionTankConfig); fmt.Println("Нажата: Настройка танка") }, theme)
	layout.Add(tankConfigButton)

	settingsButton := core.NewButton("Настройки", 240, 40, func() { onAction(ActionSettings); fmt.Println("Нажата: Настройки") }, theme)
	layout.Add(settingsButton)

	exitButton := core.NewButton("Выход", 240, 40, func() { onAction(ActionExit) }, theme)
	layout.Add(exitButton)

	return &MainMenu{
		layout:      layout,
		onAction:    onAction,
		inputSystem: input.NewInputSystem(),
	}
}

func (m *MainMenu) Update() error {
	m.inputSystem.UpdateUI()

	// Обработка мыши
	mousePos := m.inputSystem.GetMousePosition()
	isMousePressed := m.inputSystem.IsUIActionJustPressed(input.UIActionSelect)
	isMouseHeld := m.inputSystem.IsUIActionPressed(input.UIActionSelect)
	isMouseReleased := m.inputSystem.IsUIActionJustReleased(input.UIActionSelect)
	hoveredWidget := core.Widget(nil)
	for _, w := range m.layout.Widgets() {
		if w.IsVisible() {
			w.HandleMouseEvent(mousePos, isMouseHeld, isMouseReleased)
			if w.Contains(mousePos) {
				hoveredWidget = w
			}
			if isMousePressed && w.Contains(mousePos) && w.IsFocusable() {
				m.setFocus(w)
			}
		}
	}
	if isMousePressed && hoveredWidget == nil {
		m.setFocus(nil)
	}

	// Обработка клавиатуры
	// 1. Навигация по фокусу
	if m.inputSystem.IsUIActionJustPressed(input.UIActionNavigateUp) {
		m.focusPreviousWidget()
	}
	if m.inputSystem.IsUIActionJustPressed(input.UIActionNavigateDown) {
		m.focusNextWidget()
	}

	// Если фокуса нет, но была попытка навигации, установим фокус на первый элемент
	if m.focusedWidget == nil &&
		(m.inputSystem.IsUIActionJustPressed(input.UIActionNavigateUp) ||
			m.inputSystem.IsUIActionJustPressed(input.UIActionNavigateDown) ||
			m.inputSystem.IsUIActionJustPressed(input.UIActionConfirm)) {
		m.focusFirstWidget()
	}

	// 2. Передача действий и символов сфокусированному виджету
	if m.focusedWidget != nil {
		// Передаем дискретные действия
		if m.inputSystem.IsUIActionJustPressed(input.UIActionConfirm) {
			m.focusedWidget.HandleAction(input.UIActionConfirm)
		}
		if m.inputSystem.IsUIActionJustPressed(input.UIActionBackspace) {
			m.focusedWidget.HandleAction(input.UIActionBackspace)
		}

		// Передаем вводимые символы
		chars := m.inputSystem.GetInputChars()
		if len(chars) > 0 {
			m.focusedWidget.HandleChars(chars)
		}
	}

	// Обновление состояний виджетов
	for _, w := range m.layout.Widgets() {
		if w.IsVisible() {
			w.Update()
		}
	}

	return nil
}

func (m *MainMenu) setFocus(w core.Widget) {
	if m.focusedWidget == w {
		return
	}
	if m.focusedWidget != nil {
		m.focusedWidget.SetFocused(false)
	}
	m.focusedWidget = w
	if m.focusedWidget != nil {
		m.focusedWidget.SetFocused(true)
	}
}

func (m *MainMenu) focusNextWidget() {
	focusable := m.getFocusableWidgets()
	if len(focusable) == 0 {
		return
	}

	var currentIndex int = -1
	for i, w := range focusable {
		if w == m.focusedWidget {
			currentIndex = i
			break
		}
	}

	nextIndex := 0
	if currentIndex != -1 {
		nextIndex = (currentIndex + 1) % len(focusable)
	}

	m.setFocus(focusable[nextIndex])
}

func (m *MainMenu) focusPreviousWidget() {
	focusable := m.getFocusableWidgets()
	if len(focusable) == 0 {
		return
	}

	var currentIndex int = -1
	for i, w := range focusable {
		if w == m.focusedWidget {
			currentIndex = i
			break
		}
	}

	prevIndex := len(focusable) - 1
	if currentIndex != -1 {
		prevIndex = (currentIndex - 1 + len(focusable)) % len(focusable)
	}

	m.setFocus(focusable[prevIndex])
}

func (m *MainMenu) focusFirstWidget() {
	focusable := m.getFocusableWidgets()
	if len(focusable) > 0 {
		m.setFocus(focusable[0])
	}
}

func (m *MainMenu) getFocusableWidgets() []core.Widget {
	var result []core.Widget
	for _, w := range m.layout.Widgets() {
		if w.IsVisible() && w.IsFocusable() {
			result = append(result, w)
		}
	}
	return result
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	m.layout.ComputeBounds(w, h) // Пересчитываем границы перед каждой отрисовкой (может быть оптимизировано)

	for _, w := range m.layout.Widgets() {
		if w.IsVisible() {
			w.Draw(screen)
		}
	}
}
