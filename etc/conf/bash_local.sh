#!/bin/bash

echo "" > ~/.hushlogin
export BASH_SILENCE_DEPRECATION_WARNING=1
export CTX_ROOT=${CTX_ROOT:=~/contexts}
export LC_ALL=en_US.utf-8
export PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
ish_sys_path_load
ish_sys_cli_prompt
ish_sys_cli_alias vi vim
ish_sys_cli_alias t "tmux attach"
[ "$PWD" = "$HOME" ] && cd ~/contexts
[ -f ~/.bash_temp ] && source ~/.bash_temp; rm ~/.bash_temp &>/dev/null
