#!/bin/sh

show() {
	echo "$TERM $PWD $HOME $PATH"
	echo "$ISH_CTX_SCRIPT <= $ISH_CTX_MODULE"
	echo "hello world $content $@"
}
show
