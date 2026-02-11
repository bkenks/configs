if status is-interactive
# Commands to run in interactive sessions can go here
end

set -gx PATH "$HOME/.local/bin" $PATH
set -gx PATH "/usr/local/go/bin" $PATH

# source $HOME/secrets.fish
starship init fish | source
/home/linuxbrew/.linuxbrew/bin/brew shellenv | source
