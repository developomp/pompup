#region nix
# https://wiki.archlinux.org/title/Nix#Desktop_integration
export XDG_DATA_DIRS=$XDG_DATA_DIRS:$HOME/.nix-profile/share
#endregion nix

#region jetbrainsToolbox
# Added by Toolbox App
export PATH="$PATH:$HOME/.local/share/JetBrains/Toolbox/scripts"
#endregion jetbrainsToolbox
