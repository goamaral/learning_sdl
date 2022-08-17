package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

var surfaces map[uint]*Surface = map[uint]*Surface{}
var newSurfaceId uint

type Surface struct {
	sdlSurface *sdl.Surface
	Window     *Window
	id         uint
	destroyed  bool
}

/* PRIVATE */
func (s *Surface) save() *Surface {
	s.id = newSurfaceId
	newSurfaceId += 1
	surfaces[s.id] = s
	return s
}

func destroySurfaces() {
	for _, surface := range surfaces {
		surface.Destroy()
	}
}

func (s *Surface) getSdlSurface() *sdl.Surface {
	return s.sdlSurface
}

/* PUBLIC */
func NewSurface(sdlSurface *sdl.Surface) *Surface {
	return (&Surface{sdlSurface: sdlSurface}).save()
}

func (s *Surface) Destroy() {
	if !s.destroyed {
		s.destroyed = true
		delete(surfaces, s.id)
	}
}

func (s *Surface) OptimizeSurface(window *Window) error {
	optimisedSdlSurface, err := s.sdlSurface.Convert(window.GetPixelFormat(), 0)
	if err != nil {
		return err
	}

	s.sdlSurface.Free()
	s.sdlSurface = optimisedSdlSurface
	s.Window = window

	return nil
}

func (s *Surface) SetTransparentColor(color Color) error {
	return s.sdlSurface.SetColorKey(true, color.ToColorKey(s.sdlSurface.Format))
}
