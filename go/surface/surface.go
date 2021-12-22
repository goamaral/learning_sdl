package surface

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Surface struct {
	*sdl.Surface
}

// Load surface from bmp
func LoadFromBmp(path string, pixelFormat *sdl.PixelFormat) Surface {
	surface, err := sdl.LoadBMP(path)
	if err != nil {
		panic(err)
	}

	optimezedSurface, err := optimizeSurface(surface, pixelFormat)
	if err != nil {
		fmt.Println("Failed to optimize surface")
		return Surface{surface}
	} else {
		surface.Free()
		return Surface{optimezedSurface}
	}
}

// Optimize surface
func optimizeSurface(surface *sdl.Surface, pixelFormat *sdl.PixelFormat) (*sdl.Surface, error) {
	return surface.Convert(pixelFormat, 0)
}
