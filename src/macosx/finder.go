package macosx

import (
	"strings"

	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
)

const FINDER = "finder"

func init() {
	Index.MergeCommands(ice.Commands{
		FINDER: {Name: "finder list", Actions: ice.MergeActions(ice.Actions{
			ice.CTX_INIT: {Hand: func(m *ice.Message, arg ...string) {
				m.Cmd(APPLICATIONS, mdb.CREATE, mdb.NAME, "Finder", mdb.ICON, "usr/icons/Finder.png", ctx.INDEX, m.PrefixKey())
				m.Cmd(APPLICATIONS, mdb.CREATE, mdb.NAME, "Safari", mdb.ICON, "usr/icons/Safari.png", ctx.INDEX, "web.chat.iframe")
				m.Cmd(APPLICATIONS, mdb.CREATE, mdb.NAME, "Preview", mdb.ICON, "usr/icons/Preview.png", ctx.INDEX, "web.wiki.feel")
				m.Cmd(APPLICATIONS, mdb.CREATE, mdb.NAME, "Terminal", mdb.ICON, "usr/icons/Terminal.png", ctx.INDEX, "web.code.xterm")
				m.Cmd(APPLICATIONS, mdb.CREATE, mdb.NAME, "vim", mdb.ICON, "usr/icons/vim.png", ctx.INDEX, "web.code.vimer")
				if m.Cmd(DOCK).Length() == 0 {
					m.Cmd(DOCK, mdb.CREATE, mdb.NAME, "Finder", mdb.ICON, "usr/icons/Finder.png", ctx.INDEX, m.PrefixKey())
					m.Cmd(DOCK, mdb.CREATE, mdb.NAME, "Terminal", mdb.ICON, "usr/icons/Terminal.png", ctx.INDEX, "web.code.xterm")
				}
				m.Cmd(FINDER, mdb.CREATE, mdb.NAME, "Applications", ctx.INDEX, Prefix(APPLICATIONS))
				m.Cmd(FINDER, mdb.CREATE, mdb.NAME, "contexts", ctx.INDEX, "nfs.dir")
			}},
		}, CmdHashAction(mdb.NAME)), Hand: func(m *ice.Message, arg ...string) {
			if len(arg) == 0 {
				mdb.HashSelect(m, arg...)
				m.Display("")
			} else if len(arg) == 1 || strings.HasSuffix(arg[1], nfs.PS) {
				switch kit.Select("", arg, 1) {
				case ice.USR_LOCAL_WORK:
					ctx.ProcessCmds(m, "web.dream")
				case ice.USR_LOCAL_REPOS:
					ctx.ProcessCmds(m, "web.code.git.service")
				case ice.USR_LOCAL_IMAGE:
					ctx.ProcessCmds(m, "web.wiki.feel")
				case ice.USR_LOCAL_DAEMON:
					ctx.ProcessCmds(m, "web.code.install")
				default:
					m.Cmdy(nfs.DIR, arg[1:]).Display("")
				}
			} else {
				switch kit.Ext(arg[1]) {
				case "svg":
					ctx.ProcessCmds(m, "web.wiki.draw", arg[1])
				case "shy":
					ctx.ProcessCmds(m, "web.wiki.word", arg[1])
				default:
					ls := nfs.SplitPath(m, arg[1])
					ctx.ProcessCmds(m, "web.code.vimer", ls[0], ls[1])
				}
			}
		}},
	})
}
