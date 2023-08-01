package render2d

import (
	"gogame/ecs"
	"gogame/fp"
	"gogame/physics2d"

	"github.com/veandco/go-sdl2/sdl"
)

var Window *sdl.Window
var Renderer *sdl.Renderer
var DeltaTime float32

type Drawable interface {
	ZIndex() int32
}

type RenderSystemParams struct {
	VSYNC fp.Maybe[bool]
}

func CreateRenderSystem(params RenderSystemParams) ecs.System {
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

	if !fp.IsNone(params.VSYNC) {
		if fp.Just(params.VSYNC) {
			Renderer, err = sdl.CreateRenderer(
				Window,
				-1,
				sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC,
			)
		} else {
			Renderer, err = sdl.CreateRenderer(
				Window,
				-1,
				sdl.RENDERER_ACCELERATED,
			)
		}
	} else {
		Renderer, err = sdl.CreateRenderer(
			Window,
			-1,
			sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC,
		)
	}
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
		totalTriangles := []sdl.Vertex{}
		for pair := range ecs.FindComponents[Polygon](world) {
			entity := pair.First
			polygon := pair.Second
			totalTriangles = append(totalTriangles, getTriangles(entity, polygon)...)
		}
		Renderer.RenderGeometry(nil, totalTriangles, nil)
		Renderer.Present()
	}
}

func getTriangles(entity ecs.Entity, drawable Drawable) []sdl.Vertex {
	polygon, isPolygon := drawable.(*Polygon)
	if isPolygon {
		transforms := []*physics2d.Transform{}
		transforms = append(transforms, ecs.GetComponent[physics2d.Transform](&entity))
		// TODO Get all parent transforms too!
		return polygon.toVertices(transforms)
	}

	return []sdl.Vertex{}
}
