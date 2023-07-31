package main

import (
	"gogame/ecs"
	"gogame/render2d"
)

func gameWorld() *ecs.World {
	world := ecs.CreateWorld()

	player := world.CreateEntity(
		render2d.Rectangle{
			Rect:  render2d.Rect(100, 100, 100, 100),
			Color: render2d.RGB(255, 0, 0),
		},
	)
	playerRect := ecs.GetComponent[render2d.Rectangle](player)

	// Render Pipeline
	world.AddSystems(
		renderSystem,
		func(_ *ecs.World) {
			playerRect.Rect.X += 1
		},
	)

	// Physics Pipeline
	//world.AddSystems(
	//    physics2d.ForcesSystem,
	//    physics2d.CollisionSystem,
	//)

	// Game Pipeline
	//world.AddSystems(
	//    playerMovementSystem,
	//)

	return &world
}
