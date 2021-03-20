#include "System.hpp"

int main(int argc, char** args) {
  int window_width = 640;
  int window_height = 480;

  System system;
  system.init();
  std::shared_ptr<Window> main_window_p = system.create_window("main", "SDL Tutorial", window_width, window_height);

  main_window_p->renderer_reset();
    SDL_Rect render_rectangle = { window_width / 4, window_height / 4, window_width / 2, window_height / 2 };
    main_window_p->render_rectangle(&render_rectangle, Color("red"));

    render_rectangle = { window_width / 6, window_height / 6, window_width * 2 / 3, window_height * 2 / 3 };
    main_window_p->render_rectangle_outline(&render_rectangle, Color("green"));

    main_window_p->render_line(0, window_height / 2, window_width, window_height / 2, Color("blue"));

    for (int i = 0; i < window_height; i += 4) main_window_p->render_point(window_width / 2, i, Color("yellow"));
  main_window_p->renderer_render();

  system.delay(2000);

  return 0;
}