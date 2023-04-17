module shylinux.com/x/contexts

go 1.11

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
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
