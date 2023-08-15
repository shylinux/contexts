module shylinux.com/x/contexts

go 1.13

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)

require (
	shylinux.com/x/ice v1.3.13
	shylinux.com/x/icebergs v1.5.19
	shylinux.com/x/toolkits v0.7.10
)

require shylinux.com/x/webview v0.0.2
