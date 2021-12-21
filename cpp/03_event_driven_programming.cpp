#include "System.hpp"

int main(int argc, char** args) {
  System system;
  system.init();
  system.create_window("main", "SDL Tutorial", 640, 480);

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
          }
          break;
      }
    }
  }

  return 0;
}