package engine

import "github.com/veandco/go-sdl2/sdl"

// Init engine
func Init() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
}

// Quit engine
func Quit() {
	sdl.Quit()
}

// Main loop
func Loop(handler func(sdl.Event) bool) {
	// Main loop
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			// Window closed
			case *sdl.QuitEvent:
				running = false
			default:
				running = handler(event)
			}
		}
	}
}
