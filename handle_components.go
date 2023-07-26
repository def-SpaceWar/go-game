package main

import (
	"gogame/ecs"
	"gogame/ecs/primitive"
)

// Resets all the component lists
func resetComponents() {
	primitive.Position = primitive.Position[:0]
	primitive.Velocity = primitive.Velocity[:0]
}

// Add spaces for new entity on all the components
func onEntityCreate(e ecs.Entity) {
	for len(primitive.Position) <= int(e) {
		primitive.Position = append(primitive.Position, nil)
		primitive.Velocity = append(primitive.Velocity, nil)
	}
}

// Deletes all of the destroyed entity's components
func onEntityDestroy(e ecs.Entity) {
	primitive.Position[e] = nil
	primitive.Velocity[e] = nil
}
