package render2d

import "github.com/veandco/go-sdl2/sdl"

func Rect(x, y, w, h int32) *sdl.Rect {
	return &sdl.Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
}

type Rectangle struct {
	Rect   *sdl.Rect
	Color  uint32
	zIndex int32
}

func (r Rectangle) ZIndex() int32 {
	return r.zIndex
}
