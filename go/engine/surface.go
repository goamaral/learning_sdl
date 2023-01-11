package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Surface struct {
	abstractResource

	sdlSurface *sdl.Surface
	Window     *Window
}

/* PRIVATE */
var surfaceManager = newResourceManager()

func newSurface(sdlSurface *sdl.Surface, window *Window) *Surface {
	return surfaceManager.Save(&Surface{sdlSurface: sdlSurface, Window: window}).(*Surface)
}

/* PUBLIC */
func (s *Surface) Destroy() {
	if !s.destroyed {
		s.destroyed = true
		delete(surfaceManager.resources, s.id)
		s.sdlSurface.Free()
	}
}

func (s *Surface) OptimizeSurface() error {
	optimisedSdlSurface, err := s.sdlSurface.Convert(s.Window.GetPixelFormat(), 0)
	if err != nil {
		return err
	}

	s.sdlSurface.Free()
	s.sdlSurface = optimisedSdlSurface

	return nil
}

func (s *Surface) SetTransparentColor(color Color) error {
	return s.sdlSurface.SetColorKey(true, color.ToColorKey(s.sdlSurface.Format))
}

func (s *Surface) GetPixelFormat() *sdl.PixelFormat {
	return s.sdlSurface.Format
}
