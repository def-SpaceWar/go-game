package main

import (
	"gogame/ecs"
	"gogame/fp"
	"gogame/render2d"
)

var renderSystem = render2d.CreateRenderSystem(render2d.RenderSystemParams{
    VSYNC: fp.Some(false),
})

func main() {
	ecs.CurrentWorld = gameWorld()

	for ecs.Running {
		ecs.CurrentWorld.RunSystems()
	}
}
