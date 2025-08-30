package states

import "github.com/hajimehoshi/ebiten/v2"

type GameState interface {
	Update(dt float64) error
	Draw(screen *ebiten.Image)
	Enter()
	Exit()
}

type StateMachine struct {
	current GameState
}

func NewStateMachine(initial GameState) *StateMachine {
	sm := &StateMachine{current: initial}
	initial.Enter()
	return sm
}

func (sm *StateMachine) Change(s GameState) {
	if sm.current != nil {
		sm.current.Exit()
	}
	sm.current = s
	s.Enter()
}

func (sm *StateMachine) Update(dt float64) error {
	return sm.current.Update(dt)
}

func (sm *StateMachine) Draw(screen *ebiten.Image) {
	sm.current.Draw(screen)
}