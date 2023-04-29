package macosx

import (
	"strings"

	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
)

const FINDER = "finder"

func init() {
	Index.MergeCommands(ice.Commands{
		FINDER: {Name: "finder path auto create", Actions: CmdHashAction(), Hand: func(m *ice.Message, arg ...string) {
			if len(arg) == 0 || strings.HasSuffix(arg[0], nfs.PS) {
				switch kit.Select("", arg, 0) {
				case ice.USR_LOCAL_WORK:
					ctx.ProcessCmds(m, "web.dream")
				case ice.USR_LOCAL_REPOS:
					ctx.ProcessCmds(m, "web.code.git.service")
				case ice.USR_LOCAL_IMAGE:
					ctx.ProcessCmds(m, "web.wiki.feel")
				case ice.USR_LOCAL_DAEMON:
					ctx.ProcessCmds(m, "web.code.install")
				default:
					m.Cmdy(nfs.DIR, arg).Display("")
				}
			} else {
				switch kit.Ext(arg[0]) {
				case "svg":
					ctx.ProcessCmds(m, "web.wiki.draw", arg[0])
				case "shy":
					ctx.ProcessCmds(m, "web.wiki.word", arg[0])
				default:
					ls := nfs.SplitPath(m, arg[0])
					ctx.ProcessCmds(m, "web.code.vimer", ls[0], ls[1])
				}
			}
		}},
	})
}
