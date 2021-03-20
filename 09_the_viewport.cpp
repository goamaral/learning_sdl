#include "System.hpp"

int main(int argc, char** args) {
  int window_width = 640;
  int window_height = 480;

  System system;
  system.init();
  std::shared_ptr<Window> main_window_p = system.create_window("main", "SDL Tutorial", window_width, window_height);

  try {
    main_window_p->load_surface_from_png("resources/images/texture.png", "viewport");
  } catch (std::string) {
    return 1;
  }

  main_window_p->surface_to_texture("viewport");

  main_window_p->renderer_reset();
    SDL_Rect top_left_viewport = { 0, 0, window_width / 2, window_height / 2 };
    main_window_p->renderer_set_viewport(&top_left_viewport);
    main_window_p->render_texture("viewport");

    SDL_Rect top_right_viewport = { window_width / 2, 0, window_width / 2, window_height / 2 };
    main_window_p->renderer_set_viewport(&top_right_viewport);
    main_window_p->render_texture("viewport");

    SDL_Rect bottom_viewport = { 0, window_height / 2, window_width, window_height / 2 };
    main_window_p->renderer_set_viewport(&bottom_viewport);
    main_window_p->render_texture("viewport");
  main_window_p->renderer_render();

  system.delay(2000);

  return 0;
}