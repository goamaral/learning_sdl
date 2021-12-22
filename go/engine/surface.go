package engine

import (
	"fmt"

	srf "github.com/Goamaral/learning_sdl/surface"
)

var name_SurfaceMap map[string]srf.Surface

// Init surface map if not initialized
func initSurfaceMap() {
	if name_SurfaceMap == nil {
		name_SurfaceMap = map[string]srf.Surface{}
	}
}

// Save surface
func SaveSurface(name string, surface *srf.Surface) error {
	initSurfaceMap()
	_, exists := name_SurfaceMap[name]
	if exists {
		return fmt.Errorf("surface %s already exists", name)
	}

	name_SurfaceMap[name] = *surface

	return nil
}

// Get surface
func GetSurface(name string) (srf.Surface, error) {
	initSurfaceMap()
	surface, exists := name_SurfaceMap[name]
	if !exists {
		return srf.Surface{}, fmt.Errorf("surface %s does not exist", name)
	}

	return surface, nil
}

// Destroy surface
func DestroySurface(name string) error {
	initSurfaceMap()
	surface, exists := name_SurfaceMap[name]
	if !exists {
		return fmt.Errorf("surface %s does not exist", name)
	}

	surface.Free()
	delete(name_SurfaceMap, name)

	return nil
}

// Free surfaces
func FreeSurfaces() {
	initSurfaceMap()

	for name := range name_SurfaceMap {
		DestroySurface(name)
	}
}
