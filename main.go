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
	currentWorld := gameWorld()
	running := true
	for running {
		for _, system := range currentWorld.Systems {
			state, nextWorld := system(currentWorld)
			if state == ecs.QUIT {
				running = false
			} else if state == ecs.SWITCH {
				currentWorld = nextWorld
			}
		}
	}
}
