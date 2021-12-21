#include "System.hpp"

#include <algorithm>

int main(int argc, char** args) {
  int window_width = 640;
  int window_height = 480;

  System system;
  system.init();
  std::shared_ptr<Window> main_window_p = system.create_window("main", "SDL Tutorial", window_width, window_height);

  try {
    main_window_p->load_surface_from_png("../resources/images/fadein.png", "fadein");
    main_window_p->load_surface_from_png("../resources/images/fadeout.png", "fadeout");
  } catch (std::string) {
    return 1;
  }

  main_window_p->surface_to_texture("fadein");
  main_window_p->surface_to_texture("fadeout");
  main_window_p->texture("fadeout")->set_blend_mode(SDL_BLENDMODE_BLEND);

  Uint8 alpha = 255;

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

            case SDLK_UP:
              alpha = std::min(255, alpha + 32);
              break;

            case SDLK_DOWN:
              alpha = std::max(0, alpha - 32);
              break;
          }

          main_window_p->texture("fadeout")->set_alpha_modulation(alpha);
          break;
      }
    }

    main_window_p->renderer_reset();
      main_window_p->render_texture("fadein");
      main_window_p->render_texture("fadeout");
    main_window_p->renderer_render();
  }

  return 0;
}