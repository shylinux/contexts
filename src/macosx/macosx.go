package macosx

import (
	ice "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/ctx"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/core/chat"
	kit "shylinux.com/x/toolkits"
)

const MACOSX = "macosx"

var Index = &ice.Context{Name: MACOSX}

func init() { chat.Index.Register(Index, nil) }

func CmdHashAction(arg ...string) ice.Actions {
	file := kit.FileLines(2)
	return ice.MergeActions(ice.Actions{
		mdb.SELECT: {Name: "list hash auto create", Hand: func(m *ice.Message, arg ...string) { mdb.HashSelect(m, arg...).Display(file) }},
	}, ctx.CmdAction(), mdb.HashAction(mdb.SHORT, kit.Select("", arg, 0), mdb.FIELD, kit.Select("time,hash,type,name,icon,index,args", arg, 1), kit.Slice(arg, 2)))
}
