CXX = g++
CXXFLAGS = -std=c++11 -Wall
LDLIBS = -L./lib -lEngine -I./lib -lSDL2 -lSDL2_image -lSDL2_ttf

all: 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17

engine:
	(cd ./lib && make)

02: engine
	$(CXX) 02_getting_an_image_on_the_screen.cpp -o 02_getting_an_image_on_the_screen $(LDLIBS) $(CXXFLAGS)

03: engine
	$(CXX) 03_event_driven_programming.cpp -o 03_event_driven_programming $(LDLIBS) $(CXXFLAGS)

04: engine
	$(CXX) 04_key_presses.cpp -o 04_key_presses $(LDLIBS) $(CXXFLAGS)

05: engine
	$(CXX) 05_optimized_surface_loading_and_soft_stretching.cpp -o 05_optimized_surface_loading_and_soft_stretching $(LDLIBS) $(CXXFLAGS)

06: engine
	$(CXX) 06_extension_libraries_and_loading_other_image_formats.cpp -o 06_extension_libraries_and_loading_other_image_formats $(LDLIBS) $(CXXFLAGS)

07: engine
	$(CXX) 07_texture_loading_and_rendering.cpp -o 07_texture_loading_and_rendering $(LDLIBS) $(CXXFLAGS)

08: engine
	$(CXX) 08_geometry_rendering.cpp -o 08_geometry_rendering $(LDLIBS) $(CXXFLAGS)

09: engine
	$(CXX) 09_the_viewport.cpp -o 09_the_viewport $(LDLIBS) $(CXXFLAGS)

10: engine
	$(CXX) 10_color_keying.cpp -o 10_color_keying $(LDLIBS) $(CXXFLAGS)

11: engine
	$(CXX) 11_clip_rendering_and_sprite_sheets.cpp -o 11_clip_rendering_and_sprite_sheets $(LDLIBS) $(CXXFLAGS)

12: engine
	$(CXX) 12_color_modulation.cpp -o 12_color_modulation $(LDLIBS) $(CXXFLAGS)

13: engine
	$(CXX) 13_alpha_blending.cpp -o 13_alpha_blending $(LDLIBS) $(CXXFLAGS)

14: engine
	$(CXX) 14_animated_sprites_and_vsync.cpp -o 14_animated_sprites_and_vsync $(LDLIBS) $(CXXFLAGS)

15: engine
	$(CXX) 15_rotation_and_flipping.cpp -o 15_rotation_and_flipping $(LDLIBS) $(CXXFLAGS)

16: engine
	$(CXX) 16_true_type_fonts.cpp -o 16_true_type_fonts $(LDLIBS) $(CXXFLAGS)

17: engine
	$(CXX) 17_mouse_events.cpp -o 17_mouse_events $(LDLIBS) $(CXXFLAGS)