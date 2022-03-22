package engine

// Reset window renderer (draw color and view port)
func (w *Window) Reset(color Color) {
	err := w.renderer.SetDrawColor(color.red, color.green, color.blue, color.alpha)
	if err != nil {
		panic(err)
	}

	err = w.renderer.SetViewport(nil)
	if err != nil {
		panic(err)
	}

	err = w.renderer.Clear()
	if err != nil {
		panic(err)
	}
}

// Presents window renderer
func (w *Window) Present() {
	w.renderer.Present()
}
