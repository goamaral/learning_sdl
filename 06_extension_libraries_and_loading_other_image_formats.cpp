#include "System.hpp"

int main(int argc, char** args) {
  System system;
  system.init();
  system.init_png_img();
  std::shared_ptr<Window> main_window_p = system.create_window("main", "SDL Tutorial", 640, 480);

  try {
    main_window_p->load_surface_from_png("resources/images/png_loaded.png", "png_loaded");
  } catch (std::string) {
    return 1;
  }

  main_window_p->render_surface("png_loaded");
  system.delay(2000);

  return 0;
}