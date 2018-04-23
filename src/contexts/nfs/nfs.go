package nfs // {{{
// }}}
import ( // {{{
	"contexts"

	"bufio"
	"encoding/json"
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/skip2/go-qrcode"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// }}}

type NFS struct {
	io io.ReadWriteCloser
	*bufio.Reader
	*bufio.Writer
	send   map[int]*ctx.Message
	target *ctx.Context

	in   *os.File
	out  *os.File
	buf  []string
	cli  *ctx.Message
	x, y int

	*ctx.Message
	*ctx.Context
}

func (nfs *NFS) print(str string, arg ...interface{}) bool { // {{{
	switch {
	case nfs.io != nil:
		str := fmt.Sprintf(str, arg...)
		nfs.y += strings.Count(str, "\n")
		fmt.Fprintf(nfs.in, "%s", str)
	case nfs.out != nil:
		str := fmt.Sprintf(str, arg...)
		nfs.y += strings.Count(str, "\n")
		fmt.Fprintf(nfs.out, "%s", str)
	default:
		return false
	}
	return true
}

// }}}
func (nfs *NFS) clear(line string) {
	termbox.SetCursor(nfs.x, nfs.y)
	termbox.Flush()
	nfs.print("                              ")

	termbox.SetCursor(nfs.x, nfs.y)
	termbox.Flush()

	nfs.print("%s%s", nfs.cli.Conf("PS1"), line)

}
func (nfs *NFS) Read(p []byte) (n int, err error) {
	if nfs.Cap("stream") != "stdio" {
		return nfs.in.Read(p)
	}

	his := len(nfs.buf)
	buf := make([]rune, 0, 1024)
	back := buf

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC:
				termbox.Close()
				os.Exit(1)
			case termbox.KeyCtrlJ, termbox.KeyEnter:
				buf = append(buf, '\n')
				nfs.print("\n")
				nfs.y++

				b := []byte(string(buf[:len(buf)]))
				n = len(b)
				nfs.Log("fuck", nil, "%d %v", n, b)
				copy(p, b)
				return
			case termbox.KeyCtrlP:
				his = (his + len(nfs.buf) - 1) % len(nfs.buf)
				b := nfs.buf[his]
				buf = buf[:len(b)]
				n := copy(buf, []rune(b))
				buf = buf[:n]

				nfs.clear(string(buf))
			case termbox.KeyCtrlN:
				his = (his + len(nfs.buf) - 1) % len(nfs.buf)
				b := nfs.buf[his]
				buf = buf[:len(b)]
				n := copy(buf, []rune(b))
				buf = buf[:n]

				nfs.clear(string(buf))
			case termbox.KeyCtrlH:
				if len(buf) > 0 {
					buf = buf[:len(buf)-1]
					nfs.clear(string(buf))
					n--
				}
			case termbox.KeyCtrlL:
				termbox.Sync()
				nfs.y = 0
				nfs.clear(string(buf))
			case termbox.KeyCtrlU:
				nfs.clear("")
				if len(buf) > 0 {
					back = buf
				}
				buf = make([]rune, 0, 1024)
				n = 0
			case termbox.KeyCtrlY:
				buf = back
				n = len(buf)
				nfs.clear(string(buf))
			case termbox.KeySpace:
				nfs.print(" ")
				buf = append(buf, ' ')
				n++
			default:
				print(string(ev.Ch))
				buf = append(buf, ev.Ch)
				n++
			}
		}
	}
	return

}

