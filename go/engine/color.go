package engine

import "github.com/veandco/go-sdl2/sdl"

type Color sdl.Color

type ColorName string

var (
	Color_BLACK   = Color{0, 0, 0, 255}
	Color_WHITE   = Color{255, 255, 255, 255}
	Color_RED     = Color{255, 0, 0, 255}
	Color_GREEN   = Color{0, 255, 0, 255}
	Color_BLUE    = Color{0, 0, 255, 255}
	Color_YELLOW  = Color{255, 255, 0, 255}
	Color_MAGENTA = Color{255, 0, 255, 255}
	Color_CYAN    = Color{0, 255, 255, 255}
)

func (c Color) ToColorKey(format *sdl.PixelFormat) uint32 {
	return sdl.MapRGB(format, c.R, c.G, c.B)
}
