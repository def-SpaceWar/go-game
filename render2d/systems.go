package render2d

import (
	"gogame/ecs"
	"gogame/physics2d"

	"github.com/veandco/go-sdl2/sdl"
)

var Window *sdl.Window
var Renderer *sdl.Renderer
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

	Renderer, err = sdl.CreateRenderer(
        Window,
        -1,
        sdl.RENDERER_ACCELERATED |
        sdl.RENDERER_PRESENTVSYNC,
    )
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

		Renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)
        Renderer.Clear()
		for pair := range ecs.FindComponents[Polygon](world) {
			entity := pair.First
			rect := pair.Second
			draw(entity, rect)
		}
        Renderer.Present()
	}
}

func draw(entity *ecs.Entity, drawable Drawable) {
	polygon, isPolgyon := drawable.(*Polygon)
	if isPolgyon {
        transforms := []*physics2d.Transform{}
        transforms = append(transforms, ecs.GetComponent[physics2d.Transform](entity))
        // TODO Get all parent transforms too!
		Renderer.RenderGeometry(
            nil,
            polygon.toVertices(transforms),
            nil,
        )
	}
}
