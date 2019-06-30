package ctx

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"toolkit"
)

type Cache struct {
	Value string
	Name  string
	Help  string
	Hand  func(m *Message, x *Cache, arg ...string) string
}
type Config struct {
	Value interface{}
	Name  string
	Help  string
	Hand  func(m *Message, x *Config, arg ...string) string
}
type Command struct {
	Form map[string]int
	Name string
	Help interface{}
	Auto func(m *Message, c *Context, key string, arg ...string) (ok bool)
	Hand func(m *Message, c *Context, key string, arg ...string) (e error)
}
type Context struct {
	Name string
	Help string

	Caches   map[string]*Cache
	Configs  map[string]*Config
	Commands map[string]*Command

	message  *Message
	requests []*Message
	sessions []*Message

	contexts map[string]*Context
	context  *Context
	root     *Context

	exit chan bool
	Server
}
type Server interface {
	Spawn(m *Message, c *Context, arg ...string) Server
	Begin(m *Message, arg ...string) Server
	Start(m *Message, arg ...string) bool
	Close(m *Message, arg ...string) bool
}

func (c *Context) Context() *Context {
	return c.context
}
func (c *Context) Message() *Message {
	return c.message
}

type Message struct {
	time time.Time
	code int

	source *Context
	target *Context

	Hand bool
	Meta map[string][]string
	Data map[string]interface{}

	callback func(msg *Message) (sub *Message)
	freeback []func(msg *Message) (done bool)
	Sessions map[string]*Message

	messages []*Message
	message  *Message
	root     *Message
}
type LOGGER interface {
	Log(*Message, string, string, ...interface{})
}
type DEBUG interface {
	Wait(*Message, ...interface{}) interface{}
	Goon(interface{}, ...interface{})
}

func (m *Message) Time(arg ...interface{}) string {
	t := m.time
	if len(arg) > 0 {
		if d, e := time.ParseDuration(arg[0].(string)); e == nil {
			arg = arg[1:]
			t.Add(d)
		}
	}

	str := m.Conf("time_format")
	if len(arg) > 1 {
		str = fmt.Sprintf(arg[0].(string), arg[1:]...)
	} else if len(arg) > 0 {
		str = fmt.Sprintf("%v", arg[0])
	}

	if str == "stamp" {
		return kit.Format(t.Unix())
	}
	return t.Format(str)
}
func (m *Message) Code() int {
	return m.code
}
func (m *Message) Source() *Context {
	return m.source
}
func (m *Message) Target() *Context {
	return m.target
}
func (m *Message) Message() *Message {
	return m.message
}

