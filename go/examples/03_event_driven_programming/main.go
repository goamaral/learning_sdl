/*
	Event Driven Programming
	https://lazyfoo.net/tutorials/SDL/03_event_driven_programming
*/

package main

import (
	"github.com/rs/zerolog/log"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/Goamaral/learning_sdl/engine"
)

func main() {
	// Init engine
	err := engine.Init()
	if err != nil {
		log.Error().Err(err).Msg("Failed to initialized engine")
		return
	}
	defer engine.Quit()

	// Create window
	_, err = engine.NewWindow(640, 480, false)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create window")
		return
	}

	// Run event loop
	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		event := getEvent()
		switch eventType := event.(type) {
		// Key pressed
		case *sdl.KeyboardEvent:
			switch eventType.Keysym.Sym {
			case sdl.K_ESCAPE:
				return false
			}
		}
		return true
	})
}
