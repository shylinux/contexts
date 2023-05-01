package macos

import (
	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/log"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

const FINDER = "finder"

func init() {
	Index.MergeCommands(ice.Commands{
		FINDER: {Name: "finder list", Actions: ice.MergeActions(ice.Actions{
			ice.CTX_INIT: {Hand: func(m *ice.Message, arg ...string) {
				if m.Cmd(DOCK).Length() == 0 {
					DockAppend(m, FINDER, m.PrefixKey())
					DockAppend(m, "Safari", web.CHAT_IFRAME)
					DockAppend(m, "Terminal", web.CODE_XTERM)
					DockAppend(m, "", web.CODE_VIMER)
				}
			}},
			mdb.SEARCH: {Hand: func(m *ice.Message, arg ...string) {
				mdb.IsSearchForEach(m, arg, func() []string { return []string{web.LINK, DESKTOP, m.MergePodCmd("", DESKTOP, log.DEBUG, ice.TRUE)} })
			}},
		}, CmdHashAction(mdb.NAME))},
	})
}
func DockAppend(m *ice.Message, name, index string, arg ...string) {
	name, icon := kit.Select(kit.Select("", kit.Split(index, ice.PT), -1), name), ""
	kit.If(nfs.Exists(m, kit.PathJoin(USR_ICONS, name, nfs.PNG)), func() { icon = kit.PathJoin(USR_ICONS, name, nfs.PNG) })
	m.Cmd(Prefix(DOCK), mdb.CREATE, mdb.NAME, name, mdb.ICON, icon, ctx.INDEX, index, arg)
}