func (m *Message) Detail(arg ...interface{}) string {
	noset, index := true, 0
	if len(arg) > 0 {
		switch v := arg[0].(type) {
		case int:
			noset, index, arg = false, v, arg[1:]
		}
	}
	if noset && len(arg) > 0 {
		index = -2
	}

	return m.Insert("detail", index, arg...)
}
func (m *Message) Detaili(arg ...interface{}) int {
	return kit.Int(m.Detail(arg...))
}
func (m *Message) Details(arg ...interface{}) bool {
	return kit.Right(m.Detail(arg...))
}
func (m *Message) Option(key string, arg ...interface{}) string {
	if len(arg) > 0 {
		m.Insert(key, 0, arg...)
		if _, ok := m.Meta[key]; ok {
			m.Add("option", key)
		}
	}

	for msg := m; msg != nil; msg = msg.message {
		if !msg.Has(key) {
			continue
		}
		for _, k := range msg.Meta["option"] {
			if k == key {
				return msg.Get(key)
			}
		}
	}
	return ""
}
func (m *Message) Optioni(key string, arg ...interface{}) int {
	return kit.Int(m.Option(key, arg...))

}
func (m *Message) Options(key string, arg ...interface{}) bool {
	return kit.Right(m.Option(key, arg...))
}
func (m *Message) Optionv(key string, arg ...interface{}) interface{} {
	if len(arg) > 0 {
		switch arg[0].(type) {
		case nil:
		// case []string:
		// 	m.Option(key, v...)
		// case string:
		// 	m.Option(key, v)
		default:
			m.Put("option", key, arg[0])
		}
	}

	for msg := m; msg != nil; msg = msg.message {
		if msg.Meta == nil || !msg.Has(key) {
			continue
		}
		for _, k := range msg.Meta["option"] {
			if k == key {
				if v, ok := msg.Data[key]; ok {
					return v
				}
				return msg.Meta[key]
			}
		}
	}
	return nil
}
func (m *Message) Optionx(key string, arg ...string) interface{} {
	value := m.Conf(key)
	if value == "" {
		value = m.Option(key)
	}

	if len(arg) > 0 {
		value = fmt.Sprintf(arg[0], value)
	}
	return value
}
func (m *Message) Append(key string, arg ...interface{}) string {
	if len(arg) > 0 {
		m.Insert(key, 0, arg...)
		if _, ok := m.Meta[key]; ok {
			m.Add("append", key)
		}
	}

	ms := []*Message{m}
	for i := 0; i < len(ms); i++ {
		ms = append(ms, ms[i].messages...)
		if !ms[i].Has(key) {
			continue
		}
		for _, k := range ms[i].Meta["append"] {
			if k == key {
				return ms[i].Get(key)
			}
		}
	}
	return ""
}
func (m *Message) Appendi(key string, arg ...interface{}) int64 {
	i, _ := strconv.ParseInt(m.Append(key, arg...), 10, 64)
	return i
}
func (m *Message) Appends(key string, arg ...interface{}) bool {
	return kit.Right(m.Append(key, arg...))
}
func (m *Message) Appendv(key string, arg ...interface{}) interface{} {
	if len(arg) > 0 {
		m.Put("append", key, arg[0])
	}

	ms := []*Message{m}
	for i := 0; i < len(ms); i++ {
		ms = append(ms, ms[i].messages...)
		if !ms[i].Has(key) {
			continue
		}
		for _, k := range ms[i].Meta["append"] {
			if k == key {
				if v, ok := ms[i].Data[key]; ok {
					return v
				}
				return ms[i].Meta[key]
			}
		}
	}
	return nil
}
func (m *Message) Result(arg ...interface{}) string {
	noset, index := true, 0
	if len(arg) > 0 {
		switch v := arg[0].(type) {
		case int:
			noset, index, arg = false, v, arg[1:]
		}
	}
	if noset && len(arg) > 0 {
		index = -2
	}

	return m.Insert("result", index, arg...)
}
func (m *Message) Resulti(arg ...interface{}) int {
	return kit.Int(m.Result(arg...))
}
func (m *Message) Results(arg ...interface{}) bool {
	return kit.Right(m.Result(arg...))
}

