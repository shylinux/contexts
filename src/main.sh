#!/bin/sh

main() {
	case "$1" in
		app) # 生产环境
			export PATH="$PWD/bin:$PWD/local/bin:$PWD/usr/local/go/bin:$PATH"
			shift && prepare_ice && while true; do bin/ice.bin serve start dev dev "$@" && break; done
			;;
		dev) # 开发环境
			shift && prepare_system; require miss.sh; [ -f ~/.gitconfig ] || ish_dev_git_prepare
			if [ -n "$ctx_pod" ]; then
				git clone $ctx_dev/chat/pod/$ctx_pod; cd $ctx_pod && source etc/miss.sh dev dev "$@"
			else
				git clone $ctx_dev contexts; cd contexts && source etc/miss.sh dev dev "$@"
			fi
			;;
		cmd) # 命令环境
		   	ish_sys_dev_init >/dev/null; shift; [ -n "$*" ] && ish_sys_dev_run "$@"
			;;
		*)
			echo "hello world"
			;;
	esac
}
main "$@"
