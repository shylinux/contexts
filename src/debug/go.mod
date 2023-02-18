module shylinux.com/x/contexts

go 1.11

require (
	shylinux.com/x/golang-story v0.5.7
	shylinux.com/x/linux-story v0.5.4
	shylinux.com/x/mysql-story v0.5.7
	shylinux.com/x/nginx-story v0.5.7
	shylinux.com/x/redis-story v0.5.8
)

require (
	golang.org/x/crypto v0.6.0 // indirect
	shylinux.com/x/gogit v0.0.7
	shylinux.com/x/ice v1.2.7
	shylinux.com/x/icebergs v1.5.2
	shylinux.com/x/toolkits v0.7.3
	shylinux.com/x/webview v0.0.2
)

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)

replace (
	shylinux.com/x/golang-story => ./usr/golang-story
	shylinux.com/x/linux-story => ./usr/linux-story
	shylinux.com/x/mysql-story => ./usr/mysql-story
	shylinux.com/x/nginx-story => ./usr/nginx-story
	shylinux.com/x/redis-story => ./usr/redis-story
)
