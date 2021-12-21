#include "System.hpp"

#include <algorithm>

int main(int argc, char** args) {
  int window_width = 640;
  int window_height = 480;

  System system;
  system.init();
  std::shared_ptr<Window> main_window_p = system.create_window("main", "SDL Tutorial", window_width, window_height);

  try {
    main_window_p->load_surface_from_png("../resources/images/arrow.png", "arrow");
  } catch (std::string) {
    return 1;
  }

  std::shared_ptr<Texture> arrow_texture_p = main_window_p->surface_to_texture("arrow");

  SDL_Event event;
  bool running = true;
  double degrees = 0;
  SDL_RendererFlip flip = SDL_FLIP_NONE;
  int x = (window_width - arrow_texture_p->width()) / 2;
  int y = (window_height - arrow_texture_p->height()) / 2;

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

            case SDLK_a:
              degrees -= 60;
              break;

            case SDLK_d:
              degrees += 60;
              break;

            case SDLK_q:
              flip = SDL_FLIP_HORIZONTAL;
              break;

            case SDLK_w:
              flip = SDL_FLIP_NONE;
              break;

            case SDLK_e:
              flip = SDL_FLIP_VERTICAL;
              break;
          }
          break;
      }
    }

    main_window_p->renderer_reset();
      main_window_p->render_texture("arrow", x, y, NULL, degrees, flip);
    main_window_p->renderer_render();
  }

  return 0;
}