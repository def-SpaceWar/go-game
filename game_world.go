package main

import (
	"gogame/ecs"
	"gogame/ecs_time"
	"gogame/fp"
	"gogame/physics2d"
	"gogame/render2d"
)

func gameWorld() *ecs.World {
	world := ecs.CreateWorld()

	world.CreateEntity(
		physics2d.NewTransform(physics2d.TransformParams{
			Position: physics2d.Vec(200, 200),
			Rotation: fp.Some[float32](0),
			Scale:    fp.None[physics2d.Vector](),
		}),
		render2d.NewPolygon(render2d.PolygonParams{
			Points: render2d.Rect(0, 0, 100, 100),
			Color:  fp.Some(render2d.RGB(255, 0, 0)),
		}),
	)

	player := world.CreateEntity(
		physics2d.NewTransform(physics2d.TransformParams{
			Position: physics2d.Vec(100, 300),
			Rotation: fp.Some[float32](0),
			Scale:    fp.None[physics2d.Vector](),
		}),
		render2d.NewPolygon(render2d.PolygonParams{
			Points: render2d.Rect(0, 0, 100, 100),
			Color:  fp.Some(render2d.RGB(255, 0, 0)),
		}),
	)
	transform := ecs.GetComponent[physics2d.Transform](player)
	polygon := ecs.GetComponent[render2d.Polygon](player)

	world.CreateChildEntity(player,
		render2d.NewPolygon(render2d.PolygonParams{
			Points: render2d.Rect(100, 100, 100, 100),
			Color:  fp.Some(render2d.RGB(255, 100, 0)),
		}),
		render2d.NewPolygon(render2d.PolygonParams{
			Points: render2d.Rect(-100, -100, 100, 100),
			Color:  fp.Some(render2d.RGB(255, 100, 0)),
		}),
	)

	world.AddSystems(
		ecs_time.CreateTimedSystem(
			"render",
			renderSystem,
		),
		ecs_time.CreateTimedSystem(
            "animate",
			func(_ *ecs.World) (ecs.WorldState, *ecs.World) {
                deltaTime, _ := ecs_time.DeltaTimes.Get("animate")
				transform.Position.X += 100 * deltaTime
				*transform.Rotation += 5 * deltaTime
				polygon.Color.R += 50 * deltaTime
				polygon.Color.G += 10 * deltaTime
				polygon.Color.B += 20 * deltaTime
				if transform.Position.X > float32(render2d.Renderer.GetViewport().W+200) {
					transform.Position.X = -200
				}
                return ecs.CONTINUE, nil
			},
		),
	)

	return &world
}
