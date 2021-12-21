#include <System.hpp>

int main(int argc, char** args) {
  System system;
  system.init();
  std::shared_ptr<Window> main_window_p = system.create_window("main", "SDL Tutorial", 640, 480);

  try {
    main_window_p->load_surface_from_bmp("../resources/images/hello_world.bmp", "hello_world");
  } catch (std::string) {
    return 1;
  }

  main_window_p->render_surface("hello_world");
  system.delay(2000);

  return 0;
}