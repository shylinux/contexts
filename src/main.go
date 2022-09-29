package main

import (
	"shylinux.com/x/ice"
	_ "shylinux.com/x/icebergs/misc/alpha"
	_ "shylinux.com/x/icebergs/misc/chrome"
	_ "shylinux.com/x/icebergs/misc/coder"
	_ "shylinux.com/x/icebergs/misc/input"

	_ "shylinux.com/x/icebergs/misc/java"
	_ "shylinux.com/x/icebergs/misc/lark"
	_ "shylinux.com/x/icebergs/misc/mp"
	_ "shylinux.com/x/icebergs/misc/node"
	_ "shylinux.com/x/icebergs/misc/wework"
	_ "shylinux.com/x/icebergs/misc/wx"

	_ "shylinux.com/x/golang-story/src/compile"
	_ "shylinux.com/x/golang-story/src/project"
	_ "shylinux.com/x/golang-story/src/runtime"

	_ "shylinux.com/x/golang-story/src/leveldb"
	_ "shylinux.com/x/golang-story/src/rocksdb"
	_ "shylinux.com/x/golang-story/src/tcmalloc"

	_ "shylinux.com/x/golang-story/src/docker"
	_ "shylinux.com/x/golang-story/src/gotags"
	_ "shylinux.com/x/golang-story/src/grafana"
	_ "shylinux.com/x/golang-story/src/kubernetes"
	_ "shylinux.com/x/golang-story/src/prometheus"

	_ "shylinux.com/x/linux-story/iso/alpine"

	_ "shylinux.com/x/linux-story/src/busybox"
	_ "shylinux.com/x/linux-story/src/ctags"
	_ "shylinux.com/x/linux-story/src/ffmpeg"
	_ "shylinux.com/x/linux-story/src/gcc"
	_ "shylinux.com/x/linux-story/src/gdb"
	_ "shylinux.com/x/linux-story/src/glibc"
	_ "shylinux.com/x/linux-story/src/kernel"
	_ "shylinux.com/x/linux-story/src/qemu"
	_ "shylinux.com/x/linux-story/src/sysctl"

	_ "shylinux.com/x/mysql-story/src/client"
	_ "shylinux.com/x/mysql-story/src/server"
	_ "shylinux.com/x/nginx-story/src/client"
	_ "shylinux.com/x/nginx-story/src/server"
	_ "shylinux.com/x/redis-story/src/client"
	_ "shylinux.com/x/redis-story/src/server"

	_ "shylinux.com/x/golang-story/src/data"
	_ "shylinux.com/x/golang-story/src/data/leecode"

	_ "shylinux.com/x/contexts/src/hi"
)

func main() { print(ice.Run()) }
