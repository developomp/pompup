export ZSH="/home/pomp/.oh-my-zsh"
ZSH_THEME="powerlevel10k/powerlevel10k"
plugins=(
  git
  zsh-interactive-cd
  zsh-syntax-highlighting
)

source $ZSH/oh-my-zsh.sh
source /usr/share/nvm/init-nvm.sh

export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
export PATH="$PATH:$HOME/.yarn/bin:$HOME/.local/bin"
export PATH="$PATH:$HOME/.deno/bin"

export EDITOR=vim
export VISUAL=vim
