package main

import (
	"shylinux.com/x/ice"
	_ "shylinux.com/x/icebergs/misc/node"

	_ "shylinux.com/x/contexts/src/h2"
)

func main() { print(ice.Run()) }
