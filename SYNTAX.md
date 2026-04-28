# Inkstamp Syntax Reference


## Foreground (16-color)

| Command | Effect |
|---------|--------|
| `fg=black` | Black text |
| `fg=red` | Red text |
| `fg=green` | Green text |
| `fg=yellow` | Yellow text |
| `fg=blue` | Blue text |
| `fg=magenta` | Magenta text |
| `fg=cyan` | Cyan text |
| `fg=white` | White text |
| `fg=darkgray` | Dark gray text |
| `fg=lred` | Light red text |
| `fg=lgreen` | Light green text |
| `fg=lyellow` | Light yellow text |
| `fg=lblue` | Light blue text |
| `fg=lmagenta` | Light magenta text |
| `fg=lcyan` | Light cyan text |
| `fg=lwhite` | Light white text |

---


## Background (16-color)
| Command | Effect |
|---------|--------|
| `bg=black` | Black background |
| `bg=red` | Red background |
| `bg=green` | Green background |
| `bg=yellow` | Yellow background |
| `bg=blue` | Blue background |
| `bg=magenta` | Magenta background |
| `bg=cyan` | Cyan background |
| `bg=white` | White background |
| `bg=darkgray` | Dark gray background |
| `bg=lred` | Light red background |
| `bg=lgreen` | Light green background |
| `bg=lyellow` | Light yellow background |
| `bg=lblue` | Light blue background |
| `bg=lmagenta` | Light magenta background |
| `bg=lcyan` | Light cyan background |
| `bg=lwhite` | Light white background |

---

## Hex Colors
| Command | Effect |
|---------|--------|
| `fg=#RRGGBB` | Hex color for foreground |
| `bg=#RRGGBB` | Hex color for background |

Example: `[fg=#FF5733]Orange text[reset]`

---

## RGB Colors
| Command | Effect |
|---------|--------|
| `fg=rgb(R,G,B)` | RGB color for foreground |
| `bg=rgb(R,G,B)` | RGB color for background |

Example: `[fg=rgb(255,105,180)]Hot pink[reset]`

---

## 256-Color Palette
| Command | Effect |
|---------|--------|
| `fg=#NNN` | 256-color (0-255) for foreground |
| `bg=#NNN` | 256-color (0-255) for background |

Example: `[fg=214]Orange text[reset]`

---


## Styles
| Command | Effect |
|---------|--------|
| `bold` | Bold/bright text |
| `dim` | Dim/faint text |
| `italic` | Italic text |
| `underline=single` | Single underlined text |
| `underline=double` | Double underlined text |
| `blink=slow` | Slow blinking text |
| `blink=fast` | Fast blinking text |
| `reverse` | Reverse video (swap foreground and background colors) |
| `hidden` | Hidden text |
| `strike` | Strikethrough text |

---

## Resets
| Command | Effect |
|---------|--------|
| `reset` | Reset all colors and styles |
| `fg=reset` | Reset foreground color only |
| `bg=reset` | Reset background color only |
| `bold=reset` | Reset bold style only |
| `dim=reset` | Reset dim style only |
| `italic=reset` | Reset italic style only |
| `underline=reset` | Reset underline style only |
| `blink=reset` | Reset blink style only |
| `reverse=reset` | Reset reverse style only |
| `hidden=reset` | Reset hidden style only |
| `strike=reset` | Reset strikethrough style only |

---


## Placeholders
| Syntax | Effect |
|---------|--------|
| `[0]` | Insert first argument |
| `[1]` | Insert second argument |
| `[999]` | Insert 1000 argument |

Placeholders work from `[0]` through `[999]`.

---

## Padding And Alignment
| Syntax | Alignment | Fill | Example |
|---------|--------|--------|--------|
| `[0:<20]` | Left | Spaces | `"value               "` |
| `[0:<20:.]` | Left | Dots | `"value..............."` |
| `[0:>10]` | Right | Spaces | `"     value"` |
| `[0:>10:-]` | Right | Dashes | `"-----value"` |
| `[0:^15]` | Center | Spaces | `"     value     "` |
| `[0:^15:=]` | Center | Equals | `"=====value====="` |

---
