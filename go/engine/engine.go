package engine

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Init engine
func Init() error {
	if err := ttf.Init(); err != nil {
		return errors.Wrap(err, "failed to initialize sdl ttf api")
	}
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return errors.Wrap(err, "failed to initialize sdl")
	}
	return nil
}

// Quit engine
func Quit() {
	destroyViewports()
	textureManager.Destroy()
	surfaceManager.Destroy()
	rendererManager.Destroy()
	windowManager.Destroy()
	sdl.Quit()
}

// Event loop
func EventLoop(renderHandler func(func() sdl.Event) bool) {
	if renderHandler != nil {
		running := true
		ok := true
		for running && ok {
			running = renderHandler(func() sdl.Event {
				var sdlEvent sdl.Event
				ok, sdlEvent = getEvent()

				return sdlEvent
			})
		}
	}
}

// Get events
func getEvent() (bool, sdl.Event) {
	sdlEvent := sdl.PollEvent()
	if sdlEvent != nil {
		if sdlEvent.GetType() == sdl.QUIT {
			return false, nil
		}
	}

	return true, sdlEvent
}
