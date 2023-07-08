module shylinux.com/x/contexts

go 1.13

replace (
	shylinux.com/x/go-git/v5 => ./usr/go-git
	shylinux.com/x/go-qrcode => ./usr/go-qrcode
	shylinux.com/x/websocket => ./usr/websocket
	shylinux.com/x/webview => ./usr/webview
)

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)

require (
	shylinux.com/x/ice v1.3.3
	shylinux.com/x/icebergs v1.5.11
	shylinux.com/x/toolkits v0.7.6
	shylinux.com/x/webview v0.0.2
)
