package shapes

import (
	m "github.com/Crusazer/tanks-race/pkg/math"
)

type Rectangle struct {
	Center     m.Vector2
	Width      float64
	Height     float64
	Rotation   float64
	HalfWidth  float64
	HalfHeight float64
}

func NewRectangle(center m.Vector2, width, height, rotation float64) *Rectangle {
	return &Rectangle{
		Center:     center,
		Width:      width,
		Height:     height,
		Rotation:   rotation,
		HalfWidth:  width / 2,
		HalfHeight: height / 2,
	}
}

func (r *Rectangle) Type() ShapeType {
	return RectangleShape
}

func (r *Rectangle) Bounds() AABB {
	corners := r.GetCorners()
	minX, minY := corners[0].X, corners[0].Y
	maxX, maxY := corners[0].X, corners[0].Y

	for _, corner := range corners[1:] {
		if corner.X < minX {
			minX = corner.X
		}
		if corner.X > maxX {
			maxX = corner.X
		}
		if corner.Y < minY {
			minY = corner.Y
		}
		if corner.Y > maxY {
			maxY = corner.Y
		}
	}

	return AABB{
		Min: m.Vector2{X: minX, Y: minY},
		Max: m.Vector2{X: maxX, Y: maxY},
	}
}

// GetAxes возвращает оси для SAT (нормали к сторонам)
func (r *Rectangle) GetAxes() []m.Vector2 {
	// Получаем углы прямоугольника
	corners := r.GetCorners()

	// Создаем оси (нормали к сторонам)
	axes := make([]m.Vector2, 2)

	// Ось 1: нормаль к верхней стороне
	edge1 := corners[1].Sub(corners[0])
	axes[0] = m.Vector2{X: -edge1.Y, Y: edge1.X}.Normalize()

	// Ось 2: нормаль к правой стороне
	edge2 := corners[2].Sub(corners[1])
	axes[1] = m.Vector2{X: -edge2.Y, Y: edge2.X}.Normalize()

	return axes
}

// GetCorners возвращает 4 угла прямоугольника
func (r *Rectangle) GetCorners() []m.Vector2 {
	corners := make([]m.Vector2, 4)

	// Локальные координаты углов относительно центра
	hw := r.HalfWidth
	hh := r.HalfHeight

	localCorners := []m.Vector2{
		{X: -hw, Y: -hh}, // Нижний-левый
		{X: hw, Y: -hh},  // Нижний-правый
		{X: hw, Y: hh},   // Верхний-правый
		{X: -hw, Y: hh},  // Верхний-левый
	}

	// Поворачиваем и сдвигаем
	for i, corner := range localCorners {
		rotated := corner.Rotate(r.Rotation)
		corners[i] = rotated.Add(r.Center)
	}

	return corners
}

// Project проектирует прямоугольник на ось
func (r *Rectangle) Project(axis m.Vector2) (min, max float64) {
	corners := r.GetCorners()

	min = corners[0].Dot(axis)
	max = min

	for i := 1; i < 4; i++ {
		projection := corners[i].Dot(axis)
		if projection < min {
			min = projection
		}
		if projection > max {
			max = projection
		}
	}

	return min, max
}