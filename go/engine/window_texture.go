package engine

import (
	"github.com/pkg/errors"
)

// Get next texture id
// FUTURE: Make it thread safe
func (w *Window) GetNextTextureID() uint64 {
	nextID := w.lastTextureID + 1
	w.lastTextureID = nextID

	return nextID
}

// Save texture
func (w *Window) SaveTexture(texture *Texture) {
	texture.ID = w.GetNextSurfaceID()
	w.textureMap[texture.ID] = texture
}

// Get texture
func (w *Window) GetTexture(id uint64) (*Texture, error) {
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
func (w *Window) RenderTexture(id uint64) error {
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
