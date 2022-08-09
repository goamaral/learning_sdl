package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Surface struct {
	*sdl.Surface
	ID uint32
}

// Optimize surface
func (s *Surface) OptimizeSurface(pixelFormat *sdl.PixelFormat) error {
	optimisedSdlSurface, err := s.Convert(pixelFormat, 0)
	if err != nil {
		return err
	}

	s.Surface.Free()
	s.Surface = optimisedSdlSurface

	return nil
}

// Set color key
func (s *Surface) SetColorKey(color Color) error {
	return s.Surface.SetColorKey(true, sdl.MapRGB(s.Format, color.Red, color.Green, color.Blue))
}
