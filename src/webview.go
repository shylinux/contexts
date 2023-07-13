package main

import (
	"os"

	"shylinux.com/x/ice"
	_ "shylinux.com/x/icebergs/misc/alpha"
	_ "shylinux.com/x/icebergs/misc/chrome"
	_ "shylinux.com/x/icebergs/misc/java"
	_ "shylinux.com/x/icebergs/misc/node"
	"shylinux.com/x/icebergs/misc/webview"
	_ "shylinux.com/x/icons"
	kit "shylinux.com/x/toolkits"
)

func main() {
	os.Chdir(kit.HomePath(ice.CONTEXTS))
	go ice.Run(ice.SERVE, ice.START)
	defer ice.Pulse.Cmd(ice.EXIT)
	webview.Run(nil)
}
