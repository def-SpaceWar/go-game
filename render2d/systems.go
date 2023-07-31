package render2d

import (
	"gogame/ecs"

	"github.com/veandco/go-sdl2/sdl"
)

var Window *sdl.Window
var Surface *sdl.Surface
var DeltaTime float32

type Drawable interface {
	ZIndex() int32
}

func CreateRenderSystem() ecs.System {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	Window, err = sdl.CreateWindow(
		"test",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		800,
		600,
		sdl.WINDOW_SHOWN,
	)
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
			case *sdl.KeyboardEvent:
				event := event.(*sdl.KeyboardEvent)
				switch event.State {
				case sdl.PRESSED:
					if event.Repeat == 0 {
						println("a")
					}
				case sdl.RELEASED:
					println("b")
				}
			}
		}

		Surface.FillRect(nil, 0)
		for rect := range ecs.FindComponents[Rectangle](world) {
			draw(rect)
		}
		Window.UpdateSurface()
	}
}

func draw(drawable Drawable) {
	rect, isRect := drawable.(*Rectangle)
	if isRect {
		Surface.FillRect(rect.Rect, rect.Color)
	}
}
