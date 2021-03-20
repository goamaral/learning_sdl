#ifndef WINDOW_H
#define WINDOW_H

#include <string>
#include <unordered_map>

#include <SDL2/SDL.h>
#include <SDL2/SDL_assert.h>
#include <SDL2/SDL_image.h>

#include "Texture.hpp"
#include "Surface.hpp"
#include "Color.hpp"

class Window {
  public:
    Window();
    ~Window();

    // STATIC METHODS

    // INSTANCE METHODS
    // GENERAL
    void init(std::string, int, int);

    // RENDERER
    void renderer_render();
    void renderer_set_draw_color(Color);
    void renderer_set_viewport(SDL_Rect*);
    void renderer_clear();
    void renderer_reset();

    // SURFACES
    std::shared_ptr<Surface> load_surface_from_bmp(std::string, std::string);
    std::shared_ptr<Surface> load_surface_from_png(std::string, std::string);
    void render_surface(std::string, bool = false);

    // TEXTURES
    std::shared_ptr<Texture> surface_to_texture(std::string);
    void render_texture(std::string, SDL_Rect* = NULL);

    // GEOMETRIES
    void render_rectangle(SDL_Rect*, Color);
    void render_rectangle_outline(SDL_Rect*, Color);
    void render_line(int, int, int, int, Color);
    void render_point(int, int, Color);

    // GETTERS
    std::shared_ptr<Surface> surface(std::string);
    std::shared_ptr<Texture> texture(std::string);

  private:
    SDL_Window* $sdl_p;
    SDL_Renderer* $sdl_renderer_p;
    SDL_Surface* $sdl_surface_p;
    std::string $title;
    int $width;
    int $height;
    std::unordered_map<std::string, std::shared_ptr<Surface>> $surfaces;
    std::unordered_map<std::string, std::shared_ptr<Texture>> $textures;
};

#endif