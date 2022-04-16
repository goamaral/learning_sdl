package engine

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Viewport struct {
	sdl.Rect
	ID     uint32
	window *Window
}

// Render texture
func (v *Viewport) RenderTexture(texture *Texture, x int32, y int32) error {
	// Set viewport
	err := v.window.SetViewport(v)
	if err != nil {
		return errors.Wrap(err, "failed to set viewport")
	}

	// Render texture
	err = v.window.renderer.Copy(texture.Texture, nil, &sdl.Rect{X: x, Y: y, W: v.W, H: v.H})
	if err != nil {
		return errors.Wrap(err, "failed to copy texture to window renderer")
	}

	return nil
}
