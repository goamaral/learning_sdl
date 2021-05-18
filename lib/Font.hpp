

#ifndef FONT_H
#define FONT_H

#include <SDL2/SDL.h>
#include <SDL2/SDL_ttf.h>

class Font {
  public:
    Font(TTF_Font*);
    ~Font();

    // INSTANCE METHODS
    // GETTERS
    TTF_Font* sdl_p();

    private:
      TTF_Font* $sdl_p;
};

#endif