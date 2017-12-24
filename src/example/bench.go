package main

import (
	"context"
	_ "context/aaa"
	_ "context/cli"
	_ "context/ssh"

	_ "context/mdb"
	_ "context/nfs"
	_ "context/tcp"
	_ "context/web"

	_ "context/lex"

	"os"
)

func main() {
	ctx.Start(os.Args[1:]...)
}