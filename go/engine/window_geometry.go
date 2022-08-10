package engine

import "github.com/pkg/errors"

// Render rectangle in window renderer
func (w *Window) RenderRectangle(rectangle *Rectangle) error {
	err := w.SetDrawColor(rectangle.Color)
	if err != nil {
		return errors.Wrap(err, "failed to set window renderer draw color")
	}

	switch rectangle.RenderMode {
	case GEOMETRY_RENDER_MODE_FILL:
		err = w.renderer.FillRect(&rectangle.Rect)
	case GEOMETRY_RENDER_MODE_OUTLINE:
		err = w.renderer.DrawRect(&rectangle.Rect)
	default:
		err = errors.Errorf("Unknown render mode %d", rectangle.RenderMode)
	}
	if err != nil {
		return errors.Wrap(err, "failed to render rectangle in window renderer")
	}

	return nil
}

// Render line in window renderer
func (w *Window) RenderLine(x1, y1, x2, y2 int32, color Color) error {
	err := w.SetDrawColor(color)
	if err != nil {
		return errors.Wrap(err, "failed to set window renderer draw color")
	}

	err = w.renderer.DrawLine(x1, y1, x2, y2)
	if err != nil {
		return errors.Wrap(err, "failed to render line in window renderer")
	}

	return nil
}

// Render point in window renderer
func (w *Window) RenderPoint(x, y int32, color Color) error {
	err := w.SetDrawColor(color)
	if err != nil {
		return errors.Wrap(err, "failed to set window renderer draw color")
	}

	err = w.renderer.DrawPoint(x, y)
	if err != nil {
		return errors.Wrap(err, "failed to render line in window renderer")
	}

	return nil
}
