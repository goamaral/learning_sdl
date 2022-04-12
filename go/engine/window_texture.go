package engine

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

// Save texture
func (w *Window) SaveTexture(texture *Texture) {
	texture.ID = w.GetNextResourceID()
	w.textureMap[texture.ID] = *texture
}

// Get texture
func (w *Window) GetTexture(id uint64) (Texture, error) {
	texture, exists := w.textureMap[id]
	if !exists {
		return texture, errors.Errorf("texture %d not found", id)
	}

	return texture, nil
}

// Destroy texture
func (w *Window) DestroyTexture(id uint64) error {
	texture, exists := w.textureMap[id]
	if !exists {
		return errors.Errorf("texture %d not found", id)
	}

	delete(w.textureMap, id)
	return texture.Destroy()
}

// Render texture
func (w *Window) RenderTexture(id uint64, viewportId uint64, dst *sdl.Rect) error {
	// Set viewport
	err := w.SetViewportByID(viewportId)
	if err != nil {
		return errors.Wrap(err, "failed to set viewport")
	}

	// Get texture
	texture, err := w.GetTexture(id)
	if err != nil {
		return errors.Wrap(err, "failed to get texture")
	}

	// Set destination rect default width and height
	if dst != nil {
		if dst.W == 0 {
			dst.W = texture.W
		}
		if dst.H == 0 {
			dst.H = texture.H
		}
	}

	// Render texture
	err = w.renderer.Copy(texture.Texture, nil, dst)
	if err != nil {
		return errors.Wrap(err, "failed to copy texture to window renderer")
	}

	return nil
}
