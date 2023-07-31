package render2d

import "github.com/veandco/go-sdl2/sdl"

func RGB(r, g, b uint8) uint32 {
    return sdl.MapRGB(Surface.Format, r, g, b)
}

func RGBA(r, g, b, a uint8) uint32 {
    return sdl.MapRGBA(Surface.Format, r, g, b, a)
}
