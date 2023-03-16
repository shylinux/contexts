#! /bin/sh

if [ -f $PWD/.ish/plug.sh ]; then source $PWD/.ish/plug.sh; elif [ -f $HOME/.ish/plug.sh ]; then source $HOME/.ish/plug.sh; else
	export ctx_dev=${ctx_dev:="https://shylinux.com"}
	temp=$(mktemp); if curl -h &>/dev/null; then curl -o $temp -fsSL $ctx_dev; else wget -O $temp -q $ctx_dev; fi; source $temp intshell
fi; require conf.sh; require miss.sh; ish_dev_git_prepare

ish_miss_prepare_compile
ish_miss_prepare_develop
ish_miss_prepare_project

ish_miss_make; if [ -n "$*" ]; then ish_miss_serve "$@"; fi
