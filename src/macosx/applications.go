package macosx

import (
	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

const APPLICATIONS = "applications"

func init() {
	Index.MergeCommands(ice.Commands{
		APPLICATIONS: {Name: "applications hash auto create", Help: "应用", Actions: CmdHashAction("index,args"), Hand: func(m *ice.Message, arg ...string) {
			mdb.HashSelect(m, arg...).Options(ice.MSG_HEIGHT, kit.Select("240", "32", len(arg) == 0)).Table(func(value ice.Maps) { m.PushImages(web.IMAGE, "/require/"+value[mdb.ICON]) })
		}},
	})
}
