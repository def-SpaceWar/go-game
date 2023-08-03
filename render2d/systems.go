package render2d

import (
	"gogame/ecs"
	"gogame/fp"
	"gogame/physics2d"

	"github.com/veandco/go-sdl2/sdl"
)

var Window *sdl.Window
var Renderer *sdl.Renderer

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

	vsync := true
	if !fp.IsNone(params.VSYNC) {
		vsync = fp.Just(params.VSYNC)
	}
	if vsync {
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
	if err != nil {
		panic(err)
	}

	return func(world *ecs.World) (ecs.WorldState, *ecs.World) {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
                Renderer.Destroy()
                Window.Destroy()
                sdl.Quit()
                return ecs.QUIT, nil
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

        return ecs.CONTINUE, nil
	}
}

func getTriangles(entity ecs.Entity, drawable Drawable) []sdl.Vertex {
	polygon, isPolygon := drawable.(*Polygon)
	if isPolygon {
		transforms := []*physics2d.Transform{}
		for e := &entity; e != nil; e = e.Parent {
			transform := ecs.GetComponent[physics2d.Transform](e)
			if transform != nil {
				transforms = append(transforms, transform)
			}
		}
		return polygon.toVertices(transforms)
	}

	return []sdl.Vertex{}
}
