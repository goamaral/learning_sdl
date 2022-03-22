package engine

import "github.com/pkg/errors"

// Reset window renderer (draw color and view port)
func (w *Window) Reset(color *Color) error {
	err := w.SetDrawColor(color)
	if err != nil {
		return errors.Wrap(err, "failed to rest window renderer")
	}

	err = w.renderer.SetViewport(nil)
	if err != nil {
		return errors.Wrap(err, "failed to rest window renderer")
	}

	err = w.renderer.Clear()
	if err != nil {
		return errors.Wrap(err, "failed to rest window renderer")
	}

	return nil
}

// Presents window renderer
func (w *Window) Present() {
	w.renderer.Present()
}

// Set window renderer draw color
func (w *Window) SetDrawColor(color *Color) error {
	err := w.renderer.SetDrawColor(color.red, color.green, color.blue, color.alpha)
	if err != nil {
		return errors.Wrap(err, "failed to set window renderer draw color")
	}

	return nil
}
