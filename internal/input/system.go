// internal/input/system.go
package input

import (
	"github.com/Crusazer/tanks-race/pkg/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type InputSystem struct {
	gameMapper *GameMapper
	uiMapper   *UIMapper
}

func NewInputSystem() *InputSystem {
	return &InputSystem{
		gameMapper: NewGameMapper(),
		uiMapper:   NewUIMapper(),
	}
}

func (is *InputSystem) Update() {
	is.gameMapper.Update()
	is.uiMapper.Update()
}

func (is *InputSystem) UpdateGame() {
	is.gameMapper.Update()
}

func (is *InputSystem) UpdateUI() {
	is.uiMapper.Update()
}

// Для игровых действий
func (is *InputSystem) IsGameActionPressed(action GameAction) bool {
	return is.gameMapper.IsActionPressed(action)
}

func (is *InputSystem) IsGameActionJustPressed(action GameAction) bool {
	return is.gameMapper.IsActionJustPressed(action)
}

func (is *InputSystem) IsGameActionJustReleased(action GameAction) bool {
	return is.gameMapper.IsActionJustReleased(action)
}

// Для UI-действий
func (is *InputSystem) IsUIActionPressed(action UIAction) bool {
	return is.uiMapper.IsActionPressed(action)
}

func (is *InputSystem) IsUIActionJustPressed(action UIAction) bool {
	return is.uiMapper.IsActionJustPressed(action)
}

func (is *InputSystem) GetMousePosition() math.Vector2 {
	x, y := ebiten.CursorPosition()
	return math.Vector2{X: float64(x), Y: float64(y)}
}

func (is *InputSystem) IsUIActionJustReleased(action UIAction) bool {
	return is.uiMapper.IsActionJustReleased(action)
}

func (is *InputSystem) GetInputChars() []rune {
	return is.uiMapper.GetInputChars()
}
