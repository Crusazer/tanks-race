package prefabs

import (
	"fmt"
	"math"

	"github.com/Crusazer/tanks-race/internal/game/systems/race"
	"github.com/Crusazer/tanks-race/internal/graphics/assets"
	"github.com/Crusazer/tanks-race/internal/graphics/renderer"
	"github.com/Crusazer/tanks-race/internal/input"
	"github.com/Crusazer/tanks-race/internal/physics/dynamics"
	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tank struct {
	Body         *dynamics.Body
	Position     m.Vector2
	Rotation     float64
	Velocity     m.Vector2
	turretOffset m.Vector2
	body         *renderer.Sprite
	turret       *renderer.Sprite
}

func NewTank() (*Tank, error) {
	bodyImg, err := assets.LoadTexture("tanks/default/tank_body.png")
	if err != nil {
		return nil, fmt.Errorf("failed to load body: %w", err)
	}

	turretImg, err := assets.LoadTexture("tanks/default/tank_turret.png")
	if err != nil {
		return nil, fmt.Errorf("failed to load turret: %w", err)
	}

	position := m.Vector2{X: 110, Y: 110}
	body := &dynamics.Body{
		Position: position,
		MaxSpeed: 320.0,
		Mass:     1.0,
		Inertia:  0.8,
	}

	tank := &Tank{
		Body:         body,
		Position:     position,
		Rotation:     0,
		Velocity:     m.Vector2{X: 0, Y: 0},
		turretOffset: m.Vector2{X: -23, Y: 0},
		body: &renderer.Sprite{
			Image:    bodyImg,
			Position: position,
			Rotation: 0,
			Scale:    m.Vector2{X: 1, Y: 1},
			Origin:   m.Vector2{X: float64(bodyImg.Bounds().Dx()) / 2, Y: float64(bodyImg.Bounds().Dy()) / 2},
		},
		turret: &renderer.Sprite{
			Image:    turretImg,
			Position: position,
			Rotation: 0,
			Scale:    m.Vector2{X: 1, Y: 1},
			Origin:   m.Vector2{X: 34, Y: 28},
		},
	}
	return tank, nil
}

func (t *Tank) Move(dt float64) {
	input := input.GetInput()
	race.Drive(t.Body, input, dt)
	dynamics.Integrate(t.Body, dt)

	// корпус
	t.body.Position = t.Body.Position
	t.body.Rotation = t.Body.Rotation

	// башня
	local := t.turretOffset
	cos, sin := math.Cos(t.Body.Rotation), math.Sin(t.Body.Rotation)
	world := m.Vector2{
		X: local.X*cos - local.Y*sin,
		Y: local.X*sin + local.Y*cos,
	}
	t.turret.Position = t.Body.Position.Add(world)
}

func (t *Tank) UpdateAiming(cam *renderer.Camera) {
	mx, my := ebiten.CursorPosition()
	mouseWorld := cam.ScreenToWorld(m.Vector2{X: float64(mx), Y: float64(my)})
	dirX := mouseWorld.X - t.turret.Position.X
	dirY := mouseWorld.Y - t.turret.Position.Y
	t.turret.Rotation = math.Atan2(dirY, dirX)
}

func (tank *Tank) Draw(screen *ebiten.Image, cam *renderer.Camera) {
	tank.body.Draw(screen, cam)
	tank.turret.Draw(screen, cam)
}
