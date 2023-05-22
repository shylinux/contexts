module shylinux.com/x/contexts

go 1.13

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)

replace (
	shylinux.com/x/go-git => ./usr/go-git
	shylinux.com/x/go-qrcode => ./usr/go-qrcode
	shylinux.com/x/websocket => ./usr/websocket
)

require (
	shylinux.com/x/ice v1.3.2
	shylinux.com/x/icebergs v1.5.6
	shylinux.com/x/toolkits v0.7.5
)
