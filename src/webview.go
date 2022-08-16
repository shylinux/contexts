package main

import (
	"os"
	"time"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/tcp"
	"shylinux.com/x/icebergs/base/web"
	"shylinux.com/x/icebergs/misc/webview"
	kit "shylinux.com/x/toolkits"
)

type view struct{ *webview.WebView }

func main() {
	os.Chdir(kit.HomePath(ice.CONTEXTS))
	go ice.Run(web.SERVE, tcp.START)
	time.Sleep(time.Second)
	webview.Run(func(w *webview.WebView) ice.Any { return view{w} })
}
