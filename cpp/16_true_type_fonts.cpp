#include "System.hpp"

int main(int argc, char** args) {
  int window_width = 640;
  int window_height = 480;

  System system;
  system.init();
  system.init_ttf_font();
  std::shared_ptr<Window> main_window_p = system.create_window("main", "SDL Tutorial", window_width, window_height);

  main_window_p->load_font_from_ttf("../resources/fonts/lazy.ttf", "lazy", 24);
  main_window_p->font_to_surface("lazy", "The quick brown fox jumps over the lazy dog");
  std::shared_ptr<Texture> lazy_texture_p = main_window_p->surface_to_texture("lazy");

  int x = (window_width - lazy_texture_p->width()) / 2;
  int y = (window_height - lazy_texture_p->height()) / 2;

  main_window_p->renderer_reset(Color("white"));
    main_window_p->render_texture("lazy", x, y);
  main_window_p->renderer_render();

  system.delay(5000);

  return 0;
}