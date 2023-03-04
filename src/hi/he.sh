#!/bin/sh

show() {
	echo "$SHELL $TERM $PWD $HOME"
	echo "hello world $content $@"
	ish_sys_path_list
}
show
