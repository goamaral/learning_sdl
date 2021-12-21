#ifndef BUTTON_H
#define BUTTON_H

#include <string>
#include <unordered_map>

#include <SDL2/SDL.h>

enum ButtonState
{
  BUTTON_STATE_DEFAULT,
  BUTTON_STATE_HOVERED,
  BUTTON_STATE_SELECTED
};

class Button {
  public:
    SDL_Point position;
    int width, height;
    ButtonState state;
    std::unordered_map<ButtonState, std::string> texture_map;

    Button(std::string, std::string, std::string); // TODO
    ~Button(); // TODO

    // INSTANCE METHODS
    bool isInside(int, int);
    void consumeMouseEvent(SDL_Event);

    // GETTERS
    std::string currentTexture(); // TODO

    private:
};

#endif