func (m *Message) Table(cbs ...interface{}) *Message {
	if len(m.Meta["append"]) == 0 {
		return m
	}

	// 遍历函数
	if len(cbs) > 0 {
		nrow := len(m.Meta[m.Meta["append"][0]])
		for i := 0; i < nrow; i++ {
			line := map[string]string{}
			for _, k := range m.Meta["append"] {
				line[k] = m.Meta[k][i]
			}

			switch cb := cbs[0].(type) {
			case func(map[string]string):
				cb(line)
			case func(map[string]string) bool:
				if !cb(line) {
					return m
				}
			case func(int, map[string]string):
				cb(i, line)
			}
		}
		return m
	}

	//计算列宽
	space := m.Confx("table_space")
	depth, width := 0, map[string]int{}
	for _, k := range m.Meta["append"] {
		if len(m.Meta[k]) > depth {
			depth = len(m.Meta[k])
		}
		width[k] = kit.Width(k, len(space))
		for _, v := range m.Meta[k] {
			if kit.Width(v, len(space)) > width[k] {
				width[k] = kit.Width(v, len(space))
			}
		}
	}

	// 回调函数
	var cb func(maps map[string]string, list []string, line int) (goon bool)
	if len(cbs) > 0 {
		cb = cbs[0].(func(maps map[string]string, list []string, line int) (goon bool))
	} else {
		row := m.Confx("table_row_sep")
		col := m.Confx("table_col_sep")
		compact := kit.Right(m.Confx("table_compact"))
		cb = func(maps map[string]string, lists []string, line int) bool {
			for i, v := range lists {
				if k := m.Meta["append"][i]; compact {
					v = maps[k]
				}

				if m.Echo(v); i < len(lists)-1 {
					m.Echo(col)
				}
			}
			m.Echo(row)
			return true
		}
	}

	// 输出表头
	row := map[string]string{}
	wor := []string{}
	for _, k := range m.Meta["append"] {
		row[k], wor = k, append(wor, k+strings.Repeat(space, width[k]-kit.Width(k, len(space))))
	}
	if !cb(row, wor, -1) {
		return m
	}

	// 输出数据
	for i := 0; i < depth; i++ {
		row := map[string]string{}
		wor := []string{}
		for _, k := range m.Meta["append"] {
			data := ""
			if i < len(m.Meta[k]) {
				data = m.Meta[k][i]
			}

			row[k], wor = data, append(wor, data+strings.Repeat(space, width[k]-kit.Width(data, len(space))))
		}
		m.Log("fuck", "waht %v", row)
		if !cb(row, wor, i) {
			break
		}
	}

	return m
}
func (m *Message) Sort(key string, arg ...string) *Message {
	cmp := "str"
	if len(arg) > 0 {
		cmp = arg[0]
	}

	number := map[int]int{}
	table := []map[string]string{}
	m.Table(func(index int, line map[string]string) {
		table = append(table, line)
		switch cmp {
		case "int":
			number[index] = kit.Int(line[key])
		case "int_r":
			number[index] = -kit.Int(line[key])
		case "time":
			number[index] = kit.Time(line[key])
		case "time_r":
			number[index] = -kit.Time(line[key])
		}
	})

	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			result := false
			switch cmp {
			case "str":
				if table[i][key] > table[j][key] {
					result = true
				}
			case "str_r":
				if table[i][key] < table[j][key] {
					result = true
				}
			default:
				if number[i] > number[j] {
					result = true
				}
			}

			if result {
				table[i], table[j] = table[j], table[i]
				number[i], number[j] = number[j], number[i]
			}
		}
	}

	for _, k := range m.Meta["append"] {
		delete(m.Meta, k)
	}

	for _, v := range table {
		for _, k := range m.Meta["append"] {
			m.Add("append", k, v[k])
		}
	}
	return m
}
func (m *Message) Copy(msg *Message, arg ...string) *Message {
	if msg == nil || m == msg {
		return m
	}

	for i := 0; i < len(arg); i++ {
		meta := arg[i]

		switch meta {
		case "target":
			m.target = msg.target
		case "callback":
			m.callback = msg.callback
		case "detail", "result":
			if len(msg.Meta[meta]) > 0 {
				m.Add(meta, msg.Meta[meta][0], msg.Meta[meta][1:])
			}
		case "option", "append":
			if msg.Meta == nil {
				msg.Meta = map[string][]string{}
			}
			if msg.Meta[meta] == nil {
				break
			}
			if i == len(arg)-1 {
				arg = append(arg, msg.Meta[meta]...)
			}

			for i++; i < len(arg); i++ {
				if v, ok := msg.Data[arg[i]]; ok {
					m.Put(meta, arg[i], v)
				} else if v, ok := msg.Meta[arg[i]]; ok {
					m.Set(meta, arg[i], v) // TODO fuck Add
				}
			}
		default:
			if msg.Hand {
				meta = "append"
			} else {
				meta = "option"
			}

			if v, ok := msg.Data[arg[i]]; ok {
				m.Put(meta, arg[i], v)
			}
			if v, ok := msg.Meta[arg[i]]; ok {
				m.Add(meta, arg[i], v)
			}
		}
	}

	return m
}
func (m *Message) Echo(str string, arg ...interface{}) *Message {
	if len(arg) > 0 {
		return m.Add("result", fmt.Sprintf(str, arg...))
	}
	return m.Add("result", str)
}

