package main

import (
	"gogame/ecs"
	"gogame/ecs/primitive"
	"gogame/physics"
)

func game() *ecs.World {
	resetComponents()
	world := ecs.CreateWorld([]ecs.System{
        primitive.CreateTimeSystem(),
        primitive.ForcesSystem,
        primitive.CreateRenderSystem(Window, Surface),
    }, onEntityCreate, onEntityDestroy)

	player := world.CreateEntity()
	primitive.Position[player] = physics.Vec(100, 100)
	primitive.Velocity[player] = physics.Vec(100, 100)

	enemy := world.CreateEntity()
	primitive.Position[enemy] = physics.Vec(
        float32(Surface.W) - 100,
        float32(Surface.H) - 100,
    )

	return &world
}
