package hi

import (
	"shylinux.com/x/ice"
)

type hi struct {
	ice.Zone

	list string `name:"list zone id auto insert" help:"示例"`
}

func (h hi) List(m *ice.Message, arg ...string) {
	h.Zone.List(m, arg...)
}

func init() { ice.Cmd("web.code.hi.hi", hi{}) }

