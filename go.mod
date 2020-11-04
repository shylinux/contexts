module github.com/shylinux/contexts

go 1.13

require (
	github.com/Baozisoftware/qrcode-terminal-go v0.0.0-20170407111555-c0650d8dff0f // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/shylinux/golang-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/icebergs v0.2.9
	github.com/shylinux/linux-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/mysql-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/nginx-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/redis-story v0.0.0-00010101000000-000000000000
	github.com/shylinux/toolkits v0.1.9
)

replace github.com/shylinux/icebergs => ./usr/icebergs

replace github.com/shylinux/toolkits => ./usr/toolkits

replace github.com/shylinux/linux-story => ./usr/linux-story

replace github.com/shylinux/nginx-story => ./usr/nginx-story

replace github.com/shylinux/golang-story => ./usr/golang-story

replace github.com/shylinux/redis-story => ./usr/redis-story

replace github.com/shylinux/mysql-story => ./usr/mysql-story
