package macos

import (
	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

const APPLICATIONS = "applications"

func init() {
	Index.MergeCommands(ice.Commands{
		APPLICATIONS: {Name: "applications hash auto create", Help: "应用", Actions: ice.MergeActions(ice.Actions{
			ice.CTX_INIT: {Hand: func(m *ice.Message, arg ...string) {
				Install(m, "Finder", "nfs.dir")
				Install(m, "Safari", "web.chat.iframe")
				Install(m, "Photos", "web.wiki.feel")
				Install(m, "Terminal", "web.code.xterm")
				Install(m, "Calendar", "web.team.plan", ctx.ARGS, "month")
				Install(m, "Grapher", "web.wiki.draw")
				Install(m, "Books", "web.wiki.word")
				m.Cmd(FINDER, mdb.CREATE, mdb.NAME, "Applications", ctx.INDEX, Prefix(APPLICATIONS))
				m.Cmd(FINDER, mdb.CREATE, mdb.NAME, "Pictures", ctx.INDEX, "web.wiki.feel")
				m.Cmd(FINDER, mdb.CREATE, mdb.NAME, "Trash", ctx.INDEX, "nfs.trash")
			}},
			"install": {Hand: func(m *ice.Message, arg ...string) { Install(m, arg[0], arg[1], arg[2:]...) }},
		}, CmdHashAction("index,args")), Hand: func(m *ice.Message, arg ...string) {
			mdb.HashSelect(m, arg...).Options(ice.MSG_HEIGHT, kit.Select("240", "32", len(arg) == 0)).Table(func(value ice.Maps) { m.PushImages(web.IMAGE, "/require/"+value[mdb.ICON]) })
			m.Sort(mdb.NAME)
		}},
	})
}
func Install(m *ice.Message, name, index string, arg ...string) {
	name, icon := kit.Select(kit.Select("", kit.Split(index, ice.PT), -1), name), ""
	kit.If(nfs.Exists(m, "usr/icons/"+name+".png"), func() { icon = "usr/icons/" + name + ".png" })
	m.Cmd(Prefix(APPLICATIONS), mdb.CREATE, mdb.NAME, name, mdb.ICON, icon, ctx.INDEX, index, arg)
}
