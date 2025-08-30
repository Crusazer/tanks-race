package dynamics

import (
	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/Crusazer/tanks-race/internal/physics/shapes"
)

type Body struct {
	Position m.Vector2
	Velocity m.Vector2
	AngVel   float64
	Force    m.Vector2
	Torque   float64

	Mass     float64
	Rotation float64
	Inertia  float64

	Shape shapes.Shape
}

func (b *Body) AddForce(force m.Vector2) {
	b.Force = b.Force.Add(force)
}

func (b *Body) AddTorque(torque float64) {
	b.Torque += torque
}

func (b *Body) ClampSpeed() {
	maxSpeed := 500.0
	maxAngVel := 10.0
	
	if b.Velocity.Length() > maxSpeed {
		b.Velocity = b.Velocity.Normalize().Scale(maxSpeed)
	}
	if b.AngVel > maxAngVel {
		b.AngVel = maxAngVel
	}
}