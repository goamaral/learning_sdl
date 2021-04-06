#include "System.hpp"

int main(int argc, char** args) {
  int window_width = 640;
  int window_height = 480;

  System system;
  system.init();
  std::shared_ptr<Window> main_window_p = system.create_window("main", "SDL Tutorial", window_width, window_height);

  try {
    main_window_p->load_surface_from_png("resources/images/dots.png", "dots");
  } catch (std::string) {
    return 1;
  }

  main_window_p->surface_to_texture("dots");

  main_window_p->renderer_reset();
    SDL_Rect top_left_source_rectangle = { 0, 0, 100, 100 };
    main_window_p->render_texture("dots", 0, 0, &top_left_source_rectangle);

    SDL_Rect top_right_source_rectangle = { 100, 0, 100, 100 };
    main_window_p->render_texture("dots", window_width - 100, 0, &top_right_source_rectangle);

    SDL_Rect bottom_left_source_rectangle = { 0, 100, 100, 100 };
    main_window_p->render_texture("dots", 0, window_height - 100, &bottom_left_source_rectangle);

    SDL_Rect bottom_right_source_rectangle = { 100, 100, 100, 100 };
    main_window_p->render_texture("dots", window_width - 100, window_height - 100, &bottom_right_source_rectangle);
  main_window_p->renderer_render();

  system.delay(2000);

  return 0;
}