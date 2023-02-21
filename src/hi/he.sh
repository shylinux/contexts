#!/bin/sh

show() {
	echo "$SHELL $TERM $PWD $HOME"
	echo "$PATH"|tr ":" "\n"
	echo "$ISH_CTX_SCRIPT <= $ISH_CTX_MODULE"
	echo "hello world $content $@"
}
show
