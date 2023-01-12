package engine

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Supports bmp, png
func (w *Window) LoadSurface(path string) (*Surface, error) {
	pathParts := strings.Split(path, ".")
	extension := pathParts[len(pathParts)-1]
	var sdlSurface *sdl.Surface
	var err error

	switch extension {
	case "bmp":
		sdlSurface, err = sdl.LoadBMP(path)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load bmp surface")
		}
	case "png":
		sdlSurface, err = img.Load(path)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load png surface")
		}
	default:
		return nil, errors.Errorf("unsupported surface file format %s")
	}

	surface := newSurface(sdlSurface, w)
	err = surface.OptimizeSurface()
	if err != nil {
		surface.Destroy()
		log.Warn().Err(err).Str("path", path).Msg("Failed to optimize loaded surface")
	}

	return surface, nil
}

func (w *Window) RenderSurface(surface *Surface, scaled bool) error {
	if scaled {
		err := surface.sdlSurface.BlitScaled(nil, w.Surface.sdlSurface, nil)
		if err != nil {
			return errors.Wrap(err, "failed to copy scaled surface to window surface")
		}
	} else {
		err := surface.sdlSurface.Blit(nil, w.Surface.sdlSurface, nil)
		if err != nil {
			return errors.Wrap(err, "failed to copy surface to window surface")
		}
	}

	err := w.sdlWindow.UpdateSurface()
	if err != nil {
		return errors.Wrap(err, "failed to render window surface")
	}

	return nil
}
