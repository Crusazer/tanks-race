package prefabs

import (
	"github.com/Crusazer/tanks-race/internal/game/entity"
	"github.com/Crusazer/tanks-race/internal/graphics/assets"
	"github.com/Crusazer/tanks-race/internal/physics"
	"github.com/Crusazer/tanks-race/internal/physics/dynamics"
	"github.com/Crusazer/tanks-race/internal/physics/shapes"
	m "github.com/Crusazer/tanks-race/pkg/math"

	"log"
)

func CreateBox(em *entity.Manager, world *physics.World, x, y float64) {
	const width = 100.0
	const height = 100.0
	const mass = 1

	center := m.Vector2{X: width / 2, Y: height / 2}

	body := &dynamics.Body{
		Position: m.Vector2{X: x, Y: y},
		Mass:     mass,
		Shape:    shapes.NewRectangle(center, width, height, 0),
		Inertia:  mass * (width*width + height*height) / 12.0,
	}
	world.AddBody(body)

	hull := em.Create()
	hullImg, err := assets.LoadImage("box.png")
	if err != nil {
		log.Fatal(err)
		return
	}
	sprite := &entity.Sprite{
		Image:   hullImg,
		Width:   width,
		Height:  height,
		OriginX: center.X,
		OriginY: center.Y,
		Layer:   0,
		Scale:   m.Vector2{X: 10, Y: 10},
	}
	em.SetComponent(hull, entity.PhysicsComponent, &entity.Physics{Body: body})
	em.SetComponent(hull, entity.SpriteComponent, sprite)
}
