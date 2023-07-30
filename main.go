package main

import (
	"gogame/ecs"
	"gogame/render2d"
)

var renderSystem = render2d.CreateRenderSystem()

func main() {
	ecs.CurrentWorld = gameWorld()

	for ecs.Running {
		ecs.CurrentWorld.RunSystems()
	}
}
