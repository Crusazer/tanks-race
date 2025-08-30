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
	const width = 85.0
	const height = 128.0
	const mass = 10
	center := m.Vector2{X: width / 2, Y: height / 2}

	// Body
	hullBody := &dynamics.Body{
		Position: m.Vector2{X: x, Y: y},
		Mass:     mass,
		Shape:    shapes.NewRectangle(center, width, height, 0),
		Inertia:  mass * (width*width + height*height) / 12.0,
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
		Scale:   m.Vector2{X: 1, Y: 1},
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
		Width:   56,
		Height:  128,
		OriginX: 29, // Центр вращения башни по X относительно спрайта башни
		OriginY: 100, // Центр вращение башни по Y относительно спрайта башни
		Layer:   1,
		Scale:   m.Vector2{X: 1, Y: 1},
	}

	em.SetComponent(turret, entity.PositionComponent, &entity.Position{Position: m.Vector2{
		X: hullBody.Position.X + turretSprite.OriginX, // Центр вращения относительно мыши (вектор центр вращения -> мышка)
		Y: hullBody.Position.Y + turretSprite.OriginY,
	}})
	em.SetComponent(turret, entity.SpriteComponent, turretSprite)
	em.SetComponent(turret, entity.TurretComponent, &entity.Turret{
		HullID: hull.ID,
		Angle:  0,
		Offset: m.Vector2{X: 0, Y: 23}, // Смещение крепления башни относительно корпуса в пикселях
	})

	log.Println("Tank created")
}