func (nfs *NFS) Spawn(m *ctx.Message, c *ctx.Context, arg ...string) ctx.Server { // {{{
	c.Caches = map[string]*ctx.Cache{
		"pos":    &ctx.Cache{Name: "读写位置", Value: "0", Help: "读写位置"},
		"nline":  &ctx.Cache{Name: "缓存命令行数", Value: "0", Help: "缓存命令行数"},
		"return": &ctx.Cache{Name: "缓存命令行数", Value: "0", Help: "缓存命令行数"},

		"nsend":  &ctx.Cache{Name: "消息发送数量", Value: "0", Help: "消息发送数量"},
		"nrecv":  &ctx.Cache{Name: "消息接收数量", Value: "0", Help: "消息接收数量"},
		"target": &ctx.Cache{Name: "消息接收模块", Value: "ssh", Help: "消息接收模块"},
		"result": &ctx.Cache{Name: "前一条指令执行结果", Value: "", Help: "前一条指令执行结果"},
		"sessid": &ctx.Cache{Name: "会话令牌", Value: "", Help: "会话令牌"},
	}
	c.Configs = map[string]*ctx.Config{}

	if len(arg) > 1 {
		if info, e := os.Stat(arg[1]); e == nil {
			c.Caches["name"] = &ctx.Cache{Name: "name", Value: info.Name(), Help: "文件名"}
			c.Caches["mode"] = &ctx.Cache{Name: "mode", Value: info.Mode().String(), Help: "文件权限"}
			c.Caches["size"] = &ctx.Cache{Name: "size", Value: fmt.Sprintf("%d", info.Size()), Help: "文件大小"}
			c.Caches["time"] = &ctx.Cache{Name: "time", Value: info.ModTime().Format("15:03:04"), Help: "创建时间"}
		}
	}

	s := new(NFS)
	s.Context = c
	return s

}

// }}}
func (nfs *NFS) Begin(m *ctx.Message, arg ...string) ctx.Server { // {{{
	nfs.Context.Master(nil)
	if nfs.Context == Index {
		Pulse = m
	}
	return nfs
}

