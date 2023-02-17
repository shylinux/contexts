main() {
	case "$1" in
		app) # 生产环境
			export ctx_log=${ctx_log:=/dev/stdout}
			shift && prepare_ice && bin/ice.bin forever start dev dev "$@"
			;;
		dev) # 开发环境
			shift && prepare_package && source etc/miss.sh "$@"
			;;
		cmd) # 命令环境
		   	ish_sys_dev_init >/dev/null
			shift; [ -n "$*" ] && ish_sys_dev_run "$@"
			;;
		*)
			;;
	esac
}
main "$@"

