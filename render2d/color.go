package render2d

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Color struct {
	R float32
	G float32
	B float32
	A float32
}

func RGB(r, g, b float32) Color {
	return Color{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}

func (color *Color) ToSDL() *sdl.Color {
    color.R = cap255(color.R)
    color.G = cap255(color.G)
    color.B = cap255(color.B)
    color.A = cap255(color.A)

    return &sdl.Color{
        R: uint8(color.R),
        G: uint8(color.G),
        B: uint8(color.B),
        A: uint8(color.A),
    }
}

func cap255(v float32) float32 {
    ints, remain := math.Modf(float64(v))
    myInt := int32(ints)
    myInt = myInt % 256
    result := float32(myInt) + float32(remain)
    if result > 255 {
        result -= 255
    }
    return result
}