// }}}
func (nfs *NFS) Start(m *ctx.Message, arg ...string) bool { // {{{
	nfs.Message = m
	if socket, ok := m.Data["io"]; ok {
		nfs.io = socket.(io.ReadWriteCloser)
		nfs.Reader = bufio.NewReader(nfs.io)
		nfs.Writer = bufio.NewWriter(nfs.io)
		nfs.send = make(map[int]*ctx.Message)
		nfs.target = m.Target()
		if target, ok := m.Data["target"]; ok {
			nfs.target = target.(*ctx.Context)
		}

		msg := m.Spawn(nfs.target)
		nfs.Caches["target"] = &ctx.Cache{Name: "target", Value: nfs.target.Name, Help: "文件名"}

		for {
			line, e := nfs.Reader.ReadString('\n')
			m.Assert(e)

			if line = strings.TrimSpace(line); len(line) > 0 {
				ls := strings.SplitN(line, ":", 2)
				ls[0] = strings.TrimSpace(ls[0])
				ls[1], e = url.QueryUnescape(strings.TrimSpace(ls[1]))
				m.Assert(e)

				switch ls[0] {
				case "detail":
					msg.Add("detail", ls[1])
				case "result":
					msg.Add("result", ls[1])
				default:
					msg.Add("option", ls[0], ls[1])
				}
				continue
			}

			if msg.Has("detail") {
				msg.Log("info", nil, "%d recv", m.Capi("nrecv", 1))
				msg.Log("info", nil, "detail: %v", msg.Meta["detail"])
				msg.Log("info", nil, "option: %v", msg.Meta["option"])
				msg.Options("stdio", false)

				func() {
					cmd := msg
					cmd.Call(func(sub *ctx.Message) *ctx.Message {
						for _, v := range sub.Meta["result"] {
							n, e := fmt.Fprintf(nfs.Writer, "result: %s\n", url.QueryEscape(v))
							sub.Assert(e)
							sub.Log("fuck", nil, "write %d", n)
						}

						sub.Append("nsend", sub.Option("nsend"))
						for _, k := range sub.Meta["append"] {
							for _, v := range sub.Meta[k] {
								n, e := fmt.Fprintf(nfs.Writer, "%s: %s\n", k, v)
								sub.Assert(e)
								sub.Log("fuck", nil, "write %d", n)
							}
						}

						sub.Log("info", nil, "%d send", sub.Optioni("nsend"))
						sub.Log("info", nil, "result: %v", sub.Meta["result"])
						sub.Log("info", nil, "append: %v", sub.Meta["append"])

						n, e := fmt.Fprintf(nfs.Writer, "\n")
						sub.Assert(e)
						sub.Log("sub", nil, "write %d", n)
						e = nfs.Writer.Flush()
						sub.Assert(e)

						if sub.Has("io") {
							if f, ok := sub.Data["io"].(io.ReadCloser); ok {
								io.Copy(nfs.Writer, f)
								nfs.Writer.Flush()
								f.Close()
							}
						}
						return sub
					})
				}()

			} else {
				msg.Meta["append"] = msg.Meta["option"]
				delete(msg.Meta, "option")

				msg.Log("info", nil, "%s send", msg.Meta["nsend"])
				msg.Log("info", nil, "result: %v", msg.Meta["result"])
				msg.Log("info", nil, "append: %v", msg.Meta["append"])

				send := nfs.send[msg.Appendi("nsend")]
				send.Copy(msg, "result")
				send.Copy(msg, "append")

				if send.Has("io") {
					if f, ok := send.Data["io"].(io.WriteCloser); ok {
						io.CopyN(f, nfs.Reader, int64(send.Appendi("size")))
						f.Close()
					}
				}

				send.Back(send)
			}

			msg = m.Spawn(nfs.target)
			m.Cap("target", nfs.target.Name)
		}
		return true
	}

	if in, ok := m.Data["in"]; ok {
		nfs.in = in.(*os.File)
	}
	if out, ok := m.Data["out"]; ok {
		nfs.out = out.(*os.File)
	}
	if len(arg) > 1 {
		if m.Cap("stream", arg[1]); arg[0] == "open" {
			return false
		}
	}

	cli := m.Reply()
	nfs.cli = cli
	yac := m.Find(cli.Conf("yac"))
	bio := bufio.NewScanner(nfs)
	if m.Cap("stream") == "stdio" {
		termbox.Init()
	}
	nfs.Context.Master(nil)
	pos := 0

	if buf, ok := m.Data["buf"]; ok {
		nfs.buf = buf.([]string)
		m.Capi("nline", len(nfs.buf))
		goto out
	}

	if len(arg) > 2 {
		nfs.print("%v\n", arg[2])
	}
	nfs.print("%s", cli.Conf("PS1"))

	for rest, text := "", ""; pos < m.Capi("nline") || bio.Scan(); {
		if pos == m.Capi("nline") {
			if text = bio.Text(); len(text) > 0 && text[len(text)-1] == '\\' {
				rest += text[:len(text)-1]
				continue
			}

			if text, rest = rest+text, ""; len(text) == 0 && len(nfs.buf) > 0 && nfs.in == os.Stdin {
				pos--
			} else {
				nfs.buf = append(nfs.buf, text)
				m.Capi("nline", 1)
			}
		}

		for ; pos < m.Capi("nline"); pos++ {

			for text = nfs.buf[pos] + "\n"; text != ""; {
				line := m.Spawn(yac.Target())
				line.Optioni("pos", pos)
				line.Options("stdio", true)
				line.Put("option", "cli", cli.Target())
				text = line.Cmd("parse", "line", "void", text).Get("result")
				cli.Target(line.Data["cli"].(*ctx.Context))
				if line.Has("return") {
					goto out
				}
				if line.Has("back") {
					pos = line.Appendi("back")
				}

				if result := strings.TrimRight(strings.Join(line.Meta["result"][1:len(line.Meta["result"])-1], ""), "\n"); len(result) > 0 {
					nfs.print("%s", result+"\n")
				}
			}
		}

		nfs.print("%s", cli.Conf("PS1"))
	}

out:
	if len(arg) > 1 {
		cli.Cmd("end")
	} else {
		m.Cap("status", "stop")
	}
	if m.Cap("stream") == "stdio" {
		termbox.Close()
	}
	return false
}

// }}}
func (nfs *NFS) Close(m *ctx.Message, arg ...string) bool { // {{{
	switch nfs.Context {
	case m.Target():
		if nfs.in != nil {
			m.Log("info", nil, "%d close %s", Pulse.Capi("nfile", -1)+1, m.Cap("name"))
			nfs.in.Close()
			nfs.in = nil
		}
	case m.Source():
	}
	return true
}

