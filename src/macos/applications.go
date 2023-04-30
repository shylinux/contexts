package macos

import (
	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

const APPLICATIONS = "applications"

func init() {
	Index.MergeCommands(ice.Commands{
		APPLICATIONS: {Name: "applications hash auto create", Help: "应用", Actions: ice.MergeActions(ice.Actions{
			ice.CTX_INIT: {Hand: func(m *ice.Message, arg ...string) {
				m.Cmd(APPLICATIONS, mdb.CREATE, mdb.NAME, "Finder", mdb.ICON, "usr/icons/Finder.png", ctx.INDEX, "nfs.dir")
				m.Cmd(APPLICATIONS, mdb.CREATE, mdb.NAME, "Safari", mdb.ICON, "usr/icons/Safari.png", ctx.INDEX, "web.chat.iframe")
				m.Cmd(APPLICATIONS, mdb.CREATE, mdb.NAME, "Preview", mdb.ICON, "usr/icons/Preview.png", ctx.INDEX, "web.wiki.feel")
				m.Cmd(APPLICATIONS, mdb.CREATE, mdb.NAME, "Terminal", mdb.ICON, "usr/icons/Terminal.png", ctx.INDEX, "web.code.xterm")
				m.Cmd(APPLICATIONS, mdb.CREATE, mdb.NAME, "Calendar", mdb.ICON, "usr/icons/Calendar.png", ctx.INDEX, "web.team.plan", ctx.ARGS, "month")
				m.Cmd(APPLICATIONS, mdb.CREATE, mdb.NAME, "Grapher", mdb.ICON, "usr/icons/Grapher.png", ctx.INDEX, "web.wiki.draw")
				m.Cmd(APPLICATIONS, mdb.CREATE, mdb.NAME, "Books", mdb.ICON, "usr/icons/Books.png", ctx.INDEX, "web.wiki.word")
				m.Cmd(APPLICATIONS, mdb.CREATE, mdb.NAME, "vim", mdb.ICON, "usr/icons/vim.png", ctx.INDEX, "web.code.vimer")

				m.Cmd(FINDER, mdb.CREATE, mdb.NAME, "Applications", ctx.INDEX, Prefix(APPLICATIONS))
				m.Cmd(FINDER, mdb.CREATE, mdb.NAME, "Pictures", ctx.INDEX, "web.wiki.feel")
				m.Cmd(FINDER, mdb.CREATE, mdb.NAME, "Trash", ctx.INDEX, "nfs.trash")
			}},
		}, CmdHashAction("index,args")), Hand: func(m *ice.Message, arg ...string) {
			mdb.HashSelect(m, arg...).Options(ice.MSG_HEIGHT, kit.Select("240", "32", len(arg) == 0)).Table(func(value ice.Maps) { m.PushImages(web.IMAGE, "/require/"+value[mdb.ICON]) })
		}},
	})
}
