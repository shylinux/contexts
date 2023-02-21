#!/bin/sh

_down_tar() { # 下载文件 file path
	[ -f $1 ] && return; _down_big_file "$@" && tar -xf $1
}
_down_tars() { # 下载文件 file...
	for file in "$@"; do _down_tar $file publish/$file; done
}
prepare_package() {
	_down_tars contexts.bin.tar.gz contexts.src.tar.gz
	local back=$PWD; cd ~/; _down_tars contexts.home.tar.gz; cd $back
	export VIM=$PWD/usr/install/vim-vim-12be734/_install/share/vim/vim82/
	export LD_LIBRARY_PATH=$PWD/usr/local/lib

	ish_sys_path_load
	git config --global init.templatedir $PWD/usr/install/git-2.31.1/_install/share/git-core/templates/
	git config --global url."$ctx_dev".insteadOf https://shylinux.com
	git config --global init.defaultBranch master
}
main() {
	case "$1" in
		app) # 生产环境
			shift && prepare_ice && bin/ice.bin forever start dev dev "$@"
			;;
		dev) # 开发环境
			shift && prepare_package && source etc/miss.sh "$@"
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
