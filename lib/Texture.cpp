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
void Texture::set_color_modulation(Color color) {
  int result = SDL_SetTextureColorMod($sdl_p, color.red, color.green, color.blue);
  SDL_assert(result == 0);
}

void Texture::set_blend_mode(SDL_BlendMode blend_mode) {
  int result = SDL_SetTextureBlendMode($sdl_p, blend_mode);
  SDL_assert(result == 0);
}

void Texture::set_alpha_modulation(Uint8 alpha) {
  int result = SDL_SetTextureAlphaMod($sdl_p, alpha);
  SDL_assert(result == 0);
}

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