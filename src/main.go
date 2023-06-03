package main

import (
	"shylinux.com/x/ice"
	_ "shylinux.com/x/icebergs/misc/java"
	_ "shylinux.com/x/icebergs/misc/node"
)

func main() { print(ice.Run()) }
