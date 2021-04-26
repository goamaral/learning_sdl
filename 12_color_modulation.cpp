#include "System.hpp"

#include <algorithm>

int main(int argc, char** args) {
  int window_width = 640;
  int window_height = 480;

  System system;
  system.init();
  std::shared_ptr<Window> main_window_p = system.create_window("main", "SDL Tutorial", window_width, window_height);

  try {
    main_window_p->load_surface_from_png("resources/images/colors.png", "colors");
  } catch (std::string) {
    return 1;
  }

  main_window_p->surface_to_texture("colors");

  Color color(255, 255, 255, 255);

  SDL_Event event;
  bool running = true;

  while (running) {
    if (SDL_PollEvent(&event)) {
      switch (event.type) {
        case SDL_QUIT:
          running = false;
          break;

        case SDL_KEYDOWN:
          switch(event.key.keysym.sym) {
            case SDLK_ESCAPE:
              running = false;
              break;

            case SDLK_q:
              color.red = std::min(255, color.red + 32);
              break;

            case SDLK_a:
              color.red = std::max(0, color.red - 32);
              break;

            case SDLK_w:
              color.green = std::min(255, color.green + 32);
              break;

            case SDLK_s:
              color.green = std::max(0, color.green - 32);
              break;

            case SDLK_e:
              color.blue = std::min(255, color.blue + 32);
              break;

            case SDLK_d:
              color.blue = std::max(0, color.blue - 32);
              break;
          }

          main_window_p->texture("colors")->set_color_modulation(color);
          break;
      }
    }

    main_window_p->renderer_reset();
      main_window_p->render_texture("colors");
    main_window_p->renderer_render();
  }

  return 0;
}