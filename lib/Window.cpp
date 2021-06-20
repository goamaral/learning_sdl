#include "Window.hpp"

Window::Window() {
  $sdl_p = NULL;
  $sdl_renderer_p = NULL;
  $sdl_surface_p = NULL;
}

Window::~Window() {
  $surfaces.clear();
  $textures.clear();
  $fonts.clear();

  if ($sdl_surface_p) SDL_FreeSurface($sdl_surface_p);
  if ($sdl_renderer_p) SDL_DestroyRenderer($sdl_renderer_p);
  if ($sdl_p) SDL_DestroyWindow($sdl_p);
}

// STATIC METHODS

// INSTANCE METHODS
// GENERAL
void Window::init(std::string title, int width, int height) {
  $title = title;
  $width = width;
  $height = height;

  $sdl_p = SDL_CreateWindow($title.c_str(), SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED, $width, $height, SDL_WINDOW_SHOWN);
  SDL_assert($sdl_p != NULL);

  $sdl_renderer_p = SDL_CreateRenderer($sdl_p, -1, SDL_RENDERER_ACCELERATED | SDL_RENDERER_PRESENTVSYNC); 
  SDL_assert($sdl_renderer_p != NULL);

  int result = SDL_SetRenderDrawColor($sdl_renderer_p, 0xFF, 0xFF, 0xFF, 0xFF);
  SDL_assert(result == 0);

  $sdl_surface_p = SDL_GetWindowSurface($sdl_p);
  SDL_assert($sdl_surface_p != NULL);
}

// RENDERER - Future Renderer class
void Window::renderer_render() {
  SDL_RenderPresent($sdl_renderer_p);
}

void Window::renderer_set_draw_color(Color color) {
  int result = SDL_SetRenderDrawColor($sdl_renderer_p, color.red, color.green, color.blue, color.alpha);
  SDL_assert(result == 0);
}

void Window::renderer_set_viewport(SDL_Rect* render_rectangle_p) {
  int result = SDL_RenderSetViewport($sdl_renderer_p, render_rectangle_p);
  SDL_assert(result == 0);
}

void Window::renderer_clear() {
  int result = SDL_RenderClear($sdl_renderer_p);
  SDL_assert(result == 0);
}

void Window::renderer_reset(Color color) {
  renderer_set_draw_color(color);
  renderer_set_viewport(NULL);
  renderer_clear();
}

// SURFACES
std::shared_ptr<Surface> Window::load_surface_from_bmp(std::string image_location, std::string key) {
  std::shared_ptr<Surface> surface_p(Surface::load_from_bmp(image_location, $sdl_surface_p->format));
  $surfaces.insert({ key, surface_p });

  return surface_p;
}

std::shared_ptr<Surface> Window::load_surface_from_png(std::string image_location, std::string key) {
  std::shared_ptr<Surface> surface_p(Surface::load_from_png(image_location, $sdl_surface_p->format));
  $surfaces.insert({key, surface_p});

  return surface_p;
}

std::shared_ptr<Surface> Window::font_to_surface(std::string key, std::string text, Color color) {
  std::shared_ptr<Font> font_p = $fonts.at(key);

  SDL_Surface* sdl_surface_p = TTF_RenderText_Solid(font_p->sdl_p(), text.c_str(), color.sdl());
  if (sdl_surface_p == NULL) {
    const char* error_message = SDL_GetError();
    SDL_LogCritical(SDL_LOG_CATEGORY_ERROR, "%s", error_message);
    throw std::string(error_message);
  }

  std::shared_ptr<Surface> surface_p(new Surface(sdl_surface_p));
  $surfaces.insert({key, surface_p});

  return surface_p;
}

void Window::render_surface(std::string surface_key, bool scaled) {
  std::shared_ptr<Surface> surface_p = $surfaces.at(surface_key);

  if (scaled) {
    Surface::copy_scaled(surface_p->sdl_p(), $sdl_surface_p);
  } else {
    Surface::copy(surface_p->sdl_p(), $sdl_surface_p);
  }

  int result = SDL_UpdateWindowSurface($sdl_p);
  SDL_assert(result == 0);
}

// FONTS
std::shared_ptr<Font> Window::load_font_from_ttf(std::string font_location, std::string key, int font_size) {
  TTF_Font* sdl_font_p = TTF_OpenFont(font_location.c_str(), font_size);

  if (sdl_font_p == NULL) {
    const char* error_message = SDL_GetError();
    SDL_LogCritical(SDL_LOG_CATEGORY_ERROR, "%s", error_message);
    throw std::string(error_message);
  }

  std::shared_ptr<Font> font_p(new Font(sdl_font_p));
  $fonts.insert({key, font_p});

  return font_p;
}

// TEXTURES
std::shared_ptr<Texture> Window::surface_to_texture(std::string key) {
  std::shared_ptr<Surface> surface_p = $surfaces.at(key);
  SDL_Texture* sdl_texture_p = SDL_CreateTextureFromSurface($sdl_renderer_p, surface_p->sdl_p());
  SDL_assert(sdl_texture_p != NULL);

  std::shared_ptr<Texture> texture_p(new Texture(sdl_texture_p, surface_p->width(), surface_p->height()));
  $textures.insert({key, texture_p});

  return texture_p;
}

void Window::render_texture(std::string key, int x, int y, SDL_Rect* source_rectangle_p, double angle, SDL_RendererFlip flip) {
  std::shared_ptr<Texture> texture_p = $textures.at(key);
  SDL_Rect destination_rectangle = { x, y, texture_p->width(), texture_p->height() };

  if (source_rectangle_p != NULL) {
    destination_rectangle.w = source_rectangle_p->w;
    destination_rectangle.h = source_rectangle_p->h;
  }

  int result = SDL_RenderCopyEx($sdl_renderer_p, texture_p->sdl_p(), source_rectangle_p, &destination_rectangle, angle, NULL, flip);
  SDL_assert(result == 0);
}

// GEOMETRIES
void Window::render_rectangle(SDL_Rect* render_rectangle_p, Color color) {
  renderer_set_draw_color(color);
  int result = SDL_RenderFillRect($sdl_renderer_p, render_rectangle_p);
  SDL_assert(result == 0);
}

void Window::render_rectangle_outline(SDL_Rect* render_rectangle_p, Color color) {
  renderer_set_draw_color(color);
  int result = SDL_RenderDrawRect($sdl_renderer_p, render_rectangle_p);
  SDL_assert(result == 0);
}

void Window::render_line(int x1, int y1, int x2, int y2, Color color) {
  renderer_set_draw_color(color);
  int result = SDL_RenderDrawLine($sdl_renderer_p, x1, y1, x2, y2);
  SDL_assert(result == 0);
}

void Window::render_point(int x, int y, Color color) {
  renderer_set_draw_color(color);
  int result = SDL_RenderDrawPoint($sdl_renderer_p, x, y);
  SDL_assert(result == 0);
}

// GETTERS
std::shared_ptr<Surface> Window::surface(std::string key) {
  return $surfaces.at(key);
}

std::shared_ptr<Texture> Window::texture(std::string key) {
  return $textures.at(key);
}

std::shared_ptr<Font> Window::font(std::string key) {
  return $fonts.at(key);
}