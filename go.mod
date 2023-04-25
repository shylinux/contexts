module shylinux.com/x/contexts

go 1.13

replace (
	shylinux.com/x/golang-story => ./usr/golang-story
	shylinux.com/x/linux-story => ./usr/linux-story
	shylinux.com/x/nginx-story => ./usr/nginx-story
	shylinux.com/x/redis-story => ./usr/redis-story
)

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
	shylinux.com/x/golang-story v0.6.0
	shylinux.com/x/linux-story v0.5.6
	shylinux.com/x/nginx-story v0.5.9
	shylinux.com/x/redis-story v0.6.0
)

require (
	shylinux.com/x/ice v1.3.2
	shylinux.com/x/icebergs v1.5.6
	shylinux.com/x/toolkits v0.7.5
)
