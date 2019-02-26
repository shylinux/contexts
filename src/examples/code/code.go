package code

import (
	"contexts/ctx"
	"contexts/web"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

var Index = &ctx.Context{Name: "code", Help: "代码中心",
	Caches: map[string]*ctx.Cache{},
	Configs: map[string]*ctx.Config{
		"skip_login": &ctx.Config{Name: "skip_login", Value: map[string]interface{}{
			"/consul": "true",
		}, Help: "免登录"},
		"counter": &ctx.Config{Name: "counter", Value: map[string]interface{}{
			"nopen": "0", "nsave": "0",
		}, Help: "counter"},
		"counter_service": &ctx.Config{Name: "counter_service", Value: "http://localhost:9094/code/counter", Help: "counter"},
		"web_site": &ctx.Config{Name: "web_site", Value: []interface{}{
			map[string]interface{}{"_name": "MDN", "site": "https://developer.mozilla.org"},
			map[string]interface{}{"_name": "github", "site": "https://github.com"},
		}, Help: "web_site"},
		"componet_command": &ctx.Config{Name: "component_command", Value: "pwd", Help: "默认命令"},
		"componet_group":   &ctx.Config{Name: "component_group", Value: "index", Help: "默认组件"},
		"componet": &ctx.Config{Name: "componet", Value: map[string]interface{}{
			"login": []interface{}{
				map[string]interface{}{"name": "head", "template": "head"},
				map[string]interface{}{"name": "login", "help": "login", "template": "componet",
					"componet_ctx": "aaa", "componet_cmd": "auth", "arguments": []interface{}{"@sessid", "ship", "username", "@username", "password", "@password"},
					"inputs": []interface{}{
						map[string]interface{}{"type": "text", "name": "username", "label": "username", "value": ""},
						map[string]interface{}{"type": "password", "name": "password", "label": "password", "value": ""},
						map[string]interface{}{"type": "button", "value": "login"},
					},
					"display_append": "", "display_result": "",
				},
				map[string]interface{}{"name": "tail", "template": "tail"},
			},
			"index": []interface{}{
				map[string]interface{}{"name": "head", "template": "head"},
				map[string]interface{}{"name": "toolkit", "help": "toolkit", "template": "toolkit"},
				// map[string]interface{}{"name": "login", "help": "login", "template": "componet",
				// 	"componet_ctx": "aaa", "componet_cmd": "login", "arguments": []interface{}{"@username", "@password"},
				// 	"inputs": []interface{}{
				// 		map[string]interface{}{"type": "text", "name": "username", "label": "username"},
				// 		map[string]interface{}{"type": "password", "name": "password", "label": "password"},
				// 		map[string]interface{}{"type": "button", "value": "login"},
				// 	},
				// 	"display_append": "", "display_result": "",
				// },
				// map[string]interface{}{"name": "userinfo", "help": "userinfo", "template": "componet",
				// 	"componet_ctx": "aaa", "componet_cmd": "login", "arguments": []interface{}{"@sessid"},
				// 	"pre_run": true,
				// },
				map[string]interface{}{"name": "clipboard", "help": "clipboard", "template": "clipboard"},
				map[string]interface{}{"name": "buffer", "help": "buffer", "template": "componet",
					"componet_ctx": "cli", "componet_cmd": "tmux", "arguments": []interface{}{"buffer"}, "inputs": []interface{}{
						map[string]interface{}{"type": "text", "name": "limit", "label": "limit", "value": "3"},
						map[string]interface{}{"type": "text", "name": "index", "label": "index"},
						map[string]interface{}{"type": "button", "value": "refresh"},
					},
					"pre_run": true,
				},
				// map[string]interface{}{"name": "time", "help": "time", "template": "componet",
				// 	"componet_ctx": "cli", "componet_cmd": "time", "arguments": []interface{}{"@string"},
				// 	"inputs": []interface{}{
				// 		map[string]interface{}{"type": "text", "name": "time_format",
				// 			"label": "format", "value": "2006-01-02 15:04:05",
				// 		},
				// 		map[string]interface{}{"type": "text", "name": "string", "label": "string"},
				// 		map[string]interface{}{"type": "button", "value": "refresh"},
				// 	},
				// },
				// map[string]interface{}{"name": "json", "help": "json", "template": "componet",
				// 	"componet_ctx": "nfs", "componet_cmd": "json", "arguments": []interface{}{"@string"},
				// 	"inputs": []interface{}{
				// 		map[string]interface{}{"type": "text", "name": "string", "label": "string"},
				// 		map[string]interface{}{"type": "button", "value": "refresh"},
				// 	},
				// },
				map[string]interface{}{"name": "dir", "help": "dir", "template": "componet",
					"componet_ctx": "nfs", "componet_cmd": "dir", "arguments": []interface{}{"@dir", "dir_sort", "@sort_order", "@sort_field"},
					"pre_run": true, "display_result": "",
					"inputs": []interface{}{
						map[string]interface{}{"type": "choice", "name": "dir_type",
							"label": "dir_type", "value": "both", "choice": []interface{}{
								map[string]interface{}{"name": "both", "value": "both"},
								map[string]interface{}{"name": "file", "value": "file"},
								map[string]interface{}{"name": "dir", "value": "dir"},
							},
						},
						map[string]interface{}{"type": "choice", "name": "sort_field",
							"label": "sort_field", "value": "time", "choice": []interface{}{
								map[string]interface{}{"name": "filename", "value": "filename"},
								map[string]interface{}{"name": "is_dir", "value": "type"},
								map[string]interface{}{"name": "line", "value": "line"},
								map[string]interface{}{"name": "size", "value": "size"},
								map[string]interface{}{"name": "time", "value": "time"},
							},
						},
						map[string]interface{}{"type": "choice", "name": "sort_order",
							"label": "sort_order", "value": "time_r", "choice": []interface{}{
								map[string]interface{}{"name": "str", "value": "str"},
								map[string]interface{}{"name": "str_r", "value": "str_r"},
								map[string]interface{}{"name": "int", "value": "int"},
								map[string]interface{}{"name": "int_r", "value": "int_r"},
								map[string]interface{}{"name": "time", "value": "time"},
								map[string]interface{}{"name": "time_r", "value": "time_r"},
							},
						},
						map[string]interface{}{"type": "text", "name": "dir", "label": "dir"},
					},
				},
				map[string]interface{}{"name": "upload", "help": "upload", "template": "componet",
					"componet_ctx": "web", "componet_cmd": "upload", "form_type": "upload",
					"inputs": []interface{}{
						map[string]interface{}{"type": "file", "name": "upload"},
						map[string]interface{}{"type": "submit", "value": "submit"},
					},
					"display_result": "",
				},
				// map[string]interface{}{"name": "download", "help": "download", "template": "componet",
				// 	"componet_ctx": "cli.shy", "componet_cmd": "source", "arguments": []interface{}{"@cmds"},,
				// 	"display_result": "", "download_file": "",
				// 	"inputs": []interface{}{
				// 		map[string]interface{}{"type": "text", "name": "download_file", "value": "data_2006_0102_1504.txt", "class": "file_name"},
				// 		map[string]interface{}{"type": "text", "name": "cmds", "value": "",
				// 			"class": "file_cmd", "clipstack": "clistack",
				// 		},
				// 	},
				// },
				map[string]interface{}{"name": "cmd", "help": "cmd", "template": "componet",
					"componet_ctx": "cli.shy", "componet_cmd": "source", "arguments": []interface{}{"@cmd"},
					"inputs": []interface{}{
						map[string]interface{}{"type": "text", "name": "cmd", "value": "",
							"class": "cmd", "clipstack": "void",
						},
					},
				},
				map[string]interface{}{"name": "ctx", "help": "ctx", "template": "componet",
					"componet_ctx": "cli.shy", "componet_cmd": "context", "arguments": []interface{}{"@ctx", "list"},
					"display_result": "",
					"inputs": []interface{}{
						map[string]interface{}{"type": "text", "name": "ctx", "value": ""},
						map[string]interface{}{"type": "button", "value": "refresh"},
					},
				},
				// map[string]interface{}{"name": "ccc", "help": "ccc", "template": "componet",
				// 	"componet_ctx": "cli.shy", "componet_cmd": "context", "arguments": []interface{}{"@current_ctx", "@ccc"},
				// 	"display_result": "",
				// 	"inputs": []interface{}{
				// 		map[string]interface{}{"type": "choice", "name": "ccc",
				// 			"label": "ccc", "value": "cmd", "choice": []interface{}{
				// 				map[string]interface{}{"name": "cmd", "value": "cmd"},
				// 				map[string]interface{}{"name": "config", "value": "config"},
				// 				map[string]interface{}{"name": "cache", "value": "cache"},
				// 			},
				// 		},
				// 		map[string]interface{}{"type": "button", "value": "refresh"},
				// 	},
				// },
				// map[string]interface{}{"name": "cmd", "help": "cmd", "template": "componet",
				// 	"componet_ctx": "cli.shy", "componet_cmd": "context", "arguments": []interface{}{"@current_ctx", "cmd", "list"},
				// 	"pre_run": true, "display_result": "",
				// 	"inputs": []interface{}{
				// 		map[string]interface{}{"type": "button", "value": "refresh"},
				// 	},
				// },
				// map[string]interface{}{"name": "history", "help": "history", "template": "componet",
				// 	"componet_ctx": "cli", "componet_cmd": "config", "arguments": []interface{}{"source_list"},
				// 	"pre_run": true, "display_result": "",
				// 	"inputs": []interface{}{
				// 		map[string]interface{}{"type": "button", "value": "refresh"},
				// 	},
				// },
				// map[string]interface{}{"name": "develop", "help": "develop", "template": "componet",
				// 	"componet_ctx": "web.code", "componet_cmd": "config", "arguments": []interface{}{"counter"},
				// 	"inputs": []interface{}{
				// 		map[string]interface{}{"type": "button", "value": "refresh"},
				// 	},
				// 	"pre_run":        true,
				// 	"display_result": "",
				// },
				// map[string]interface{}{"name": "windows", "help": "windows", "template": "componet",
				// 	"componet_ctx": "cli", "componet_cmd": "windows",
				// 	"inputs": []interface{}{
				// 		map[string]interface{}{"type": "button", "value": "refresh"},
				// 	},
				// 	"pre_run":        true,
				// 	"display_result": "",
				// },
				map[string]interface{}{"name": "runtime", "help": "runtime", "template": "componet",
					"componet_ctx": "cli", "componet_cmd": "runtime",
					"inputs": []interface{}{
						map[string]interface{}{"type": "button", "value": "refresh"},
					},
					"pre_run":        true,
					"display_result": "",
				},
				// map[string]interface{}{"name": "sysinfo", "help": "sysinfo", "template": "componet",
				// 	"componet_ctx": "cli", "componet_cmd": "sysinfo",
				// 	"inputs": []interface{}{
				// 		map[string]interface{}{"type": "button", "value": "refresh"},
				// 	},
				// 	"pre_run":        true,
				// 	"display_result": "",
				// },
				map[string]interface{}{"name": "mp", "template": "mp"},
				map[string]interface{}{"name": "tail", "template": "tail"},
			},
		}, Help: "组件列表"},
		"upgrade": &ctx.Config{Name: "upgrade", Value: map[string]interface{}{
			"system": []interface{}{"exit_shy", "common_shy", "init_shy", "bench", "boot_sh"},
			"portal": []interface{}{"code_tmpl", "code_js", "context_js"},
			"file": map[string]interface{}{
				"node_sh":    "bin/node.sh",
				"boot_sh":    "bin/boot.sh",
				"bench":      "bin/bench.new",
				"init_shy":   "etc/init.shy",
				"common_shy": "etc/common.shy",
				"exit_shy":   "etc/exit.shy",

				"code_tmpl":  "usr/template/code/code.tmpl",
				"code_js":    "usr/librarys/code.js",
				"context_js": "usr/librarys/context.js",
			},
		}, Help: "日志地址"},
	},
	Commands: map[string]*ctx.Command{
		"update": &ctx.Command{Name: "update", Help: "更新代码", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			return
		}},
		"/upgrade/": &ctx.Command{Name: "/upgrade/", Help: "下载文件", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			p := m.Cmdx("nfs.path", key)
			if strings.HasSuffix(key, "/bench") {
				bench := m.Cmdx("nfs.path", key+"."+m.Option("GOOS")+"."+m.Option("GOARCH"))
				if _, e := os.Stat(bench); e == nil {
					p = bench
				}
			}

			m.Log("fuck", "what %v", p)
			if _, e = os.Stat(p); e != nil {
				list := strings.Split(key, "/")
				m.Log("fuck", "what %v", list)
				p = m.Cmdx("nfs.path", m.Conf("upgrade", []string{"file", list[len(list)-1]}))
			}
			m.Log("fuck", "what %v", p)

			m.Log("info", "upgrade %s %s", p, m.Cmdx("aaa.hash", "file", p))
			http.ServeFile(m.Optionv("response").(http.ResponseWriter), m.Optionv("request").(*http.Request), p)
			return
		}},
		"upgrade": &ctx.Command{Name: "upgrade system|portal|script", Help: "服务升级", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if len(arg) == 0 {
				m.Cmdy("ctx.config", "upgrade")
				return
			}

			if m.Confs("upgrade", arg[0]) {
				key, arg = arg[0], arg[1:]
				m.Confm("upgrade", key, func(index int, value string) {
					arg = append(arg, value)
				})
			}

			restart := false
			for _, link := range arg {
				if file := m.Conf("upgrade", []string{"file", link}); file != "" {
					dir := path.Dir(file)
					if _, e = os.Stat(dir); e != nil {
						e = os.Mkdir(dir, 0777)
						m.Assert(e)
					}
					if m.Cmd("web.get", "dev", fmt.Sprintf("code/upgrade/%s", link),
						"GOOS", m.Conf("runtime", "host.GOOS"), "GOARCH", m.Conf("runtime", "host.GOARCH"),
						"save", file); strings.HasPrefix(file, "bin/") {
						if m.Cmd("cli.system", "chmod", "u+x", file); link == "bench" {
							m.Cmd("cli.system", "mv", "bin/bench", fmt.Sprintf("bin/bench_%s", m.Time("20060102_150405")))
							m.Cmd("cli.system", "mv", "bin/bench.new", "bin/bench")
						}
					}
					restart = true
				} else {
					m.Cmdy("web.get", "dev", fmt.Sprintf("code/upgrade/script/%s", link), "save", fmt.Sprintf("usr/script/%s", link))
				}
			}

			if restart {
				m.Cmd("cli.exit", 1)
			}
			return
		}},

		"/counter": &ctx.Command{Name: "/counter", Help: "/counter", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if len(arg) > 0 {
				m.Option("name", arg[0])
			}
			if len(arg) > 1 {
				m.Option("count", arg[1])
			}

			count := m.Optioni("count")
			switch v := m.Confv("counter", m.Option("name")).(type) {
			case string:
				i, e := strconv.Atoi(v)
				m.Assert(e)
				count += i
			}
			m.Log("info", "%v: %v", m.Option("name"), m.Confv("counter", m.Option("name"), fmt.Sprintf("%d", count)))
			m.Echo("%d", count)
			return
		}},
		"counter": &ctx.Command{Name: "counter name count", Help: "counter", Hand: func(m *ctx.Message, c *ctx.Context, key string, arg ...string) (e error) {
			if len(arg) > 1 {
				m.Copy(m.Spawn().Cmd("get", m.Conf("counter_service"), "name", arg[0], "count", arg[1]), "result")
			}
			return
		}},
	},
}

func init() {
	code := &web.WEB{}
	code.Context = Index
	web.Index.Register(Index, code)
}
