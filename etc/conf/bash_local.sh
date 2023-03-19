#!/bin/sh

export CGO_ENABLED=0
export LC_ALL=en_US.UTF-8
export BASH_SILENCE_DEPRECATION_WARNING=1
touch ~/.hushlogin

export CTX_ROOT=${CTX_ROOT:=~/contexts}
[ "$PWD" = "$HOME" ] && cd ~/contexts
[ "$PWD" = "/" ] && cd ~/contexts

if uname -s|grep -v MINGW &>/dev/null; then
	export PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
fi
ish_sys_path_load
ish_sys_cli_prompt
ish_sys_cli_alias vi vim
ish_sys_cli_alias v vim
ish_sys_cli_alias t "tmux attach"
ish_sys_cli_alias busybox "docker run -w /root -it busybox"
ish_sys_cli_alias alpine-dev "docker run -w /root -e 'LANG=en_US.UTF-8' -it alpine-dev sh"
ish_sys_cli_alias alpine "docker run -w /root -e 'LANG=en_US.UTF-8' -it alpine"
ish_sys_cli_alias centos "docker run -w /root -e 'LANG=en_US.UTF-8' -it centos"

[ -f ~/.bash_temp ] && source ~/.bash_temp; rm ~/.bash_temp &>/dev/null
