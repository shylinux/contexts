module shylinux.com/x/contexts

go 1.11

require (
	shylinux.com/x/golang-story v0.4.8
	shylinux.com/x/linux-story v0.4.6
	shylinux.com/x/mysql-story v0.4.9
	shylinux.com/x/nginx-story v0.4.9
	shylinux.com/x/redis-story v0.5.0
)

require (
	shylinux.com/x/ice v1.1.3
	shylinux.com/x/icebergs v1.3.8
	shylinux.com/x/toolkits v0.6.8
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
