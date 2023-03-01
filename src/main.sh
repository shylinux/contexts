#!/bin/sh

_down_tar() { # 下载文件 file path
	[ -f $1 ] && return; _down_big_file "$@" && tar -xf $1
}
_down_tars() { # 下载文件 file...
	for file in "$@"; do _down_tar $file publish/$file; done
}
main() {
	case "$1" in
		app) # 生产环境
			# export ctx_log=/dev/stdout
			# shift && prepare_ice && bin/ice.bin forever start dev dev "$@"
			shift && prepare_ice && while echo; do
				bin/ice.bin serve start dev dev "$@" && break
			done
			;;
		dev) # 开发环境
			shift && prepare_system
			git config --global url."$ctx_dev".insteadOf https://shylinux.com
			git clone https://shylinux.com/x/contexts
			cd contexts && source etc/miss.sh dev dev "$@"
			;;
		cmd) # 命令环境
		   	ish_sys_dev_init >/dev/null; shift; [ -n "$*" ] && ish_sys_dev_run "$@"
			;;
		*)
			require src/hi/hi.sh
			;;
	esac
}
main "$@"