// }}}

var Pulse *ctx.Message
var Index = &ctx.Context{Name: "nfs", Help: "存储中心",
	Caches: map[string]*ctx.Cache{
		"nfile": &ctx.Cache{Name: "nfile", Value: "0", Help: "已经打开的文件数量"},
	},
	Configs: map[string]*ctx.Config{
		"size": &ctx.Config{Name: "size", Value: "1024", Help: "读取文件的默认大小值"},
	},
	Commands: map[string]*ctx.Command{
		"buffer": &ctx.Command{Name: "buffer [index string]", Help: "扫描文件, file: 文件名", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			if nfs, ok := m.Target().Server.(*NFS); m.Assert(ok) && nfs.buf != nil { // {{{
				for i, v := range nfs.buf {
					m.Echo("%d: %s\n", i, v)
				}
			} // }}}
		}},
		"copy": &ctx.Command{Name: "copy name [begin [end]]", Help: "复制文件, file: 文件名", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			if nfs, ok := m.Target().Server.(*NFS); m.Assert(ok) && len(nfs.buf) > 0 { // {{{
				begin, end := 0, len(nfs.buf)
				if len(arg) > 1 {
					i, e := strconv.Atoi(arg[1])
					m.Assert(e)
					begin = i
				}
				if len(arg) > 2 {
					i, e := strconv.Atoi(arg[2])
					m.Assert(e)
					end = i
				}
				m.Put("option", "buf", nfs.buf[begin:end])
				m.Start(arg[0], "扫描文件", key)
			} // }}}
		}},
		"scan": &ctx.Command{Name: "scan file", Help: "扫描文件, file: 文件名", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			if arg[0] == "stdio" { // {{{
				m.Put("option", "in", os.Stdin).Put("option", "out", os.Stdout).Start("stdio", "扫描文件", m.Meta["detail"]...)
			} else if f, e := os.Open(arg[0]); m.Assert(e) {
				m.Put("option", "in", f).Start(fmt.Sprintf("file%d", Pulse.Capi("nfile", 1)), "扫描文件", m.Meta["detail"]...)
			}
			// }}}
		}},
		"print": &ctx.Command{Name: "print str", Help: "扫描文件, file: 文件名", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			if nfs, ok := m.Target().Server.(*NFS); m.Assert(ok) && nfs.out != nil { // {{{
				fmt.Fprintf(nfs.out, "%s\n", arg[0])
			}
			// }}}
		}},

		"listen": &ctx.Command{Name: "listen args...", Help: "启动文件服务, args: 参考tcp模块, listen命令的参数", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			msg := m.Find("tcp") // {{{
			msg.Call(func(com *ctx.Message) *ctx.Message {
				sub := com.Spawn(m.Target())
				sub.Put("option", "target", m.Source())
				sub.Put("option", "io", com.Data["io"])
				sub.Cmd(com.Meta["detail"])
				sub.Cap("stream", com.Target().Name)
				return sub
			}, m.Meta["detail"])
			// }}}
		}},
		"accept": &ctx.Command{Name: "accept args...", Help: "连接文件服务, args: 参考tcp模块, dial命令的参数", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			m.Start(fmt.Sprintf("file%d", Pulse.Capi("nfile", 1)), "打开文件")
		}},
		"dial": &ctx.Command{Name: "dial args...", Help: "连接文件服务, args: 参考tcp模块, dial命令的参数", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			msg := m.Sess("com", "tcp") // {{{
			msg.CallBack(func(com *ctx.Message) *ctx.Message {
				sub := com.Spawn(m.Target())
				sub.Put("option", "target", m.Source())
				sub.Put("option", "io", com.Data["io"])
				sub.Start(fmt.Sprintf("file%d", Pulse.Capi("nfile", 1)), "打开文件")
				sub.Cap("stream", com.Target().Name)
				m.Target(sub.Target())
				return sub
			}, m.Meta["detail"])
			// }}}
		}},
		"send": &ctx.Command{Name: "send [file] args...", Help: "连接文件服务, args: 参考tcp模块, dial命令的参数", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			if nfs, ok := m.Target().Server.(*NFS); m.Assert(ok) { // {{{
				if m.Has("nrecv") {
					if len(arg) > 1 && arg[0] == "file" {
						info, e := os.Stat(arg[1])
						m.Assert(e)
						m.Append("name", info.Name())
						m.Append("size", info.Size())
						m.Append("time", info.ModTime())
						m.Append("mode", info.Mode())

						f, e := os.Open(arg[1])
						m.Assert(e)
						m.Put("append", "io", f)
					}

				} else {
					nfs.send[m.Optioni("nsend", m.Capi("nsend", 1))] = m

					if len(arg) > 1 && arg[0] == "file" {
						info, e := os.Stat(arg[1])
						m.Assert(e)
						m.Option("name", info.Name())
						m.Option("size", info.Size())
						m.Option("time", info.ModTime())
						m.Option("mode", info.Mode())

						fmt.Fprintf(nfs.Writer, "detail: recv\n")
					}
					for _, v := range arg {
						fmt.Fprintf(nfs.Writer, "detail: %v\n", v)
					}

					for _, k := range m.Meta["option"] {
						if k == "args" {
							continue
						}
						for _, v := range m.Meta[k] {
							fmt.Fprintf(nfs.Writer, "%s: %s\n", k, v)
						}
					}
					m.Log("info", nil, "%d send", m.Optioni("nsend"))
					m.Log("info", nil, "detail: %v", m.Meta["detail"])
					m.Log("info", nil, "option: %v", m.Meta["option"])

					fmt.Fprintf(nfs.Writer, "\n")
					nfs.Writer.Flush()

					if len(arg) > 1 && arg[0] == "file" {
						f, e := os.Open(arg[1])
						m.Assert(e)
						defer f.Close()
						_, e = io.Copy(nfs.Writer, f)
					}
				}
			} // }}}
		}},
		"recv": &ctx.Command{Name: "recv [file] args...", Help: "连接文件服务, args: 参考tcp模块, dial命令的参数", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			if nfs, ok := m.Target().Server.(*NFS); m.Assert(ok) { // {{{
				if m.Has("nrecv") {
					if len(arg) > 1 && arg[0] == "file" {
						f, e := os.Create(arg[1])
						m.Assert(e)
						defer f.Close()
						io.CopyN(f, nfs.Reader, int64(m.Optioni("size")))
					}

					return
				}

				nfs.send[m.Optioni("nrecv", m.Capi("nsend", 1))] = m

				if len(arg) > 1 && arg[0] == "file" {
					f, e := os.Create(arg[1])
					m.Assert(e)
					m.Put("option", "io", f)

					fmt.Fprintf(nfs.Writer, "detail: send\n")
				}

				for _, v := range arg {
					fmt.Fprintf(nfs.Writer, "detail: %v\n", v)
				}

				for _, k := range m.Meta["option"] {
					if k == "args" {
						continue
					}
					for _, v := range m.Meta[k] {
						fmt.Fprintf(nfs.Writer, "%s: %s\n", k, v)
					}
				}

				fmt.Fprintf(nfs.Writer, "\n")
				nfs.Writer.Flush()
			} // }}}
		}},
		"open": &ctx.Command{Name: "open file", Help: "打开文件, file: 文件名", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			if m.Has("io") { // {{{
				m.Put("option", "io", m.Data["io"])
				m.Start(fmt.Sprintf("file%d", Pulse.Capi("nfile", 1)), "打开文件", m.Meta["detail"]...)
				m.Echo(m.Target().Name)
			} else if f, e := os.OpenFile(arg[0], os.O_RDWR|os.O_CREATE, os.ModePerm); e == nil {
				m.Put("option", "in", f).Put("option", "out", f)
				m.Start(fmt.Sprintf("file%d", Pulse.Capi("nfile", 1)), "打开文件", m.Meta["detail"]...)
				m.Echo(m.Target().Name)
			} // }}}
		}},
		"read": &ctx.Command{Name: "read [size [pos]]", Help: "读取文件, size: 读取大小, pos: 读取位置", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			if nfs, ok := m.Target().Server.(*NFS); m.Assert(ok) && nfs.in != nil { // {{{
				var e error
				n := m.Confi("size")
				if len(arg) > 0 {
					n, e = strconv.Atoi(arg[0])
					m.Assert(e)
				}
				if len(arg) > 1 {
					m.Cap("pos", arg[1])
				}

				buf := make([]byte, n)
				if n, e = nfs.in.ReadAt(buf, int64(m.Capi("pos"))); e != io.EOF {
					m.Assert(e)
				}
				m.Echo(string(buf))

				if m.Capi("pos", n); m.Capi("pos") == m.Capi("size") {
					m.Cap("pos", "0")
				}
			} // }}}
		}},
		"write": &ctx.Command{Name: "write string [pos]", Help: "写入文件, string: 写入内容, pos: 写入位置", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			if nfs, ok := m.Target().Server.(*NFS); m.Assert(ok) && nfs.out != nil { // {{{

				if len(arg) > 1 {
					m.Cap("pos", arg[1])
				}

				if len(arg[0]) == 0 {
					m.Assert(nfs.out.Truncate(int64(m.Capi("pos"))))
					m.Cap("size", m.Cap("pos"))
					m.Cap("pos", "0")
				} else {
					n, e := nfs.out.WriteAt([]byte(arg[0]), int64(m.Capi("pos")))
					if m.Assert(e) && m.Capi("pos", n) > m.Capi("size") {
						m.Cap("size", m.Cap("pos"))
					}
					nfs.out.Sync()
				}

				m.Echo(m.Cap("pos"))
			} // }}}
		}},
		"load": &ctx.Command{Name: "load file [size]", Help: "写入文件, string: 写入内容, pos: 写入位置", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			if f, e := os.Open(arg[0]); m.Assert(e) { // {{{
				defer f.Close()

				m.Meta = nil
				size := 1024
				if len(arg) > 1 {
					if s, e := strconv.Atoi(arg[1]); m.Assert(e) {
						size = s
					}
				}
				buf := make([]byte, size)

				if l, e := f.Read(buf); m.Assert(e) {
					m.Log("info", nil, "read %d", l)
					m.Echo(string(buf[:l]))
				}
			} // }}}
		}},
		"save": &ctx.Command{Name: "save file string...", Help: "写入文件, string: 写入内容, pos: 写入位置", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			if f, e := os.Create(arg[0]); m.Assert(e) { // {{{
				defer f.Close()

				fmt.Fprint(f, strings.Join(arg[1:], ""))
			} // }}}
		}},
		"genqr": &ctx.Command{Name: "genqr [size] file string...", Help: "写入文件, string: 写入内容, pos: 写入位置", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			size := 256 // {{{
			if len(arg) > 2 {
				if s, e := strconv.Atoi(arg[0]); e == nil {
					arg = arg[1:]
					size = s
				}
			}
			qrcode.WriteFile(strings.Join(arg[1:], ""), qrcode.Medium, size, arg[0]) // }}}
		}},
		"json": &ctx.Command{Name: "json file", Help: "写入文件, string: 写入内容, pos: 写入位置", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			buf, e := json.Marshal(m.Data["data"]) // {{{
			m.Assert(e)

			f, e := os.Create(arg[0])
			m.Assert(e)
			f.Write(buf)
			f.Close()

			m.Echo(string(buf))
			// }}}
		}},

		"pwd": &ctx.Command{Name: "pwd", Help: "写入文件, string: 写入内容, pos: 写入位置", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) {
			wd, e := os.Getwd() // {{{
			m.Assert(e)
			m.Echo(wd) // }}}
		}},
	},
	Index: map[string]*ctx.Context{
		"void": &ctx.Context{Name: "void",
			Commands: map[string]*ctx.Command{
				"scan":  &ctx.Command{},
				"open":  &ctx.Command{},
				"save":  &ctx.Command{},
				"load":  &ctx.Command{},
				"genqr": &ctx.Command{},
			},
		},
	},
}

func init() {
	nfs := &NFS{}
	nfs.Context = Index
	ctx.Index.Register(Index, nfs)
}