package macosx

import (
	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/nfs"
	kit "shylinux.com/x/toolkits"
)

const DOCK = "dock"

func init() {
	Index.MergeCommands(ice.Commands{
		DOCK: {Actions: ice.MergeActions(ice.Actions{
			mdb.INPUTS: {Hand: func(m *ice.Message, arg ...string) {
				switch mdb.HashInputs(m, arg); arg[0] {
				case mdb.NAME:
					m.Cmd(nfs.DIR, "usr/icons/", func(value ice.Maps) { m.Push(arg[0], kit.TrimExt(value[nfs.PATH], nfs.PNG)) })
				case mdb.ICON:
					m.Cmd(nfs.DIR, "usr/icons/", func(value ice.Maps) { m.Push(arg[0], value[nfs.PATH]) })
				}
			}},
		}, CmdHashAction())},
	})
}
