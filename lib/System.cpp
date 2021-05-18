#include "System.hpp"

System::System() {
  $sdl_inited = false;
  $sdl_img_inited = false;
}

System::~System() {
  $windows.clear();

  if ($sdl_ttf_font_inited) TTF_Quit();
  if ($sdl_img_inited) IMG_Quit();
  if ($sdl_inited) SDL_Quit();
}

// STATIC METHODS

// INSTANCE METHODS
// GENERAL
void System::init() {
  SDL_assert(SDL_Init(SDL_INIT_EVERYTHING) == 0);
  $sdl_inited = true;
}

void System::init_png_img() {
  if (!$sdl_png_img_inited) {
    SDL_assert((IMG_Init(IMG_INIT_PNG) & IMG_INIT_PNG) == IMG_INIT_PNG);
    $sdl_png_img_inited = true;
    $sdl_img_inited = true;
  }
}

void System::init_ttf_font() {
  if (!$sdl_ttf_font_inited) {
    SDL_assert(TTF_Init() == 0);
    $sdl_ttf_font_inited = true;
  }
}

void System::delay(Uint32 ms) {
  SDL_Delay(ms);
}

// WINDOWS
std::shared_ptr<Window> System::create_window(std::string key, std::string title, int width, int height) {
  std::shared_ptr<Window> window_p(new Window());

  window_p->init(title, width, height);
  $windows.insert({ key, window_p });

  return window_p;
}

// GETTERS
std::shared_ptr<Window> System::window(std::string window_key) {
  return $windows.at(window_key);
}