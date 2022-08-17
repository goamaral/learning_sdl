package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Init engine
func Init() error {
	return sdl.Init(sdl.INIT_EVERYTHING)
}

// Quit engine
func Quit() {
	destroyViewports()
	destroySurfaces()
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
