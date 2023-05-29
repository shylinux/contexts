module shylinux.com/x/contexts

go 1.13

replace (
	shylinux.com/x/go-git/v5 => ./usr/go-git
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)

require (
	shylinux.com/x/ice v1.3.3
	shylinux.com/x/icebergs v1.5.11
	shylinux.com/x/toolkits v0.7.6
)
