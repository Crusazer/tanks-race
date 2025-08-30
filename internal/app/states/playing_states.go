package states

import (
	"github.com/Crusazer/tanks-race/internal/game/entity"
	"github.com/Crusazer/tanks-race/internal/graphics/renderer"
	"github.com/Crusazer/tanks-race/internal/physics"
	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"math"
)

type PlayingRunningState struct {
	world  *physics.World
	em     *entity.Manager
	camera *renderer.Camera
}

func NewPlayingRunningState(w *physics.World, em *entity.Manager, cam *renderer.Camera) *PlayingRunningState {
	return &PlayingRunningState{world: w, em: em, camera: cam}
}
func (s *PlayingRunningState) Enter() {}
func (s *PlayingRunningState) Exit()  {}

func (s *PlayingRunningState) Update(dt float64) error {
	s.handleAllTanksInput()
	s.world.Update(dt)
	s.updateBounded()
	s.updateCamera()
	return nil
}

func (s *PlayingRunningState) updateBounded() {
	for _, e := range s.em.GetWithComponents(entity.TurretComponent) {
		position, ok := e.Components[entity.PositionComponent].(*entity.Position)
		if !ok {
			log.Fatal("Turret position not found")
			return
		}

		turret, ok := e.Components[entity.TurretComponent].(*entity.Turret)
		if !ok {
			log.Fatal("Turret has no hull")
			return
		}

		hull := s.em.Get(turret.HullID)
		hullPhysic := hull.Components[entity.PhysicsComponent].(*entity.Physics)

		// смещение с учётом угла корпуса
		rotated := turret.Offset.Rotate(hullPhysic.Body.Rotation)
		position.Position.X = hullPhysic.Body.Position.X + rotated.X
		position.Position.Y = hullPhysic.Body.Position.Y + rotated.Y

		// угол башни → спрайт
		if spr, ok := e.Components[entity.SpriteComponent].(*entity.Sprite); ok {
			spr.Rotation = turret.Angle
		}
	}
}

func (s *PlayingRunningState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{200, 200, 200, 255})

	for _, layer := range s.em.RenderQueue.Order {
		for _, e := range s.em.RenderQueue.Layers[layer] {
			s.drawEntity(screen, e)
		}
	}
}

func (s *PlayingRunningState) handleAllTanksInput() {
	const force = 400

	for _, e := range s.em.GetWithComponents(entity.VehicleComponent) {
		physic, ok := e.Components[entity.PhysicsComponent].(*entity.Physics)
		if !ok {
			continue
		}

		body := physic.Body
		angle := body.Rotation

		if ebiten.IsKeyPressed(ebiten.KeyW) {
			body.Force = body.Force.Add(m.Vector2{X: 0, Y: -1}.Rotate(angle).Scale(force))
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			body.Force = body.Force.Add(m.Vector2{X: 0, Y: 1}.Rotate(angle).Scale(force))
		}
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			body.Force = body.Force.Add(m.Vector2{X: -1, Y: 0}.Rotate(angle).Scale(force))
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			body.Force = body.Force.Add(m.Vector2{X: 1, Y: 0}.Rotate(angle).Scale(force))
		}
	}

	// Поворот башни за мышью
	turret_entity := s.em.GetWithComponents(entity.TurretComponent)[0]
	turret := turret_entity.Components[entity.TurretComponent].(*entity.Turret)
	turretSprite := turret_entity.Components[entity.SpriteComponent].(*entity.Sprite)
	turretPos := turret_entity.Components[entity.PositionComponent].(*entity.Position)

	mouseX, mouseY := ebiten.CursorPosition()
	mouseWorldX := float64(mouseX)/s.camera.Zoom + s.camera.Position.X - s.camera.ViewportWidth/2
	mouseWorldY := float64(mouseY)/s.camera.Zoom + s.camera.Position.Y - s.camera.ViewportHeight/2
	dx := mouseWorldX - turretPos.Position.X
	dy := mouseWorldY - turretPos.Position.Y
	angle := math.Atan2(dy, dx)
	turret.Angle = angle
	turretSprite.Rotation = angle

}

func (s *PlayingRunningState) updateCamera() {
	for _, e := range s.em.GetWithComponents(entity.VehicleComponent) {
		if comp, ok := e.Components[entity.PhysicsComponent].(*entity.Physics); ok {
			target := comp.Body.Position
			alpha := 0.12
			s.camera.Position.X += (target.X - s.camera.Position.X) * alpha
			s.camera.Position.Y += (target.Y - s.camera.Position.Y) * alpha
			return
		}
	}
}

func (s *PlayingRunningState) drawEntity(screen *ebiten.Image, e *entity.Entity) {
	var pos m.Vector2
	var rotation float64

	sprite, ok := e.Components[entity.SpriteComponent].(*entity.Sprite)
	if !ok {
		return
	}

	if comp, ok := e.Components[entity.PhysicsComponent].(*entity.Physics); ok {
		pos = comp.Body.Position
		rotation = comp.Body.Rotation
	} else if posComp, ok := e.Components[entity.PositionComponent].(*entity.Position); ok {
		pos = posComp.Position
		rotation = sprite.Rotation
	} else {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-sprite.OriginX, -sprite.OriginY) // центр текстуры → 0,0
	op.GeoM.Rotate(rotation)

	op.GeoM.Translate(pos.X, pos.Y)
	op.GeoM.Translate(-s.camera.Position.X, -s.camera.Position.Y)
	op.GeoM.Scale(s.camera.Zoom, s.camera.Zoom)
	op.GeoM.Translate(s.camera.ViewportWidth/2, s.camera.ViewportHeight/2)

	screen.DrawImage(sprite.Image, op)
}
