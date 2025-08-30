package shapes

type ShapeType int

const (
	RectangleShape ShapeType = iota
	CircleShape
)

type Shape interface {
	Type() ShapeType
	Bounds() AABB
}