#!/bin/bash

if [ -f $PWD/.ish/plug.sh ]; then source $PWD/.ish/plug.sh; elif [ -f $HOME/.ish/plug.sh ]; then source $HOME/.ish/plug.sh; else
	ctx_temp=$(mktemp); if curl -h &>/dev/null; then curl -o $ctx_temp -fsSL https://shylinux.com; else wget -O $ctx_temp -q http://shylinux.com; fi; source $ctx_temp intshell
fi

require sys/cli/file.sh
ish_sys_path_load

require miss.sh
ish_miss_prepare_compile
ish_miss_prepare_develop
ish_miss_prepare_operate

# ish_miss_prepare wubi-dict
# ish_miss_prepare word-dict

ish_miss_prepare linux-story
ish_miss_prepare nginx-story
ish_miss_prepare golang-story
ish_miss_prepare redis-story
ish_miss_prepare mysql-story
ish_miss_prepare release

ish_miss_prepare_intshell
ish_miss_prepare_contexts
ish_miss_prepare_icebergs
ish_miss_prepare_toolkits
ish_miss_prepare_volcanos
ish_miss_prepare_learning

ish_miss_make; if [ -n "$*" ]; then ish_miss_serve "$@"; fi

ish_sys_link_create ~/.bash_local $PWD/etc/conf/bash_local.sh
ish_sys_link_create ~/.vim_local.vim $PWD/etc/conf/vim_local.vim
require sys/cli/cli.sh
ish_sys_cli_prepare

require dev/vim/vim.sh
ish_dev_vim_prepare

if tmux -V; then
    require dev/tmux/tmux.sh
    ish_dev_tmux_prepare
    ish_miss_prepare_session miss miss
else
    ish_miss_serve_log
fi

