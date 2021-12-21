#include "System.hpp"

int main(int argc, char** args) {
  System system;
  system.init();
  system.init_png_img();
  std::shared_ptr<Window> main_window_p = system.create_window("main", "SDL Tutorial", 640, 480);

  try {
    main_window_p->load_surface_from_png("../resources/images/texture.png", "texture");
  } catch (std::string) {
    return 1;
  }

  main_window_p->surface_to_texture("texture");

  main_window_p->renderer_reset();
    main_window_p->render_texture("texture");
  main_window_p->renderer_render();

  system.delay(2000);

  return 0;
}