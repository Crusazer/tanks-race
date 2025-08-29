package race

import (
	"math"

	"github.com/Crusazer/tanks-race/internal/input"
	"github.com/Crusazer/tanks-race/internal/physics/dynamics"
	m "github.com/Crusazer/tanks-race/pkg/math"
)

const (
	accelMag  = 1200.0 // пиксель/с²
	turnAcc   = 22.0    // рад/с²
	linDrag   = 3.0    // 1/с
	angDrag   = 8.0   // 1/с
)

func Drive(b *dynamics.Body, in input.InputState, dt float64) {
	dir := m.Vector2{X: math.Cos(b.Rotation), Y: math.Sin(b.Rotation)}

	// линейное ускорение
	if in.Up {
		b.AddForce(dir.Mul(accelMag * dt))
	}
	if in.Down {
		b.AddForce(dir.Mul(-accelMag * dt))
	}

	// угловое ускорение
	if in.Left {
		b.AddTorque(-turnAcc * dt)
	}
	if in.Right {
		b.AddTorque(turnAcc * dt)
	}

	// экспоненциальное трение
	exp := math.Exp(-linDrag * dt)
	b.Velocity = b.Velocity.Mul(exp)
	b.AngVel *= math.Exp(-angDrag * dt)

	// полная остановка
	if b.Velocity.Len() < 2 { b.Velocity = m.Vector2{} }
	if math.Abs(b.AngVel) < 0.02 { b.AngVel = 0 }

	b.ClampSpeed()
}
