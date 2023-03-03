package main

import (
	"os"
	"time"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/cli"
	"shylinux.com/x/icebergs/base/tcp"
	"shylinux.com/x/icebergs/base/web"
	"shylinux.com/x/icebergs/misc/webview"
	kit "shylinux.com/x/toolkits"

	_ "shylinux.com/x/contexts/src/h2"
)

type view struct{ *webview.WebView }

func main() {
	if os.Chdir(kit.HomePath(ice.CONTEXTS)); len(os.Args) == 1 {
		for {
			if ice.Run(cli.SYSTEM, os.Args[0], "webview"); cli.IsSuccess(ice.Pulse) {
				break
			}
		}
	} else {
		go ice.Run(web.SERVE, tcp.START)
		defer ice.Pulse.Cmd("exit")
		time.Sleep(time.Second)
		webview.Run(func(w *webview.WebView) ice.Any { return view{w} })
	}
}
