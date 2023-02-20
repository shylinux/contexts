#! /bin/sh

function show() {
	echo "hello world"
}

require "he.sh"
require "/require/shylinux.com/x/contexts@v2.9.2/src/hi/he.sh"
require "http://localhost:9020/require/shylinux.com/x/contexts@v2.9.2/src/hi/he.sh"
require "https://shylinux.com/x/contexts@v2.9.2/src/hi/he.sh?content=what"

show
_list
echo

_list hi
echo

_list hi 1
echo

echo "hello world"
