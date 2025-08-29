package input

import "github.com/hajimehoshi/ebiten/v2"

type InputState struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

func GetInput() InputState {
	return InputState{
		Up:    ebiten.IsKeyPressed(ebiten.KeyW),
		Down:  ebiten.IsKeyPressed(ebiten.KeyS),
		Left:  ebiten.IsKeyPressed(ebiten.KeyA),
		Right: ebiten.IsKeyPressed(ebiten.KeyD),
	}
}
