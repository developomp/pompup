# pompup

pompup is a personal Arch Linux desktop setup utility tailor made for myself.

- [Gallery](./docs/gallery.md)

|                 Software | Choice                                                                                     |
| -----------------------: | :----------------------------------------------------------------------------------------- |
| Desktop Environment - 🚀 | [GNOME](https://www.gnome.org)                                                             |
|               Icons - 💎 | [Papirus](https://github.com/PapirusDevelopmentTeam/papirus-icon-theme)                    |
|           GTK theme - 🎨 | [WhiteSur](https://github.com/vinceliuice/WhiteSur-gtk-theme)                              |
|               Shell - 🐚 | [zsh](https://github.com/zsh-users/zsh) with [ohmyzsh](https://github.com/ohmyzsh/ohmyzsh) |
|            Terminal - 🖥️ | [alacritty](https://github.com/alacritty/alacritty)                                        |
|        File manager - 📂 | [Nautilus](https://gitlab.gnome.org/GNOME/nautilus)                                        |
|             Browser - 🌐 | [Brave](https://github.com/brave/brave-browser)                                            |
|   Text Editor & IDE - 📝 | [VSCodium](https://github.com/VSCodium/vscodium)                                           |

## Name?

The project's name was inspired by [rustup](https://github.com/rust-lang/rustup),
a toolchain installer for the Rust programming language. Pompup was originally
supposed to be written in Rust. I later decided to use Go instead, but the name
stuck.

## Usage

I can't think of any valid use cases of this software outside my computer,
but here are the instructions to use it:

1. Install base Arch Linux
   - https://wiki.archlinux.org/title/installation_guide
2. Create a non-root user and log in
   - give sudo permission
   - create home directory
3. Download pompup

   - using curl (included in archiso)

   ```
   curl pompup.developomp.com -Lo pompup
   ```

   - using wget

   ```
   wget pompup.developomp.com -O pompup
   ```

4. Run pompup

   ```
   ./pompup
   ```

## TODO

- GNOME extension custom repo
- https://github.com/navidys/tvxwidgets
- https://pkg.go.dev/github.com/rivo/tview#hdr-Widgets
