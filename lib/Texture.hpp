#ifndef TEXTURE_H
#define TEXTURE_H

#include <cstdio>
#include <string>

#include <SDL2/SDL.h>
#include <SDL2/SDL_assert.h>
#include <SDL2/SDL_log.h>
#include <SDL2/SDL_image.h>

#include "Surface.hpp"

class Texture {
  public:
      Texture(SDL_Texture*, int, int);
      ~Texture();

      // STATIC METHODS

      // INSTANCE METHODS

      // GETTERS
      SDL_Texture* sdl_p();
      int width();
      int height();

  private:
    SDL_Texture* $sdl_p;
    int $width;
    int $height;
};

#endif