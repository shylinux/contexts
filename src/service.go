package main

import (
	"os"

	"shylinux.com/x/ice"
	_ "shylinux.com/x/icebergs/misc/ssh"
)

func main() {
	defer func() { recover() }()
	print(ice.Run(append([]string{"ssh.service", "listen"}, os.Args[1:]...)...))
}