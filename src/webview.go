package main

import (
	"os"
	"time"

	"shylinux.com/x/ice"
	_ "shylinux.com/x/icebergs/misc/alpha"
	_ "shylinux.com/x/icebergs/misc/chrome"
	_ "shylinux.com/x/icebergs/misc/java"
	_ "shylinux.com/x/icebergs/misc/node"
	kit "shylinux.com/x/toolkits"

	"shylinux.com/x/icebergs/misc/webview"
)

func main() {
	os.Setenv("PATH", "/usr/local/bin:/usr/bin:/bin:/usr/local/sbin:/usr/sbin:/sbin")
	os.Chdir(kit.HomePath(ice.CONTEXTS))
	go ice.Run(ice.SERVE, ice.START)
	defer ice.Pulse.Cmd(ice.EXIT)
	time.Sleep(time.Second)
	webview.Run(nil)
}
