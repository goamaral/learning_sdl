package engine

var (
	COLOR_BLACK   = Color{0, 0, 0, 255}
	COLOR_WHITE   = Color{255, 255, 255, 255}
	COLOR_RED     = Color{255, 0, 0, 255}
	COLOR_GREEN   = Color{0, 255, 0, 255}
	COLOR_BLUE    = Color{0, 0, 255, 255}
	COLOR_YELLOW  = Color{255, 255, 0, 255}
	COLOR_MAGENTA = Color{255, 0, 255, 255}
	COLOR_CYAN    = Color{0, 255, 255, 255}
)

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha uint8
}
