#include "System.hpp"

#include <algorithm>

int main(int argc, char** args) {
  int window_width = 640;
  int window_height = 480;

  System system;
  system.init();
  std::shared_ptr<Window> main_window_p = system.create_window("main", "SDL Tutorial", window_width, window_height);

  try {
    main_window_p->load_surface_from_png("../resources/images/walking.png", "walking");
  } catch (std::string) {
    return 1;
  }

  main_window_p->surface_to_texture("walking");

  const int WALKING_ANIMATION_FRAMES = 4;
  SDL_Rect walking_frame_rects[WALKING_ANIMATION_FRAMES];

  walking_frame_rects[0].x = 0;
  walking_frame_rects[0].y = 0;
  walking_frame_rects[0].w = 64;
  walking_frame_rects[0].h = 205;

  walking_frame_rects[1].x = 64;
  walking_frame_rects[1].y = 0;
  walking_frame_rects[1].w = 64;
  walking_frame_rects[1].h = 205;
  
  walking_frame_rects[2].x = 128;
  walking_frame_rects[2].y = 0;
  walking_frame_rects[2].w = 64;
  walking_frame_rects[2].h = 205;

  walking_frame_rects[3].x = 196;
  walking_frame_rects[3].y = 0;
  walking_frame_rects[3].w = 64;
  walking_frame_rects[3].h = 205;

  SDL_Event event;
  bool running = true;
  int current_frame = 0;

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

    current_frame = current_frame % WALKING_ANIMATION_FRAMES;
    SDL_Rect* current_frame_rect_p = &walking_frame_rects[current_frame];
    int x = (window_width - current_frame_rect_p->w) / 2;
    int y = (window_height - current_frame_rect_p->h) / 2;

    main_window_p->renderer_reset();
      main_window_p->render_texture("walking", x, y, current_frame_rect_p);
    main_window_p->renderer_render();

    current_frame += 1;
  }

  return 0;
}