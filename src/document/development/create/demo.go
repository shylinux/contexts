package demo

import "shylinux.com/x/ice"

type demo struct {
	list string `name:"list path auto" help:"示例模块"`
}

func (s demo) List(m *ice.Message, arg ...string) {
	m.Echo("hello world")
}

func init() { ice.Cmd("web.code.demo", demo{}) }
