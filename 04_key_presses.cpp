#include "System.hpp"

int main(int argc, char** args) {
  System system;
  system.init();
  std::shared_ptr<Window> main_window_p = system.create_window("main", "SDL Tutorial", 640, 480);

  try {
    main_window_p->load_surface_from_bmp("resources/images/default.bmp", "default");
    main_window_p->load_surface_from_bmp("resources/images/up.bmp", "up");
    main_window_p->load_surface_from_bmp("resources/images/down.bmp", "down");
    main_window_p->load_surface_from_bmp("resources/images/left.bmp", "left");
    main_window_p->load_surface_from_bmp("resources/images/right.bmp", "right");
  } catch (std::string) {
    return 1;
  }

  main_window_p->render_surface("default");

  SDL_Event event;
  bool running = true;

  while (running) {
    SDL_PollEvent(&event);

    switch (event.type) {
      case SDL_QUIT:
        running = false;
        break;

      case SDL_KEYDOWN:
        switch(event.key.keysym.sym) {
          case SDLK_ESCAPE:
            running =  false;
            break;

          case SDLK_UP:
            main_window_p->render_surface("up");
            break;

          case SDLK_DOWN:
            main_window_p->render_surface("down");
            break;

          case SDLK_LEFT:
            main_window_p->render_surface("left");
            break;

          case SDLK_RIGHT:
            main_window_p->render_surface("right");
            break;

          default:
            main_window_p->render_surface("default");
            break;
        }
      break;
    }
  }

  return 0;
}