func (m *Message) Cmdp(t time.Duration, head []string, prefix []string, suffix [][]string) *Message {
	if head != nil && len(head) > 0 {
		m.Show(strings.Join(head, " "), "...\n")
	}

	for i, v := range suffix {
		m.Show(fmt.Sprintf("%v/%v %v...\n", i+1, len(suffix), v))
		m.CopyFuck(m.Cmd(prefix, v), "append")
		time.Sleep(t)
	}
	m.Show("\n")
	m.Table()
	return m
}
func (m *Message) Cmdm(args ...interface{}) *Message {
	m.Log("info", "current: %v", m.Magic("session", "current"))

	arg := []string{}
	if pod := kit.Format(m.Magic("session", "current.pod")); pod != "" {
		arg = append(arg, "context", "ssh", "remote", pod)
	}
	if ctx := kit.Format(m.Magic("session", "current.ctx")); ctx != "" {
		arg = append(arg, "context", ctx)
	}
	arg = append(arg, kit.Trans(args...)...)

	m.Spawn().Cmd(arg).CopyTo(m)
	return m
}
func (m *Message) Cmdy(args ...interface{}) *Message {
	m.Cmd(args...).CopyTo(m)
	return m
}
func (m *Message) Cmdx(args ...interface{}) string {
	msg := m.Cmd(args...)
	if msg.Result(0) == "error: " {
		return msg.Result(1)
	}
	return msg.Result(0)
}
func (m *Message) Cmds(args ...interface{}) bool {
	return m.Cmd(args...).Results(0)
}
func (m *Message) Cmd(args ...interface{}) *Message {
	if m == nil {
		return m
	}

	if len(args) > 0 {
		m.Set("detail", kit.Trans(args...))
	}
	key, arg := m.Meta["detail"][0], m.Meta["detail"][1:]

	msg := m
	if strings.Contains(key, ":") {
		ps := strings.Split(key, ":")
		msg, key, arg = msg.Sess("ssh"), "_route", append([]string{"sync", ps[0], ps[1]}, arg...)
		defer func() { m.Copy(msg, "append").Copy(msg, "result") }()
		m.Hand = true

	} else if strings.Contains(key, ".") {
		arg := strings.Split(key, ".")
		msg, key = msg.Sess(arg[0]), arg[1]
		msg.Option("remote_code", "")
	}
	if msg == nil {
		return msg
	}

	msg = msg.Match(key, true, func(msg *Message, s *Context, c *Context, key string) bool {
		msg.Hand = false
		if x, ok := c.Commands[key]; ok && x.Hand != nil {
			msg.TryCatch(msg, true, func(msg *Message) {
				msg.Log("cmd", "%s %s %v %v", c.Name, key, arg, msg.Meta["option"])

				for _, form := range []map[string]int{map[string]int{"page.limit": 1, "page.offset": 1}, x.Form} {

					if args := []string{}; form != nil {
						for i := 0; i < len(arg); i++ {
							if n, ok := form[arg[i]]; ok {
								if n < 0 {
									n += len(arg) - i
								}
								for j := i + 1; j <= i+n && j < len(arg); j++ {
									if _, ok := form[arg[j]]; ok {
										n = j - i - 1
									}
								}
								if i+1+n > len(arg) {
									msg.Add("option", arg[i], arg[i+1:])
								} else {
									msg.Add("option", arg[i], arg[i+1:i+1+n])
								}
								i += n
							} else {
								args = append(args, arg[i])
							}
						}
						arg = args
					}
				}

				target := msg.target
				msg.target = s

				msg.Hand = true
				switch v := msg.Gdb("command", key, arg).(type) {
				case string:
					msg.Echo(v)
				case nil:
					if msg.Options("auto_cmd") {
						if x.Auto != nil {
							x.Auto(msg, c, key, arg...)
						}
					} else {
						x.Hand(msg, c, key, arg...)
					}
				}
				if msg.target == s {
					msg.target = target
				}
			})
		}
		return msg.Hand
	})

	if !msg.Hand {
		msg.Log("error", "cmd run error %s", msg.Format())
	}
	return msg
}

