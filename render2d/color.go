package render2d

import "github.com/veandco/go-sdl2/sdl"

func RGB(r, g, b uint8) *sdl.Color {
    return &sdl.Color{
        R: r,
        G: g,
        B: b,
        A: 255,
    }
}
