package engine

import "github.com/pkg/errors"

// Reset window renderer (draw color and viewport)
func (w *Window) Reset(color *Color) error {
	err := w.SetDrawColor(color)
	if err != nil {
		return errors.Wrap(err, "failed to reset window renderer draw color")
	}

	err = w.renderer.SetViewport(nil)
	if err != nil {
		return errors.Wrap(err, "failed to reset window renderer viewport")
	}

	err = w.renderer.Clear()
	if err != nil {
		return errors.Wrap(err, "failed to clear window renderer")
	}

	return nil
}

// Presents window renderer
func (w *Window) Present() {
	w.renderer.Present()
}

// Set window renderer draw color
func (w *Window) SetDrawColor(color *Color) error {
	err := w.renderer.SetDrawColor(color.Red, color.Green, color.Blue, color.Alpha)
	if err != nil {
		return errors.Wrap(err, "failed to set window renderer draw color")
	}

	return nil
}

// Set window renderer viewport
func (w *Window) SetViewport(viewport *Viewport) error {
	// Reset viewport if nil
	if viewport == nil {
		err := w.renderer.SetViewport(nil)
		if err != nil {
			return errors.Wrap(err, "failed to set default window renderer viewport")
		}

		return nil
	}

	// Set viewport
	err := w.renderer.SetViewport(&viewport.Rect)
	if err != nil {
		return errors.Wrap(err, "failed to set window renderer viewport")
	}

	return nil
}
