#!/bin/bash

export LC_ALL=en_US.utf-8
export BASH_SILENCE_DEPRECATION_WARNING=1
touch ~/.hushlogin

# export CTX_ROOT=${CTX_ROOT:=~/contexts}
[ "$PWD" = "$HOME" ] && cd ~/contexts

export PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
ish_sys_path_load

ish_sys_cli_prompt
ish_sys_cli_alias vi vim
ish_sys_cli_alias t "tmux attach"

[ -f ~/.bash_temp ] && source ~/.bash_temp; rm ~/.bash_temp &>/dev/null
