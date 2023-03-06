#!/bin/sh

# require "shylinux.com/x/contexts/src/h2/h2.sh"

main() {
	case "$1" in
		app) # 生产环境
			shift && prepare_ice && while true; do
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
			# require src/hi/hi.sh
			echo "hello world"
			;;
	esac
}
main "$@"
