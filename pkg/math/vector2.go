package math

import "math"

type Vector2 struct {
	X, Y float64
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{v.X + other.X, v.Y + other.Y}
}

func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{v.X - other.X, v.Y - other.Y}
}

func (v Vector2) Mul(scalar float64) Vector2 {
	return Vector2{v.X * scalar, v.Y * scalar}
}

func (v Vector2) Div(scalar float64) Vector2 {
	if scalar == 0 {
		panic("division by zero")
	}
	return Vector2{v.X / scalar, v.Y / scalar}
}

func (v Vector2) Scale(s float64) Vector2 {
	return Vector2{v.X * s, v.Y * s}
}

func (v Vector2) Dot(other Vector2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vector2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2) Normalize() Vector2 {
	length := v.Length()
	if length == 0 {
		return Vector2{}
	}
	return v.Scale(1.0 / length)
}

func (v Vector2) Rotate(angle float64) Vector2 {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Vector2{
		X: v.X*cos - v.Y*sin,
		Y: v.X*sin + v.Y*cos,
	}
}
