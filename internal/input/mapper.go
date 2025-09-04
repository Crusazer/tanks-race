package input

import "github.com/hajimehoshi/ebiten/v2"

// GameMapper для игровых действий
type GameMapper struct {
	bindings         map[ebiten.Key]GameAction
	mouseBindings    map[ebiten.MouseButton]GameAction
	actionStates     map[GameAction]bool
	prevActionStates map[GameAction]bool
}

func NewGameMapper() *GameMapper {
	m := &GameMapper{
		bindings:         make(map[ebiten.Key]GameAction),
		mouseBindings:    make(map[ebiten.MouseButton]GameAction),
		actionStates:     make(map[GameAction]bool),
		prevActionStates: make(map[GameAction]bool),
	}

	// Привязки для WASD
	m.bindings[ebiten.KeyW] = ActionMoveUp
	m.bindings[ebiten.KeyS] = ActionMoveDown
	m.bindings[ebiten.KeyA] = ActionMoveLeft
	m.bindings[ebiten.KeyD] = ActionMoveRight

	// Привязки для стрелок
	m.bindings[ebiten.KeyArrowUp] = ActionMoveUp
	m.bindings[ebiten.KeyArrowDown] = ActionMoveDown
	m.bindings[ebiten.KeyArrowLeft] = ActionMoveLeft
	m.bindings[ebiten.KeyArrowRight] = ActionMoveRight

	// Мышь
	m.mouseBindings[ebiten.MouseButtonLeft] = ActionShoot

	return m
}

func (m *GameMapper) Update() {
	// Сохраняем предыдущее состояние
	for action := range m.actionStates {
		m.prevActionStates[action] = m.actionStates[action]
	}

	// Сброс текущего состояния
	for action := range m.actionStates {
		m.actionStates[action] = false
	}

	// Проверка клавиш
	for key, action := range m.bindings {
		if ebiten.IsKeyPressed(key) {
			m.actionStates[action] = true
		}
	}

	// Проверка мыши
	for button, action := range m.mouseBindings {
		if ebiten.IsMouseButtonPressed(button) {
			m.actionStates[action] = true
		}
	}
}

func (m *GameMapper) IsActionPressed(action GameAction) bool {
	return m.actionStates[action]
}

func (m *GameMapper) IsActionJustPressed(action GameAction) bool {
	return m.actionStates[action] && !m.prevActionStates[action]
}

func (m *GameMapper) IsActionJustReleased(action GameAction) bool {
	return !m.actionStates[action] && m.prevActionStates[action]
}

// UIMapper для UI-действий
type UIMapper struct {
	bindings         map[ebiten.Key]UIAction
	mouseBindings    map[ebiten.MouseButton]UIAction
	actionStates     map[UIAction]bool
	prevActionStates map[UIAction]bool
	inputBuffer      []rune
}

func NewUIMapper() *UIMapper {
	m := &UIMapper{
		bindings:         make(map[ebiten.Key]UIAction),
		mouseBindings:    make(map[ebiten.MouseButton]UIAction),
		actionStates:     make(map[UIAction]bool),
		prevActionStates: make(map[UIAction]bool),
	}

	// Привязки для навигации и действий
	m.bindings[ebiten.KeyArrowUp] = UIActionNavigateUp
	m.bindings[ebiten.KeyArrowDown] = UIActionNavigateDown
	m.bindings[ebiten.KeyTab] = UIActionNavigateDown
	m.bindings[ebiten.KeyEscape] = UIActionBack
	m.bindings[ebiten.KeyEnter] = UIActionConfirm
	m.bindings[ebiten.KeyBackspace] = UIActionBackspace

	// Мышь
	m.mouseBindings[ebiten.MouseButtonLeft] = UIActionSelect

	return m
}


func (m *UIMapper) Update() {
	// Сохраняем предыдущее состояние
	for action := range m.actionStates {
		m.prevActionStates[action] = m.actionStates[action]
	}

	// Сброс текущего состояния
	for action := range m.actionStates {
		m.actionStates[action] = false
	}

	// Проверка клавиш
	for key, action := range m.bindings {
		// Для Tab+Shift нужна будет отдельная логика
		if ebiten.IsKeyPressed(key) {
			m.actionStates[action] = true
		}
	}

	// Проверка мыши
	for button, action := range m.mouseBindings {
		if ebiten.IsMouseButtonPressed(button) {
			m.actionStates[action] = true
		}
	}

	// Сбор символов
	m.inputBuffer = ebiten.AppendInputChars(m.inputBuffer[:0])
}

func (m *UIMapper) IsActionPressed(action UIAction) bool {
	return m.actionStates[action]
}

func (m *UIMapper) IsActionJustPressed(action UIAction) bool {
	return m.actionStates[action] && !m.prevActionStates[action]
}

func (m *UIMapper) IsActionJustReleased(action UIAction) bool {
	return !m.actionStates[action] && m.prevActionStates[action]
}

func (m *UIMapper) GetInputChars() []rune {
	return m.inputBuffer
}