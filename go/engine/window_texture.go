package engine

import (
	"github.com/pkg/errors"
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
func (w *Window) RenderTexture(id uint64, viewportId uint64) error {
	if viewportId != 0 {
		err := w.SetViewportByID(viewportId)
		if err != nil {
			return errors.Wrap(err, "failed to set viewport")
		}
	}

	texture, err := w.GetTexture(id)
	if err != nil {
		return errors.Wrap(err, "failed to get texture")
	}

	err = w.renderer.Copy(texture.Texture, nil, nil)
	if err != nil {
		return errors.Wrap(err, "failed to copy texture to window renderer")
	}

	return nil
}
