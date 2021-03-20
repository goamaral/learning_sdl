#include "Texture.hpp"

Texture::Texture(SDL_Texture* sdl_texture_p, int width, int height) {
  $sdl_p = sdl_texture_p;
  $width = width;
  $height = height;
}

Texture::~Texture() {
  if ($sdl_p) SDL_DestroyTexture($sdl_p);
}

// STATIC METHODS

// INSTANCE METHODS

// GETTERS
SDL_Texture* Texture::sdl_p() {
  return $sdl_p;
}

int Texture::width() {
  return $width;
}

int Texture::height() {
  return $height;
}