func (m *Message) Confm(key string, args ...interface{}) map[string]interface{} {
	random := ""
	var chain interface{}
	if len(args) > 0 {
		switch arg := args[0].(type) {
		case []interface{}:
			chain, args = arg, args[1:]
		case []string:
			chain, args = arg, args[1:]
		case string:
			switch arg {
			case "%", "*":
				random, args = arg, args[1:]
			default:
				chain, args = arg, args[1:]
			}
		}
	}

	var v interface{}
	if chain == nil {
		v = m.Confv(key)
	} else {
		v = m.Confv(key, chain)
	}

	table, _ := v.([]interface{})
	value, _ := v.(map[string]interface{})
	if len(args) == 0 {
		return value
	}

	switch fun := args[0].(type) {
	case func(int, string):
		for i, v := range table {
			fun(i, kit.Format(v))
		}
	case func(int, string) bool:
		for i, v := range table {
			if fun(i, kit.Format(v)) {
				break
			}
		}
	case func(string, string):
		for k, v := range value {
			fun(k, kit.Format(v))
		}
	case func(string, string) bool:
		for k, v := range value {
			if fun(k, kit.Format(v)) {
				break
			}
		}
	case func(map[string]interface{}):
		if len(value) == 0 {
			return nil
		}
		fun(value)
	case func(string, map[string]interface{}):
		switch random {
		case "%":
			n, i := rand.Intn(len(value)), 0
			for k, v := range value {
				if val, ok := v.(map[string]interface{}); i == n && ok {
					fun(k, val)
					break
				}
				i++
			}
		case "*":
			fallthrough
		default:
			for k, v := range value {
				if val, ok := v.(map[string]interface{}); ok {
					fun(k, val)
				}
			}
		}
	case func(string, int, map[string]interface{}):
		for k, v := range value {
			if val, ok := v.([]interface{}); ok {
				for i, v := range val {
					if val, ok := v.(map[string]interface{}); ok {
						fun(k, i, val)
					}
				}
			}
		}

	case func(string, map[string]interface{}) bool:
		for k, v := range value {
			if val, ok := v.(map[string]interface{}); ok {
				if fun(k, val) {
					break
				}
			}
		}
	case func(int, map[string]interface{}):
		for i := m.Optioni("page.begin"); i < len(table); i++ {
			if val, ok := table[i].(map[string]interface{}); ok {
				fun(i, val)
			}
		}
	}
	return value
}
func (m *Message) Confx(key string, args ...interface{}) string {
	value := kit.Select(m.Conf(key), m.Option(key))
	if len(args) == 0 {
		return value
	}

	switch arg := args[0].(type) {
	case []string:
		if len(args) > 1 {
			value = kit.Select(value, arg, args[1])
		} else {
			value = kit.Select(value, arg)
		}
		args = args[1:]
	case map[string]interface{}:
		value = kit.Select(value, kit.Format(arg[key]))
	case string:
		value = kit.Select(value, arg)
	case nil:
	default:
		value = kit.Select(value, args[0])
	}

	format := "%s"
	if args = args[1:]; len(args) > 0 {
		format, args = kit.Format(args[0]), args[1:]
	}
	arg := []interface{}{format, value}
	for _, v := range args {
		arg = append(arg, v)
	}

	return kit.Format(arg...)
}
func (m *Message) Confs(key string, arg ...interface{}) bool {
	return kit.Right(m.Confv(key, arg...))
}
func (m *Message) Confi(key string, arg ...interface{}) int {
	return kit.Int(m.Confv(key, arg...))
}
func (m *Message) Confv(key string, args ...interface{}) interface{} {
	if strings.Contains(key, ".") {
		target := m.target
		defer func() { m.target = target }()

		ps := strings.Split(key, ".")
		if msg := m.Sess(ps[0], false); msg != nil {
			m.target, key = msg.target, ps[1]
		}
	}

	var config *Config
	m.Match(key, false, func(m *Message, s *Context, c *Context, key string) bool {
		if x, ok := c.Configs[key]; ok {
			config = x
			return true
		}
		return false
	})

	if len(args) == 0 {
		if config == nil {
			return nil
		}
		return config.Value
	}

	if config == nil {
		config = &Config{}
		m.target.Configs[key] = config
	}

	switch config.Value.(type) {
	case string:
		config.Value = kit.Format(args...)
	case bool:
		config.Value = kit.Right(args...)
	case int:
		config.Value = kit.Int(args...)
	case nil:
		config.Value = args[0]
	default:
		for i := 0; i < len(args); i += 2 {
			if i < len(args)-1 {
				config.Value = kit.Chain(config.Value, args[i], args[i+1])
			} else {
				return kit.Chain(config.Value, args[i])
			}
		}
	}

	return config.Value
}
func (m *Message) Conf(key string, args ...interface{}) string {
	return kit.Format(m.Confv(key, args...))
}
func (m *Message) Caps(key string, arg ...interface{}) bool {
	if len(arg) > 0 {
		return kit.Right(m.Cap(key, arg...))
	}
	return kit.Right(m.Cap(key))
}
func (m *Message) Capi(key string, arg ...interface{}) int {
	n := kit.Int(m.Cap(key))
	if len(arg) > 0 {
		return kit.Int(m.Cap(key, n+kit.Int(arg...)))
	}
	return n
}
func (m *Message) Cap(key string, arg ...interface{}) string {
	var cache *Cache
	m.Match(key, false, func(m *Message, s *Context, c *Context, key string) bool {
		if x, ok := c.Caches[key]; ok {
			cache = x
			return true
		}
		return false
	})

	if len(arg) == 0 {
		if cache == nil {
			return ""
		}
		if cache.Hand != nil {
			return cache.Hand(m, cache)
		}
		return cache.Value
	}

	if cache == nil {
		cache = &Cache{}
		m.target.Caches[key] = cache
	}

	if cache.Hand != nil {
		cache.Value = cache.Hand(m, cache, kit.Format(arg...))
	} else {
		cache.Value = kit.Format(arg...)
	}
	return cache.Value
}