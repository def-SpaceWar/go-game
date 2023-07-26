package primitive

import (
	"gogame/ecs"

	"github.com/veandco/go-sdl2/sdl"
)

func CreateRenderSystem(window *sdl.Window, surface *sdl.Surface) ecs.System {
	return func(world *ecs.World) {
		surface.FillRect(&sdl.Rect{X: 0, Y: 0, W: surface.W, H: surface.H}, 0)
		for _, e := range world.Entities() {
			position := Position[e]
			if position != nil {
				rect := sdl.Rect{
					X: int32(position.X) - 50,
					Y: int32(position.Y) - 50,
					W: 100,
					H: 100,
				}

				color := sdl.Color{R: 100, G: 255, B: 255}
				pixel := sdl.MapRGBA(surface.Format, color.R, color.G, color.B, 255)
				surface.FillRect(&rect, pixel)
			}
		}
		window.UpdateSurface()
	}
}
