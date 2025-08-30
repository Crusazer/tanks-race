package states

import (
	"image/color"
	"log"
	"math"

	"github.com/Crusazer/tanks-race/internal/game/entity"
	"github.com/Crusazer/tanks-race/internal/graphics/renderer"
	"github.com/Crusazer/tanks-race/internal/physics"
	"github.com/Crusazer/tanks-race/internal/physics/shapes"
	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
		position.Position = hullPhysic.Body.Position.Add(rotated)

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

	s.drawPhysicsBodies(screen)
}

func (s *PlayingRunningState) handleAllTanksInput() {
	const force = 1000.0
	const torque = 40000.0

	for _, e := range s.em.GetWithComponents(entity.VehicleComponent) {
		physic, ok := e.Components[entity.PhysicsComponent].(*entity.Physics)
		if !ok {
			continue
		}

		body := physic.Body
		angle := body.Rotation

		if ebiten.IsKeyPressed(ebiten.KeyW) {
			forward := m.Vector2{X: 0, Y: -1}.Rotate(angle)
			body.Force = body.Force.Add(forward.Scale(force))
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			backward := m.Vector2{X: 0, Y: -1}.Rotate(angle).Scale(-1)
			body.Force = body.Force.Add(backward.Scale(force))
		}
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			body.Torque -= torque
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			body.Torque += torque
		}
	}

	// Поворот башни за мышью
	turret_entity := s.em.GetWithComponents(entity.TurretComponent)[0]
	turret := turret_entity.Components[entity.TurretComponent].(*entity.Turret)
	turretSprite := turret_entity.Components[entity.SpriteComponent].(*entity.Sprite)
	turretPos := turret_entity.Components[entity.PositionComponent].(*entity.Position)

	mouseX, mouseY := ebiten.CursorPosition()
	mouseWorldX := (float64(mouseX)-s.camera.ViewportWidth/2)/s.camera.Zoom + s.camera.Position.X
	mouseWorldY := (float64(mouseY)-s.camera.ViewportHeight/2)/s.camera.Zoom + s.camera.Position.Y

	dx := mouseWorldX - turretPos.Position.X
	dy := mouseWorldY - turretPos.Position.Y
	angle := math.Atan2(dy, dx) - -math.Pi/2
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
	op.GeoM.Scale(sprite.Scale.X, sprite.Scale.Y)
	op.GeoM.Translate(-sprite.OriginX, -sprite.OriginY) // центр текстуры → 0,0
	op.GeoM.Rotate(rotation)

	op.GeoM.Translate(pos.X, pos.Y)
	op.GeoM.Translate(-s.camera.Position.X, -s.camera.Position.Y)
	op.GeoM.Scale(s.camera.Zoom, s.camera.Zoom)
	op.GeoM.Translate(s.camera.ViewportWidth/2, s.camera.ViewportHeight/2)

	screen.DrawImage(sprite.Image, op)
}

func (s *PlayingRunningState) drawPhysicsBodies(screen *ebiten.Image) {
	for _, body := range s.em.GetWithComponents(entity.PhysicsComponent) {
		// Для прямоугольника
		physicComp := body.Components[entity.PhysicsComponent].(*entity.Physics)
		body := physicComp.Body

		if rect, ok := body.Shape.(*shapes.Rectangle); ok {
			w, h := rect.Width, rect.Height
			corners := []m.Vector2{
				{X: -w / 2, Y: -h / 2},
				{X: w / 2, Y: -h / 2},
				{X: w / 2, Y: h / 2},
				{X: -w / 2, Y: h / 2},
			}

			for i := range corners {
				corners[i] = corners[i].Rotate(body.Rotation).Add(body.Position)
				// смещение камеры
				corners[i].X = (corners[i].X-s.camera.Position.X)*s.camera.Zoom + s.camera.ViewportWidth/2
				corners[i].Y = (corners[i].Y-s.camera.Position.Y)*s.camera.Zoom + s.camera.ViewportHeight/2
			}

			// Нарисовать линии между углами
			for i := 0; i < 4; i++ {
				next := (i + 1) % 4
				vector.StrokeLine(
					screen,
					float32(corners[i].X),
					float32(corners[i].Y),
					float32(corners[next].X),
					float32(corners[next].Y),
					1,
					color.RGBA{255, 0, 0, 255},
					false,
				)
			}
		}
	}
}
