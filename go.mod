module shylinux.com/x/contexts

go 1.11

require (
	shylinux.com/x/golang-story v0.5.4
	shylinux.com/x/linux-story v0.5.1
	shylinux.com/x/mysql-story v0.5.4
	shylinux.com/x/nginx-story v0.5.4
	shylinux.com/x/redis-story v0.5.5
)

require (
	shylinux.com/x/gogit v0.0.6
	shylinux.com/x/ice v1.2.1
	shylinux.com/x/icebergs v1.4.8
	shylinux.com/x/toolkits v0.7.2
	shylinux.com/x/webview v0.0.1
)

replace (
	shylinux.com/x/creackpty => ./usr/creackpty
	shylinux.com/x/go-qrcode => ./usr/go-qrcode
	shylinux.com/x/go-sql-mysql => ./usr/go-sql-mysql
	shylinux.com/x/gogit => ./usr/gogit
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
	shylinux.com/x/websocket => ./usr/websocket
)

replace (
	shylinux.com/x/golang-story => ./usr/golang-story
	shylinux.com/x/linux-story => ./usr/linux-story
	shylinux.com/x/mysql-story => ./usr/mysql-story
	shylinux.com/x/nginx-story => ./usr/nginx-story
	shylinux.com/x/redis-story => ./usr/redis-story
)
