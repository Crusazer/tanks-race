package physics

import (
	"github.com/Crusazer/tanks-race/internal/physics/collision"
	"github.com/Crusazer/tanks-race/internal/physics/dynamics"
	"github.com/Crusazer/tanks-race/internal/physics/shapes"
	m "github.com/Crusazer/tanks-race/pkg/math"

	"math"
)

type World struct {
	bodies     []*dynamics.Body
	rectangles []*shapes.Rectangle
	detector   *collision.Detector
}

func NewWorld() *World {
	return &World{
		bodies:     make([]*dynamics.Body, 0),
		rectangles: make([]*shapes.Rectangle, 0),
		detector:   collision.NewDetector(),
	}
}

func (w *World) AddBody(body *dynamics.Body) {
	w.bodies = append(w.bodies, body)

	// Создаем прямоугольник для столкновений
	if rectShape, ok := body.Shape.(*shapes.Rectangle); ok {
		rect := shapes.NewRectangle(
			body.Position,
			rectShape.Width,
			rectShape.Height,
			body.Rotation,
		)
		w.rectangles = append(w.rectangles, rect)
	}
}

func (w *World) GetBodies() []*dynamics.Body {
	return w.bodies
}

func (w *World) integrate(dt float64) {
	const linearDamping = 2.0  // коэффициент линейного трения (чем выше, тем быстрее затухает)
	const angularDamping = 2.5 // коэффициент углового трения

	for _, body := range w.bodies {
		if body.Mass <= 0 {
			continue // статическое тело
		}

		// --- Линейная динамика ---
		acc := body.Force.Scale(1 / body.Mass)
		body.Velocity = body.Velocity.Add(acc.Scale(dt))
		body.Position = body.Position.Add(body.Velocity.Scale(dt))

		// Линейное трение (экспоненциальное затухание)
		if body.Velocity.Length() > 0 {
			damp := math.Exp(-linearDamping * dt)
			body.Velocity = body.Velocity.Scale(damp)

			if body.Velocity.Length() < 0.001 {
				body.Velocity = m.Vector2{} // обнуляем мелкие значения
			}
		}

		// --- Угловая динамика ---
		angAcc := body.Torque / body.Inertia
		body.AngVel += angAcc * dt
		body.Rotation += body.AngVel * dt

		// Угловое трение (экспоненциальное затухание)
		if math.Abs(body.AngVel) > 0 {
			damp := math.Exp(-angularDamping * dt)
			body.AngVel *= damp

			if math.Abs(body.AngVel) < 0.001 {
				body.AngVel = 0
			}
		}

		// Сброс сил
		body.Force = m.Vector2{}
		body.Torque = 0
	}
}


func (w *World) Update(dt float64) {
	// Интеграция движения
	w.integrate(dt)

	// Обновляем позиции прямоугольников
	for i, body := range w.bodies {
		if i < len(w.rectangles) {
			w.rectangles[i].Center = body.Position
			w.rectangles[i].Rotation = body.Rotation
		}
	}

	// Проверка столкновений
	collisions := w.detector.Detect(w.rectangles, w.bodies)

	// Решение столкновений
	w.resolveCollisions(collisions)
}

func (w *World) resolveCollisions(collisions []collision.CollisionResult) {
	for _, col := range collisions {
		w.resolveCollision(col)
	}
}

func (w *World) resolveCollision(col collision.CollisionResult) {
	bodyA := col.BodyA
	bodyB := col.BodyB

	// Пропускаем, если оба тела статические
	if bodyA.Mass <= 0 && bodyB.Mass <= 0 {
		return
	}

	// Раздвигаем тела
	separation := col.SATResult.Normal.Scale(col.SATResult.Overlap)

	if bodyA.Mass <= 0 {
		// bodyA статический
		bodyB.Position = bodyB.Position.Add(separation)
	} else if bodyB.Mass <= 0 {
		// bodyB статический
		bodyA.Position = bodyA.Position.Sub(separation)
	} else {
		// Оба движутся
		totalMass := bodyA.Mass + bodyB.Mass
		ratioA := bodyB.Mass / totalMass
		ratioB := bodyA.Mass / totalMass

		bodyA.Position = bodyA.Position.Sub(separation.Scale(ratioA))
		bodyB.Position = bodyB.Position.Add(separation.Scale(ratioB))
	}

	// Простая реакция - отражение скорости
	relativeVel := bodyB.Velocity.Sub(bodyA.Velocity)
	separatingSpeed := relativeVel.Dot(col.SATResult.Normal)

	if separatingSpeed > 0 {
		return // Объекты уже разлетаются
	}

	// Коэффициент упругости (можно сделать настраиваемым)
	restitution := 0.8

	// Вычисляем импульс
	impulse := -(1 + restitution) * separatingSpeed
	if bodyA.Mass > 0 && bodyB.Mass > 0 {
		impulse /= (1/bodyA.Mass + 1/bodyB.Mass)
	} else if bodyA.Mass > 0 {
		impulse /= (1 / bodyA.Mass)
	} else {
		impulse /= (1 / bodyB.Mass)
	}

	impulseVec := col.SATResult.Normal.Scale(impulse)

	// Применяем импульс
	if bodyA.Mass > 0 {
		bodyA.Velocity = bodyA.Velocity.Sub(impulseVec.Scale(1 / bodyA.Mass))
	}
	if bodyB.Mass > 0 {
		bodyB.Velocity = bodyB.Velocity.Add(impulseVec.Scale(1 / bodyB.Mass))
	}
}
