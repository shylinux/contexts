package macos

import (
	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/mdb"
)

const NOTIFICATIONS = "notifications"

func init() {
	Index.MergeCommands(ice.Commands{NOTIFICATIONS: {Name: "notifications list", Actions: CmdHashAction(), Hand: func(m *ice.Message, arg ...string) {
		mdb.HashSelect(m, arg...).SortStrR(mdb.TIME).Display("")
	}}})
}
