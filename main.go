package main

import (
	"gogame/ecs"

	"github.com/veandco/go-sdl2/sdl"
)

var Window *sdl.Window
var Surface *sdl.Surface
var CurrentWorld *ecs.World

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	var err error
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
	defer Window.Destroy()

	Surface, err = Window.GetSurface()
	if err != nil {
		panic(err)
	}

	Surface.FillRect(nil, 0)
	Window.UpdateSurface()
	CurrentWorld = game()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Closed Window")
				running = false
			}
		}

        CurrentWorld.RunSystems()
	}
}
