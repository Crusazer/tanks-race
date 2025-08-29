package prefabs

import (
	"fmt"
	"github.com/Crusazer/tanks-race/internal/graphics/assets"
	"github.com/Crusazer/tanks-race/internal/graphics/renderer"
	"github.com/Crusazer/tanks-race/internal/input"
	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tank struct {
	Position m.Vector2
	Rotation float64
	Velocity m.Vector2

	body   *renderer.Sprite
	turret *renderer.Sprite
}

func (tank *Tank) Init() error {
	body, err := assets.LoadTexture("tanks/default/tank_body.png")
	if err != nil {
		return fmt.Errorf("failed to load body: %w", err)
	}

	turret, err := assets.LoadTexture("tanks/default/tank_turret.png")
	if err != nil {
		return fmt.Errorf("failed to load turret: %w", err)
	}

	tank.body = &renderer.Sprite{
		Image:    body,
		Position: tank.Position,
		Rotation: 0,
		Scale:    m.Vector2{X: 1, Y: 1},
		Origin:   m.Vector2{X: float64(body.Bounds().Dx()) / 2, Y: float64(body.Bounds().Dy()) / 2},
	}
	tank.turret = &renderer.Sprite{
		Image:    turret,
		Position: tank.Position,
		Rotation: 0,
		Scale:    m.Vector2{X: 1, Y: 1},
		Origin:   m.Vector2{X: float64(body.Bounds().Dx()) / 2, Y: float64(body.Bounds().Dy()) / 2},
	}
	return nil
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
	tank.turret.Position = tank.Position
}

func (tank *Tank) Draw(screen *ebiten.Image) {
	tank.body.Draw(screen)
	tank.turret.Draw(screen)
}
