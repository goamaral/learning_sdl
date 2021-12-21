#include "System.hpp"

int main(int argc, char** args) {
  int window_width = 640;
  int window_height = 480;

  System system;
  system.init();
  std::shared_ptr<Window> main_window_p = system.create_window("main", "SDL Tutorial", window_width, window_height);

  try {
    main_window_p->load_surface_from_png("../resources/images/background.png", "background");
    main_window_p->load_surface_from_png("../resources/images/player.png", "player");
  } catch (std::string) {
    return 1;
  }

  main_window_p->surface_to_texture("background");

  main_window_p->surface("player")->set_color_key(true, Color(0, 0xFF, 0xFF, 0xFF));
  std::shared_ptr<Texture> player_texture_p = main_window_p->surface_to_texture("player");

  main_window_p->renderer_reset();
    main_window_p->render_texture("background");
    main_window_p->render_texture("player", 240, 190);
  main_window_p->renderer_render();

  system.delay(2000);

  return 0;
}