package main

import (
	"gogame/ecs"
	"gogame/fp"
	"gogame/physics2d"
	"gogame/render2d"
)

func gameWorld() *ecs.World {
	world := ecs.CreateWorld()

    player := world.CreateEntity(
		physics2d.NewTransform(physics2d.TransformParams{
			Position: physics2d.Vec(100, 100),
			Rotation: fp.Some[float32](0),
			Scale:    fp.None[physics2d.Vector](),
		}),
		render2d.NewPolygon(render2d.PolygonParams{
			Points: render2d.Rect(0, 0, 100, 100),
			Color:  render2d.RGB(255, 0, 0),
		}),
	)
	transform := ecs.GetComponent[physics2d.Transform](player)
	polygon := ecs.GetComponent[render2d.Polygon](player)

	world.CreateEntity(
		physics2d.NewTransform(physics2d.TransformParams{
			Position: physics2d.Vec(200, 200),
			Rotation: fp.Some[float32](0),
			Scale:    fp.None[physics2d.Vector](),
		}),
		render2d.NewPolygon(render2d.PolygonParams{
			Points: render2d.Rect(0, 0, 100, 100),
			Color:  render2d.RGB(255, 0, 0),
		}),
	)

	// Render Pipeline
	world.AddSystems(
		renderSystem,
		func(_ *ecs.World) {
			transform.Position.X += .1
			*transform.Rotation += .05
			polygon.Color.R += 5
			polygon.Color.G += 1
			polygon.Color.B += 2
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
