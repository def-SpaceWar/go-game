package render2d

import "github.com/veandco/go-sdl2/sdl"

type Rectangle struct {
	Rect   *sdl.Rect
	Color  uint32
	ZIndex int32
}
