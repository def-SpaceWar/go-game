package main

import "fmt"

type World struct {
	entityCount Entity
	components  []Component
	systems     []System
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

func (w *World) GetComponent(e Entity, compType *string) (*Component, error) {
	for i := 0; i < len(w.components); i++ {
		component := w.components[i]
		if component.Entity() == e && component.data.Type() == compType {
			return &component, nil
		}
	}

	return nil, fmt.Errorf("Entity #%d does not have a %s component!", e, *compType)
}
