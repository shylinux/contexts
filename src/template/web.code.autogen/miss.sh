#! /bin/sh

export ctx_shy=${ctx_shy:=https://shylinux.com}
if [ -f $PWD/.ish/plug.sh ]; then source $PWD/.ish/plug.sh; elif [ -f $HOME/.ish/plug.sh ]; then source $HOME/.ish/plug.sh; else
	temp=$(mktemp); if curl -h &>/dev/null; then curl -o $temp -fsSL $ctx_shy; else wget -O $temp -q $ctx_shy; fi; source $temp intshell
fi; require conf.sh; require miss.sh

ish_miss_prepare_compile
ish_miss_prepare_develop
ish_miss_prepare_project

ish_miss_make
