package h3

import (
	"shylinux.com/x/ice"
)

type h3 struct {
	ice.Zone

	list string `name:"list zone id auto insert" help:"h3"`
}

func (s h3) List(m *ice.Message, arg ...string) {
	s.Zone.List(m, arg...)
}

func init() { ice.Cmd("web.code.h3.h3", h3{}) }
