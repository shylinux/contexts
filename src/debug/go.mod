module shylinux.com/x/contexts

go 1.11

require (
	shylinux.com/x/ice v0.3.1
	shylinux.com/x/icebergs v0.5.9
	shylinux.com/x/toolkits v0.3.6
)

require (
	shylinux.com/x/golang-story v0.1.9
	shylinux.com/x/linux-story v0.1.9
	shylinux.com/x/mysql-story v0.2.0
	shylinux.com/x/nginx-story v0.2.0
	shylinux.com/x/redis-story v0.2.1
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