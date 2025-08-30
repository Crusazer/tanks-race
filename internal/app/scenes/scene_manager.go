package scenes

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update(dt float64) error
	Draw(screen *ebiten.Image)
	Layout(w, h int) (int, int)
	Enter()
	Exit()
}

type SceneManager struct {
	current Scene
}

func NewSceneManager(initial Scene) *SceneManager {
	sm := &SceneManager{}
	sm.ChangeScene(initial)
	return sm
}

func (sm *SceneManager) ChangeScene(s Scene) {
	if sm.current != nil {
		sm.current.Exit()
	}
	sm.current = s
	s.Enter()
}

func (sm *SceneManager) Update(dt float64) error {
	return sm.current.Update(dt)
}

func (sm *SceneManager) Draw(screen *ebiten.Image) {
	sm.current.Draw(screen)
}

func (sm *SceneManager) Layout(w, h int) (int, int) {
	return sm.current.Layout(w, h)
}
