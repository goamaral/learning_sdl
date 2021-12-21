#ifndef SYSTEM_H
#define SYSTEM_H

#include <string>
#include <memory>
#include <functional>
#include <unordered_map>

#include <SDL2/SDL.h>
#include <SDL2/SDL_assert.h>
#include <SDL2/SDL_image.h>
#include <SDL2/SDL_ttf.h>

#include "Window.hpp"

class System {
  public:
    System();
    ~System();

    // STATIC METHODS

    // INSTANCE METHODS
    // GENERAL
    void init();
    void init_png_img();
    void init_ttf_font();
    void delay(Uint32);

    // WINDOWS
    std::shared_ptr<Window> create_window(std::string, std::string, int, int);

    // GETTERS
    std::shared_ptr<Window> window(std::string);

    private:
      bool $sdl_inited;
      bool $sdl_img_inited;
      bool $sdl_png_img_inited;
      bool $sdl_ttf_font_inited;
      std::unordered_map<std::string, std::shared_ptr<Window>> $windows;
    };

#endif