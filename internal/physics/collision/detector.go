package collision

import (
	"github.com/Crusazer/tanks-race/internal/physics/dynamics"
	"github.com/Crusazer/tanks-race/internal/physics/shapes"
)

type Detector struct {
	results []CollisionResult
}

type CollisionResult struct {
	BodyA     *dynamics.Body
	BodyB     *dynamics.Body
	SATResult *SATResult
}

func NewDetector() *Detector {
	return &Detector{
		results: make([]CollisionResult, 0, 100),
	}
}

func (d *Detector) Detect(rectangles []*shapes.Rectangle, bodies []*dynamics.Body) []CollisionResult {
	d.results = d.results[:0] // Очищаем слайс

	// Проверяем все пары
	for i := 0; i < len(rectangles); i++ {
		for j := i + 1; j < len(rectangles); j++ {
			// Проверяем, что индексы в пределах слайса bodies
			if i < len(bodies) && j < len(bodies) {
				result := CheckCollision(rectangles[i], rectangles[j])
				if result != nil {
					d.results = append(d.results, CollisionResult{
						BodyA:     bodies[i],
						BodyB:     bodies[j],
						SATResult: result,
					})
				}
			}
		}
	}

	return d.results
}