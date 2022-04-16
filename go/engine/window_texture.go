package engine

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
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

// Render texture
func (w *Window) RenderTexture(texture *Texture, x int32, y int32) error {
	// Set viewport
	err := w.SetViewport(nil)
	if err != nil {
		return errors.Wrap(err, "failed to set viewport")
	}

	// Render texture
	err = w.renderer.Copy(texture.Texture, nil, &sdl.Rect{X: x, Y: y, W: texture.W, H: texture.H})
	if err != nil {
		return errors.Wrap(err, "failed to copy texture to window renderer")
	}

	return nil
}
