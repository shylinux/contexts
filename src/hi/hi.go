package hi

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/ctx"
)

type hi struct {
	ice.Zone

	list string `name:"list zone id auto insert" help:"示例"`
}

func (s hi) Show(m *ice.Message, arg ...string) {
	m.Echo("hello world").StatusTime()
}

func (s hi) List(m *ice.Message, arg ...string) {
	// ctx.Display(m, "hi.js")
	// ctx.Display(m, "/require/shylinux.com/x/contexts@v2.9.2/src/hi/hi.js")
	// ctx.Display(m, "http://localhost:9020/require/shylinux.com/x/contexts@v2.9.2/src/hi/hi.js")
	ctx.Display(m, "https://shylinux.com/x/contexts@v2.9.2/src/hi/hi.js?content=what")
	s.Zone.ListPage(m, arg...)
}

func init() { ice.Cmd("web.code.hi.hi", hi{}) }
