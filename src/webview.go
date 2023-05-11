package main

import (
	"os"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/tcp"
	"shylinux.com/x/icebergs/base/web"
	_ "shylinux.com/x/icebergs/misc/node"
	"shylinux.com/x/icebergs/misc/webview"
	kit "shylinux.com/x/toolkits"
)

func main() {
	os.Chdir(kit.HomePath(ice.CONTEXTS))
	go ice.Run(web.SERVE, tcp.START)
	defer ice.Pulse.Cmd("exit")
	webview.Run(nil)
}
