package engine

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Save surface
func (w *Window) SaveSurface(surface *Surface) {
	surface.ID = w.GetNextResourceID()
	w.surfaces[surface.ID] = *surface
}

// Get surface
func (w *Window) GetSurface(id uint32) (Surface, error) {
	surface, exists := w.surfaces[id]
	if !exists {
		return surface, errors.Errorf("surface %d not found", id)
	}

	return surface, nil
}

// Destroy surface
func (w *Window) DestroySurface(id uint32) error {
	surface, exists := w.surfaces[id]
	if !exists {
		return errors.Errorf("surface %d not found", id)
	}

	delete(w.surfaces, id)
	surface.Free()

	return nil
}

// Load surface (supports bmp, png)
func (w *Window) LoadSurface(path string) (surface Surface, err error) {
	pathParts := strings.Split(path, ".")

	switch pathParts[len(pathParts)-1] {
	case "bmp":
		surface.Surface, err = sdl.LoadBMP(path)
		if err != nil {
			return surface, errors.Wrap(err, "failed to load surface from bmp")
		}
	case "png":
		surface.Surface, err = img.Load(path)
		if err != nil {
			return surface, errors.Wrap(err, "failed to load surface from png")
		}
	}

	err = surface.OptimizeSurface(w.surface.Format)
	if err != nil {
		log.Warn().Err(err).Str("path", path).Msg("Failed to optimize loaded surface")
	}

	w.SaveSurface(&surface)

	return surface, nil
}

// Render surface
func (w *Window) RenderSurface(id uint32, scaled bool) error {
	surface, err := w.GetSurface(id)
	if err != nil {
		return errors.Wrap(err, "failed to get surface")
	}

	if scaled {
		err := surface.BlitScaled(nil, w.surface, nil)
		if err != nil {
			return errors.Wrap(err, "failed to copy scaled surface to window surface")
		}
	} else {
		err := surface.Blit(nil, w.surface, nil)
		if err != nil {
			return errors.Wrap(err, "failed to copy surface to window surface")
		}
	}

	err = w.UpdateSurface()
	if err != nil {
		return errors.Wrap(err, "failed to render window surface")
	}

	return nil
}

// Convert surface to texture
func (w *Window) ConvertSurfaceToTexture(id uint32) (Texture, error) {
	// Get surface
	surface, err := w.GetSurface(id)
	if err != nil {
		return Texture{}, errors.Wrap(err, "failed to get surface")
	}

	// Create texture from surface
	texture, err := w.Renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return texture, errors.Wrap(err, "failed to create texture from surface")
	}

	// Save texture
	w.SaveTexture(&texture)

	return texture, nil
}
