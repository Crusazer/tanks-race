package entity

import (
	"github.com/Crusazer/tanks-race/internal/physics/dynamics"
	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type ID uint64

type Entity struct {
	ID         ID
	Components map[ComponentType]interface{}
}

type ComponentType int

const (
	PositionComponent ComponentType = iota
	PhysicsComponent
	SpriteComponent
	VehicleComponent
	TurretComponent
)

type Position struct {
	Position m.Vector2
}

type Physics struct {
	Body *dynamics.Body
}

type Sprite struct {
	Image    *ebiten.Image
	Width    float64
	Height   float64
	OriginX  float64
	OriginY  float64
	Rotation float64
	Layer    int
	Scale    m.Vector2
}

type Vehicle struct{}

type Turret struct {
	HullID ID // корпус, к которому прикреплена башня
	Angle  float64
	Offset m.Vector2 // смещение от центра корпуса
}
