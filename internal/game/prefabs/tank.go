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

func CreateTank(em *entity.Manager, world *physics.World, x, y float64) {
	// === КОРПУС ===
	const width = 128.0
	const height = 85.0
	const mass = 10
	center := m.Vector2{X: width / 2, Y: height / 2}

	// Body
	hullBody := &dynamics.Body{
		Position: m.Vector2{X: x, Y: y},
		Mass:     mass,
		Shape:    shapes.NewRectangle(center, width, height, 0),
	}
	world.AddBody(hullBody)

	// Entity
	hull := em.Create()
	hullImg, err := assets.LoadImage("tanks/default/tank_body.png")
	if err != nil {
		log.Fatal(err)
		return
	}
	hullSprite := &entity.Sprite{
		Image:   hullImg,
		Width:   width,
		Height:  height,
		OriginX: center.X,
		OriginY: center.Y,
		Layer:   0,
	}
	em.SetComponent(hull, entity.PhysicsComponent, &entity.Physics{Body: hullBody})
	em.SetComponent(hull, entity.SpriteComponent, hullSprite)
	em.SetComponent(hull, entity.VehicleComponent, entity.Vehicle{})

	// === БАШНЯ ===
	turret := em.Create()
	img, err := assets.LoadImage("tanks/default/tank_turret.png")
	if err != nil {
		log.Fatal(err)
		return
	}
	turretSprite := &entity.Sprite{
		Image:   img,
		Width:   128,
		Height:  56,
		OriginX: 50,
		OriginY: 20,
		Layer:   1,
	}

	em.SetComponent(turret, entity.PositionComponent, &entity.Position{Position: m.Vector2{X: x, Y: y}})
	em.SetComponent(turret, entity.SpriteComponent, turretSprite)
	em.SetComponent(turret, entity.TurretComponent, &entity.Turret{
		HullID: hull.ID,
		Angle:  0,
		Offset: m.Vector2{X: 0, Y: -8},
	})

	log.Println("Tank created")
}
