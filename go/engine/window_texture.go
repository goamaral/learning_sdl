package engine

import (
	"github.com/pkg/errors"
)

// Save texture
func (w *Window) SaveTexture(texture *Texture) {
	texture.ID = w.GetNextResourceID()
	w.textures[texture.ID] = *texture
}

// Get texture
func (w *Window) GetTexture(id uint32) (Texture, error) {
	texture, exists := w.textures[id]
	if !exists {
		return texture, errors.Errorf("texture %d not found", id)
	}

	return texture, nil
}

// Destroy texture
func (w *Window) DestroyTexture(id uint32) error {
	texture, exists := w.textures[id]
	if !exists {
		return errors.Errorf("texture %d not found", id)
	}

	delete(w.textures, id)
	return texture.Destroy()
}
