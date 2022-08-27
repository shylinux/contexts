#!/bin/bash

echo "" > ~/.hushlogin
export BASH_SILENCE_DEPRECATION_WARNING=1

export CTX_ROOT=${CTX_ROOT:=~/contexts}
export LC_ALL=en_US.utf-8

ish_sys_path_load
ish_sys_cli_prompt
ish_sys_cli_alias vi vim
ish_sys_cli_alias t "tmux attach"

