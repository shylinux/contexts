package {{.Option "zone"}}

import (
	"shylinux.com/x/ice"
)

type {{.Option "name"}} struct {
	list string `name:"list path auto" help:"示例"`
}

func (s {{.Option "name"}}) List(m *ice.Message, arg ...string) {
	m.Echo("hello world")
}

func init() { ice.Cmd("{{.Option "key"}}", {{.Option "name"}}{}) }
