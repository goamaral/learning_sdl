#include "Font.hpp"

Font::Font(TTF_Font* sdl_font_p) {
  $sdl_p = sdl_font_p;
}

Font::~Font() {
  TTF_CloseFont($sdl_p);
}

// INSTANCE METHODS
// GETTERS
TTF_Font* Font::sdl_p() {
  return $sdl_p;
}