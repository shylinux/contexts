#!/bin/sh

if [ -f $PWD/.ish/plug.sh ]; then source $PWD/.ish/plug.sh; elif [ -f $HOME/.ish/plug.sh ]; then source $HOME/.ish/plug.sh; else
	temp=$(mktemp); if curl -h &>/dev/null; then curl -o $temp -fsSL $ctx_dev; else wget -O $temp -q $ctx_dev; fi; source $temp intshell
fi

require miss.sh
ish_miss_prepare_compile
ish_miss_prepare_develop
ish_miss_prepare_project

ish_miss_prepare_contexts
ish_miss_prepare_intshell
ish_miss_prepare_learning
ish_miss_prepare_volcanos
ish_miss_prepare_toolkits
ish_miss_prepare_icebergs
ish_miss_prepare release
ish_miss_prepare icons

# ish_miss_prepare websocket
# ish_miss_prepare go-qrcode
ish_miss_prepare go-git
ish_miss_prepare matrix

_prepare_ttc() {
	ish_sys_cli_prepare; ish_dev_tmux_prepare; ish_dev_git_prepare; ish_dev_vim_prepare # ish_dev_vim_plug_prepare
	ish_sys_link_create ~/.bash_local.sh $PWD/etc/conf/bash_local.sh
	ish_sys_link_create ~/.vim_local.vim $PWD/etc/conf/vim_local.vim
	if tmux -V; then ish_miss_prepare_session miss miss; else ish_miss_serve_log; fi
}
ish_miss_make; if [ -n "$*" ]; then ish_miss_serve "$@"; else _prepare_ttc; fi
