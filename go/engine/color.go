package engine

import "github.com/veandco/go-sdl2/sdl"

type Color sdl.Color

type ColorName string

const (
	ColorName_BLACK   ColorName = "black"
	ColorName_WHITE   ColorName = "white"
	ColorName_RED     ColorName = "red"
	ColorName_GREEN   ColorName = "green"
	ColorName_BLUE    ColorName = "blue"
	ColorName_YELLOW  ColorName = "yellow"
	ColorName_MAGENTA ColorName = "magenta"
	ColorName_CYAN    ColorName = "cyan"
)

var ColorByName map[ColorName]Color = map[ColorName]Color{
	ColorName_BLACK:   {0, 0, 0, 255},
	ColorName_WHITE:   {255, 255, 255, 255},
	ColorName_RED:     {255, 0, 0, 255},
	ColorName_GREEN:   {0, 255, 0, 255},
	ColorName_BLUE:    {0, 0, 255, 255},
	ColorName_YELLOW:  {255, 255, 0, 255},
	ColorName_MAGENTA: {255, 0, 255, 255},
	ColorName_CYAN:    {0, 255, 255, 255},
}

func (c Color) ToColorKey(format *sdl.PixelFormat) uint32 {
	return sdl.MapRGB(format, c.R, c.G, c.B)
}
