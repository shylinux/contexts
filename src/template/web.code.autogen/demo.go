package {{.Option "zone"}}

import "shylinux.com/x/ice"

type {{.Option "name"}} struct {
	ice.{{.Option "type"}}

	list string {{.Option "text"}}
}

func (s {{.Option "name"}}) List(m *ice.Message, arg ...string) {
	s.{{.Option "type"}}.List(m, arg...)
}

func init() { ice.Cmd("{{.Option "key"}}", {{.Option "name"}}{}) }
