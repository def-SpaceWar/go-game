package render2d

import (
	"gogame/ecs"

	"github.com/veandco/go-sdl2/sdl"
)

var Window *sdl.Window
var Surface *sdl.Surface

func CreateRenderSystem() ecs.System {
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	Window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	Surface, err = Window.GetSurface()
	if err != nil {
		panic(err)
	}

	return func(world *ecs.World) {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				ecs.Running = false
			}
		}

		Surface.FillRect(nil, 0)
		for rect := range ecs.FindComponents[Rectangle](world) {
			Surface.FillRect(rect.Rect, rect.Color)
		}
		Window.UpdateSurface()
	}
}
