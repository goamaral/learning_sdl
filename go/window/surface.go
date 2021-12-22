package window

import (
	eng "github.com/Goamaral/learning_sdl/engine"
	srf "github.com/Goamaral/learning_sdl/surface"
)

// Get window surface
func (w *Window) GetSurface() srf.Surface {
	sdlSurface, err := w.Window.GetSurface()
	if err != nil {
		panic(err)
	}
	return srf.Surface{Surface: sdlSurface}
}

// Render surface
func (w *Window) RenderRawSurface(surface *srf.Surface) {
	err := surface.Blit(nil, w.GetSurface().Surface, nil)
	if err != nil {
		panic(err)
	}

	w.UpdateSurface()
}

// Render engine surface
func (w *Window) RenderSurface(name string) error {
	surface, err := eng.GetSurface(name)
	if err != nil {
		return err
	}

	err = surface.Blit(nil, w.GetSurface().Surface, nil)
	if err != nil {
		panic(err)
	}

	w.UpdateSurface()

	return nil
}
