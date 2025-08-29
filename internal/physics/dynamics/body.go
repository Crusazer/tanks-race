package dynamics

import m "github.com/Crusazer/tanks-race/pkg/math"

type Body struct {
	Position m.Vector2
	Velocity m.Vector2
	Rotation float64 // radians
	AngVel   float64 // rad/sec

	Mass    float64 // kg
	Inertia float64

	MaxSpeed float64
}

func (b *Body) AddForce(f m.Vector2) {
	a := f.Div(b.Mass)
	b.Velocity = b.Velocity.Add(a)
}

func (b *Body) AddTorque(t float64) {
	alpha := t / b.Inertia
	b.AngVel += alpha
}

func (b *Body) ClampSpeed() {
	if l := b.Velocity.Len(); l > b.MaxSpeed{
		b.Velocity = b.Velocity.Mul(b.MaxSpeed / l)
	}
}

