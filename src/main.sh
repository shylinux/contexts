#!/bin/sh

main() {
	case "$1" in
		app) # 生产环境
			export PATH="$PWD/bin:$PWD/local/bin:$PWD/usr/local/go/bin:$PATH"
			shift && prepare_ice && while true; do bin/ice.bin serve start dev dev "$@" && break; done
			;;
		dev) # 开发环境
			shift && prepare_system; require miss.sh; [ -f ~/.gitconfig ] || ish_dev_git_prepare
			git config --global "url.$ctx_dev/x/.insteadof" "${ctx_repos%/*}/"
			git clone $ctx_repos ${ctx_repos##*/}; cd ${ctx_repos##*/} && source etc/miss.sh
			ish_miss_serve_log dev dev "$@"
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
