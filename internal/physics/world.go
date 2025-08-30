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
		if math.Abs(body.AngVel) > 0.0001 {
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
	if bodyA == nil || bodyB == nil || col.SATResult == nil {
		return
	}

	// Нормаль столкновения (нормализуем на всякий случай)
	n := col.SATResult.Normal
	nLen := math.Hypot(n.X, n.Y)
	if nLen == 0 {
		return
	}
	n = m.Vector2{X: n.X / nLen, Y: n.Y / nLen}

	// CONTACT POINT (используем Contact1, если он задан)
	contact := col.SATResult.Contact1
	// fallback: средняя точка между центрами (хуже, но безопасно)
	if contact == (m.Vector2{}) {
		contact = bodyA.Position.Add(bodyB.Position).Scale(0.5)
	}

	// POSITONAL CORRECTION (чтобы убрать проникание)
	const percent = 0.8   // сколько исправлять (0..1)
	const slop = 0.01     // небольшая "погрешность", которую не исправляем
	overlap := col.SATResult.Overlap
	if overlap > slop {
		var invMassA, invMassB float64
		if bodyA.Mass > 0 {
			invMassA = 1.0 / bodyA.Mass
		}
		if bodyB.Mass > 0 {
			invMassB = 1.0 / bodyB.Mass
		}
		correctionMag := (math.Max(overlap-slop, 0.0) / (invMassA + invMassB)) * percent
		correction := n.Scale(correctionMag)
		if bodyA.Mass > 0 {
			bodyA.Position = bodyA.Position.Sub(correction.Scale(invMassA))
		}
		if bodyB.Mass > 0 {
			bodyB.Position = bodyB.Position.Add(correction.Scale(invMassB))
		}
	}

	// Векторы от центров масс до точки контакта
	rA := contact.Sub(bodyA.Position)
	rB := contact.Sub(bodyB.Position)

	// Скорость точки контакта учитывая угловую скорость: v + ω × r
	velA_contact := bodyA.Velocity.Add(m.Vector2{X: -bodyA.AngVel * rA.Y, Y: bodyA.AngVel * rA.X})
	velB_contact := bodyB.Velocity.Add(m.Vector2{X: -bodyB.AngVel * rB.Y, Y: bodyB.AngVel * rB.X})

	// Относительная скорость в точке контакта (B относительно A)
	relVel := velB_contact.Sub(velA_contact)

	// Нормальная составляющая скорости (если скорость направлена в разлёт — игнорируем)
	velAlongNormal := relVel.Dot(n)
	if velAlongNormal > 0 {
		return
	}

	// Параметры материала
	restitution := 0.8 // упругость, можно вынести в Body или в коллайдер
	mu := 0.5          // коэффициент трения (кулона)

	// Инверсные массы / инверсия инерции
	var invMassA, invMassB float64
	if bodyA.Mass > 0 {
		invMassA = 1.0 / bodyA.Mass
	}
	if bodyB.Mass > 0 {
		invMassB = 1.0 / bodyB.Mass
	}
	var invInertiaA, invInertiaB float64
	if bodyA.Inertia > 0 {
		invInertiaA = 1.0 / bodyA.Inertia
	}
	if bodyB.Inertia > 0 {
		invInertiaB = 1.0 / bodyB.Inertia
	}

	// Кроссы r × n (в 2D это скаляр)
	crossRA_N := rA.X*n.Y - rA.Y*n.X
	crossRB_N := rB.X*n.Y - rB.Y*n.X

	// Знаменатель для нормального импульса (учитывает вращение)
	denom := invMassA + invMassB + (crossRA_N*crossRA_N)*invInertiaA + (crossRB_N*crossRB_N)*invInertiaB
	if denom == 0 {
		return
	}

	// Нормальный импульс J
	j := -(1 + restitution) * velAlongNormal / denom
	impulse := n.Scale(j)

	// Применяем нормальный импульс: линейный + угловой
	if bodyA.Mass > 0 {
		bodyA.Velocity = bodyA.Velocity.Sub(impulse.Scale(invMassA))
		// Δω = - (rA × impulse) / I_A  (знак минус потому что A получает -impulse)
		bodyA.AngVel -= (rA.X*impulse.Y - rA.Y*impulse.X) * invInertiaA
	}
	if bodyB.Mass > 0 {
		bodyB.Velocity = bodyB.Velocity.Add(impulse.Scale(invMassB))
		bodyB.AngVel += (rB.X*impulse.Y - rB.Y*impulse.X) * invInertiaB
	}

	// --- Касательное трение (Coulomb) ---
	// tangent direction
	tangent := relVel.Sub(n.Scale(relVel.Dot(n)))
	if tangent.Length() > 1e-9 {
		t := tangent.Normalize()

		// cross for tangent
		crossRA_T := rA.X*t.Y - rA.Y*t.X
		crossRB_T := rB.X*t.Y - rB.Y*t.X

		denomT := invMassA + invMassB + (crossRA_T*crossRA_T)*invInertiaA + (crossRB_T*crossRB_T)*invInertiaB
		if denomT > 0 {
			// jt = - v_rel·t / denomT
			jt := -relVel.Dot(t) / denomT

			// ограничиваем величину по закону Кулона: |jt| <= mu * j_normal
			maxF := mu * math.Abs(j)
			if math.Abs(jt) > maxF {
				jt = math.Copysign(maxF, jt)
			}

			impulseT := t.Scale(jt)

			if bodyA.Mass > 0 {
				bodyA.Velocity = bodyA.Velocity.Sub(impulseT.Scale(invMassA))
				bodyA.AngVel -= (rA.X*impulseT.Y - rA.Y*impulseT.X) * invInertiaA
			}
			if bodyB.Mass > 0 {
				bodyB.Velocity = bodyB.Velocity.Add(impulseT.Scale(invMassB))
				bodyB.AngVel += (rB.X*impulseT.Y - rB.Y*impulseT.X) * invInertiaB
			}
		}
	}
}

