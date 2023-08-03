#!/bin/sh

touch ~/.hushlogin
export BASH_SILENCE_DEPRECATION_WARNING=1
export LC_ALL=en_US.UTF-8

export CGO_ENABLED=0
export GOPRIVATE=shylinux.com
export GOPROXY=https://goproxy.cn
export GOBIN=~/contexts/usr/local/bin

export CTX_ROOT=${CTX_ROOT:=~/contexts}
[ "$PWD" = "$HOME" ] && cd ~/contexts
[ "$PWD" = "/" ] && cd ~/contexts

if uname -s|grep -v MINGW &>/dev/null; then
	export PATH=/usr/local/bin:/usr/bin:/bin:/usr/local/sbin:/usr/sbin:/sbin
fi
ish_sys_path_load
ish_sys_cli_prompt
ish_sys_cli_alias vi vim
ish_sys_cli_alias v vim
ish_sys_cli_alias t "tmux attach"

# [ -f ~/.bash_temp ] && source ~/.bash_temp; rm ~/.bash_temp &>/dev/null
