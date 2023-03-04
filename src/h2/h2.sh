#!/bin/sh

# set -v
# export ISH_CONF_LEVEL="require request notice debug alias"
# require "h3.sh"
require "src/h2/h3.sh"
require "src/h2/h3.sh?welcome=hi"
require "shylinux.com/x/contexts/src/hi/he.sh"
# require "http://localhost:9020/x/contexts/src/hi/he.sh"
# require "http://localhost:9020/x/contexts@v2.9.3/src/hi/he.sh"

demo() {
	echo "hello world"
}
demo
demo1
request nfs.dir
ish_sys_cli_shell
