/*
	Getting an Image on the Screen
	https://lazyfoo.net/tutorials/SDL/02_getting_an_image_on_the_screen
*/

package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// Init SDL
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// Create window
	window, err := sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// Load surface from bmp
	surface, err := sdl.LoadBMP("../resources/images/hello_world.bmp")
	if err != nil {
		panic(err)
	}
	defer surface.Free()

	// Get window surface
	windowSurface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	// Optimize surface
	optimezedSurface, err := surface.Convert(windowSurface.Format, 0)
	if err != nil {
		fmt.Println("Failed to optimize surface")
	} else {
		surface = optimezedSurface
		defer optimezedSurface.Free()
	}

	// Copy surface
	err = surface.Blit(nil, windowSurface, nil)
	if err != nil {
		panic(err)
	}

	// Update surface
	window.UpdateSurface()

	// Main loop
	running := true
	for running {
		event := sdl.PollEvent()
		if event != nil {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
	}
}
