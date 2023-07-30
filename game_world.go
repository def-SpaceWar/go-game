package main

import (
	"gogame/ecs"
	"gogame/render2d"

	"github.com/veandco/go-sdl2/sdl"
)

func gameWorld() *ecs.World {
	world := ecs.CreateWorld()

    player := world.CreateEntity(
		render2d.Rectangle{
			Rect:  &sdl.Rect{X: 0, Y: 0, W: 200, H: 200},
			Color: sdl.MapRGB(render2d.Surface.Format, 255, 0, 0),
		},
	)
    playerRect := ecs.GetComponent[render2d.Rectangle](player)

	world.AddSystems(
        renderSystem,
        func (_ *ecs.World) {
            playerRect.Rect.X += 1
        },
	)

	return &world
}
