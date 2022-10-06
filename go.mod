module shylinux.com/x/contexts

go 1.11

require (
	shylinux.com/x/golang-story v0.5.1
	shylinux.com/x/linux-story v0.4.9
	shylinux.com/x/mysql-story v0.5.2
	shylinux.com/x/nginx-story v0.5.2
	shylinux.com/x/redis-story v0.5.3
	shylinux.com/x/webview v0.0.1
)

require (
	shylinux.com/x/ice v1.1.7
	shylinux.com/x/icebergs v1.4.3
	shylinux.com/x/toolkits v0.7.0
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
