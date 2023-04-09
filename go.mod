module shylinux.com/x/contexts

go 1.11

replace (
	shylinux.com/x/golang-story => ./usr/golang-story
	shylinux.com/x/linux-story => ./usr/linux-story
	shylinux.com/x/mysql-story => ./usr/mysql-story
	shylinux.com/x/nginx-story => ./usr/nginx-story
	shylinux.com/x/redis-story => ./usr/redis-story
)

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)

replace (
	shylinux.com/x/creackpty => ./usr/creackpty
	shylinux.com/x/go-qrcode => ./usr/go-qrcode
	shylinux.com/x/go-sql-mysql => ./usr/go-sql-mysql
	shylinux.com/x/gogit => ./usr/gogit
	shylinux.com/x/websocket => ./usr/websocket
	shylinux.com/x/webview => ./usr/webview
)

require (
	shylinux.com/x/golang-story v0.5.8
	shylinux.com/x/linux-story v0.5.4
	shylinux.com/x/mysql-story v0.5.7
	shylinux.com/x/nginx-story v0.5.7
	shylinux.com/x/redis-story v0.5.8
)

require (
	shylinux.com/x/ice v1.3.0
	shylinux.com/x/icebergs v1.5.4
	shylinux.com/x/toolkits v0.7.4
)

require (
	fyne.io/fyne v1.4.3 // indirect
	github.com/alecthomas/chroma v0.10.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0
	github.com/gomarkdown/markdown v0.0.0-20230322041520-c84983bdbf2a
	github.com/gomodule/redigo/redis v0.0.1
	github.com/gorilla/websocket v1.5.0
	github.com/nsf/termbox-go v1.1.1
	github.com/pkg/sftp v1.13.5 // indirect
	github.com/shylinux/icebergs v0.3.8
	github.com/shylinux/toolkits v0.2.6
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	shylinux.com/x/gogit v0.0.7
	shylinux.com/x/webview v0.0.2
)
