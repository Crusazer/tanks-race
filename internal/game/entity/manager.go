package entity

import (
	"sync"
)

type Manager struct {
	entities map[ID]*Entity
	nextID   ID
	mu       sync.RWMutex
	RenderQueue *RenderQueue
}

func NewManager() *Manager {
	return &Manager{
		entities:    make(map[ID]*Entity),
		RenderQueue: NewRenderQueue(),
	}
}

func (m *Manager) Create() *Entity {
	m.mu.Lock()
	defer m.mu.Unlock()

	id := m.nextID
	m.nextID++

	entity := &Entity{
		ID:         id,
		Components: make(map[ComponentType]interface{}),
	}

	m.entities[id] = entity
	return entity
}

func (m *Manager) Get(id ID) *Entity {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.entities[id]
}

func (m *Manager) Remove(id ID) {
	m.mu.Lock()
	defer m.mu.Unlock()

	e := m.entities[id]
	if e != nil {
		m.RenderQueue.Remove(e)
	}
	delete(m.entities, id)
}

func (m *Manager) SetComponent(e *Entity, t ComponentType, comp interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	e.Components[t] = comp

	if t == SpriteComponent {
		m.RenderQueue.Add(e)
	}
}

func (m *Manager) GetWithComponents(types ...ComponentType) []*Entity {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var result []*Entity

	for _, entity := range m.entities {
		hasAll := true
		for _, t := range types {
			if _, ok := entity.Components[t]; !ok {
				hasAll = false
				break
			}
		}
		if hasAll {
			result = append(result, entity)
		}
	}
	return result
}
