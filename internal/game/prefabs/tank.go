package prefabs

import (
	"fmt"
	"math"

	"github.com/Crusazer/tanks-race/internal/graphics/assets"
	"github.com/Crusazer/tanks-race/internal/graphics/renderer"
	"github.com/Crusazer/tanks-race/internal/input"
	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tank struct {
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
	tank := &Tank{
		Position: position,
		Rotation: 0,
		Velocity: m.Vector2{X: 0, Y: 0},
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

func (tank *Tank) Move(dt float64) {
	input := input.GetInput()

	if input.Up {
		tank.Position.Y -= 100 * dt
	}
	if input.Down {
		tank.Position.Y += 100 * dt
	}
	if input.Left {
		tank.Position.X -= 100 * dt
	}
	if input.Right {
		tank.Position.X += 100 * dt
	}

	tank.body.Position = tank.Position
	tank.turret.Position = tank.Position.Add(tank.turretOffset)
}

func (t *Tank) UpdateAiming() {
	mx, my := ebiten.CursorPosition()
	dirX := float64(mx) - t.turret.Position.X
	dirY := float64(my) - t.turret.Position.Y
	angle := math.Atan2(dirY, dirX)

	t.turret.Rotation = angle
}

func (tank *Tank) Draw(screen *ebiten.Image) {
	tank.body.Draw(screen)
	tank.turret.Draw(screen)
}
