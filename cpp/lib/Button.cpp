#include "Button.hpp"

Button::Button(std::string default_texture, std::string hovered_texture, std::string selected_texture) {
  state = BUTTON_STATE_DEFAULT;
  texture_map = {
    {BUTTON_STATE_DEFAULT, default_texture},
    {BUTTON_STATE_HOVERED, hovered_texture},
    {BUTTON_STATE_SELECTED, selected_texture},
  };
}

Button::~Button() {}

// INSTANCE METHODS
bool Button::isInside(int x, int y) {
  return x >= position.x && x <= position.x + width && y >= position.y && y <= position.y + height;
}

// Consume mouse event and apply state changes (hover in/out, select/unselect)
void Button::consumeMouseEvent(SDL_Event event) {
  switch (event.type) {
  // On hover
  case SDL_MOUSEMOTION:
    // If button not selected, apply hover changes
    if (state != BUTTON_STATE_SELECTED) {
      // Hover in
      if (isInside(event.motion.x, event.motion.y)) {
        state = BUTTON_STATE_HOVERED;
      // Hover out
      } else if (state == BUTTON_STATE_HOVERED) {
        state = BUTTON_STATE_DEFAULT;
      }
    }
    break;

  // On click
  case SDL_MOUSEBUTTONUP:
    if (isInside(event.button.x, event.button.y)) {
      // If is selected, it becomes hovered
      if (state == BUTTON_STATE_SELECTED) {
        state = BUTTON_STATE_HOVERED;
      // If is not selected, it becomes selected
      } else {
        state = BUTTON_STATE_SELECTED;
      }
    }
    break;
  }
}

// GETTERS
std::string Button::currentTexture()
{
  return texture_map[state];
}