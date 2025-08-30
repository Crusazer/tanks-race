package input

import "github.com/hajimehoshi/ebiten/v2"

type System struct {
	actions map[Action]bool
	mouseX  int
	mouseY  int
}

type Action int

const (
	ActionMoveUp Action = iota
	ActionMoveDown
	ActionMoveLeft
	ActionMoveRight
	ActionShoot
)

func NewSystem() *System{
	    return &System{
        actions: make(map[Action]bool),
    }
}

func (s *System) Update() {
    s.actions[ActionMoveUp] = ebiten.IsKeyPressed(ebiten.KeyW)
    s.actions[ActionMoveDown] = ebiten.IsKeyPressed(ebiten.KeyS)
    s.actions[ActionMoveLeft] = ebiten.IsKeyPressed(ebiten.KeyA)
    s.actions[ActionMoveRight] = ebiten.IsKeyPressed(ebiten.KeyD)
    s.actions[ActionShoot] = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
    
    s.mouseX, s.mouseY = ebiten.CursorPosition()
}

func (s *System) IsPressed(action Action) bool {
    return s.actions[action]
}

func (s *System) MousePosition() (int, int) {
    return s.mouseX, s.mouseY
}
