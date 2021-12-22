package surface

import (
	"fmt"
	"strings"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Surface struct {
	*sdl.Surface
}

// Load surface (supports bmp, png)
func Load(path string, pixelFormat *sdl.PixelFormat) Surface {
	var surface *sdl.Surface
	var err error

	pathParts := strings.Split(path, ".")
	switch pathParts[len(pathParts)-1] {
	case "bmp":
		surface, err = sdl.LoadBMP(path)
		if err != nil {
			panic(err)
		}
	case "png":
		surface, err = img.Load(path)
		if err != nil {
			panic(err)
		}
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
