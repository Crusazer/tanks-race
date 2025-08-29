package dynamics

func Integrate(b *Body, dt float64) {
	b.Position = b.Position.Add(b.Velocity.Mul(dt))
	b.Rotation += b.AngVel * dt
}