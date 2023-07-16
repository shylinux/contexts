#!/bin/sh

export ctx_shy=${ctx_shy:=https://shylinux.com}
if [ -f $PWD/.ish/plug.sh ]; then source $PWD/.ish/plug.sh; elif [ -f $HOME/.ish/plug.sh ]; then source $HOME/.ish/plug.sh; else
	temp=$(mktemp); if curl -h &>/dev/null; then curl -o $temp -fsSL $ctx_shy; else wget -O $temp -q $ctx_shy; fi; source $temp intshell
fi; require conf.sh; require miss.sh

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
ish_miss_prepare go-git
ish_miss_prepare go-qrcode
ish_miss_prepare websocket
ish_miss_prepare webview
ish_miss_prepare matrix
ish_miss_prepare word-dict
ish_miss_prepare wubi-dict

ish_miss_make
