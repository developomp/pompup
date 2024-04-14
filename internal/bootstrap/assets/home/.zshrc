export PATH="$HOME/.local/bin/:$PATH"

#region oh my zsh
if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then
  source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"
fi

export ZSH="$HOME/.oh-my-zsh"
ZSH_THEME="powerlevel10k/powerlevel10k"
plugins=(
  fzf
  git
  zsh-interactive-cd
  zsh-syntax-highlighting
)

source $ZSH/oh-my-zsh.sh
[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh
#endregion oh my zsh

#region aliases & ENV VAR
alias c='clear'

alias e='eza --long --classify=always --color=always --icons=always --all --group-directories-first --binary --group --no-time --octal-permissions'

export EDITOR=vim
export VISUAL=vim
#endregion aliases & ENV VAR

#region PATH
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
export PATH="$PATH:$HOME/.cargo/bin"
export PATH="$PATH:$HOME/.local/bin"
#endregion PATH

#region pnpm
export PNPM_HOME="$HOME/.local/share/pnpm"
case ":$PATH:" in
  *":$PNPM_HOME:"*) ;;
  *) export PATH="$PNPM_HOME:$PATH" ;;
esac
#endregion pnpm

#region nvm
source /usr/share/nvm/init-nvm.sh
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
#endregion nvm

#region conda
# >>> conda initialize >>>
# !! Contents within this block are managed by 'conda init' !!
__conda_setup="$('/usr/bin/conda' 'shell.zsh' 'hook' 2> /dev/null)"
if [ $? -eq 0 ]; then
    eval "$__conda_setup"
else
    if [ -f "/usr/etc/profile.d/conda.sh" ]; then
        . "/usr/etc/profile.d/conda.sh"
    else
        export PATH="/usr/bin:$PATH"
    fi
fi
unset __conda_setup
# <<< conda initialize <<<
#endregion conda

#region atuin
# https://github.com/atuinsh/atuin
eval "$(atuin init zsh)"
#endregion atuin

#region fzf
# https://github.com/junegunn/fzf
eval "$(fzf --zsh)"
#endregion fzf
