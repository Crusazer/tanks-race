package scenes

import (
	"github.com/Crusazer/tanks-race/internal/app/states"
	"github.com/Crusazer/tanks-race/internal/game/entity"
	"github.com/Crusazer/tanks-race/internal/game/prefabs"
	"github.com/Crusazer/tanks-race/internal/graphics/renderer"
	"github.com/Crusazer/tanks-race/internal/physics"
	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayingScene struct {
	world  *physics.World
	em     *entity.Manager
	camera *renderer.Camera
	sm     *states.StateMachine
}

func NewPlayingScene() Scene {
	w := physics.NewWorld()
	em := entity.NewManager()
	cam := &renderer.Camera{Position: m.Vector2{X: 0, Y: 0}, Zoom: 0.6}

	prefabs.CreateTank(em, w, 400, 300)

	return &PlayingScene{
		world:  w,
		em:     em,
		camera: cam,
		sm: states.NewStateMachine(
			states.NewPlayingRunningState(w, em, cam),
		),
	}
}

func (ps *PlayingScene) Enter() {}
func (ps *PlayingScene) Exit() {
	ps.world = nil
	ps.em = nil
}
func (ps *PlayingScene) Update(dt float64) error {
	return ps.sm.Update(dt)
}
func (ps *PlayingScene) Draw(screen *ebiten.Image) {
	ps.sm.Draw(screen)
}
func (ps *PlayingScene) Layout(w, h int) (int, int) { return w, h }
