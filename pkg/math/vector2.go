package math

type Vector2 struct {
	X, Y float64
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{v.X + other.X, v.Y + other.Y}
}

func (v Vector2) Mul(scalar float64) Vector2{
	return Vector2{v.X * scalar, v.Y * scalar}
}

