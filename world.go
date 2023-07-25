package main

type World struct {
    entityCount Entity
    components []Component
    systems []System
}

func (w *World) GenerateEntity() Entity {
	defer (func() { w.entityCount++ })()
	return w.entityCount
}

func (w *World) RunSystems() {
    for i := 0; i < len(w.systems); i++ {
        w.systems[i](w)
    }
}
