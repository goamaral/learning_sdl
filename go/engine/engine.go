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
	sdl.Quit()
}

// Event loop
func EventLoop(eventHandler func(sdl.Event) bool) {
	for ProcessEvents(eventHandler) {
	}
}

// Process events
func ProcessEvents(eventHandler func(sdl.Event) bool) bool {
	shouldContinue := true

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		// Window closed
		case *sdl.QuitEvent:
			shouldContinue = false
		default:
			if eventHandler != nil {
				shouldContinue = shouldContinue && eventHandler(event)
			}
		}
	}

	return shouldContinue
}
