package entity

import "sort"

/*
Служит для последовательного отображения спрайтов по слаям
(чтобы перекрывали друг друга в правильном порядке в зависимости от Layer)
*/
type RenderQueue struct {
	Layers map[int][]*Entity
	Order  []int
}

func NewRenderQueue() *RenderQueue {
	return &RenderQueue{
		Layers: make(map[int][]*Entity),
		Order:  []int{},
	}
}

func (rq *RenderQueue) Add(e *Entity) {
	sprite, ok := e.Components[SpriteComponent].(*Sprite)
	if !ok {
		return
	}
	layer := sprite.Layer

	// Если слой новый — добавляем в order
	if _, exists := rq.Layers[layer]; !exists {
		rq.Order = append(rq.Order, layer)
		sort.Ints(rq.Order)
	}
	rq.Layers[layer] = append(rq.Layers[layer], e)
}

func (rq *RenderQueue) Remove(e *Entity) {
	sprite, ok := e.Components[SpriteComponent].(*Sprite)
	if !ok {
		return
	}
	layer := sprite.Layer
	entities := rq.Layers[layer]
	for i, en := range entities {
		if en.ID == e.ID {
			rq.Layers[layer] = append(entities[:i], entities[i+1:]...)
			break
		}
	}
	if len(rq.Layers[layer]) == 0 {
		delete(rq.Layers, layer)
		newOrder := []int{}
		for _, l := range rq.Order {
			if l != layer {
				newOrder = append(newOrder, l)
			}
		}
		rq.Order = newOrder
	}
}

func (rq *RenderQueue) Clear() {
	rq.Layers = make(map[int][]*Entity)
	rq.Order = []int{}
}
