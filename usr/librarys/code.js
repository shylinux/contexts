var page = Page({
    initFlashText: function(page, field, option, append, result) {
        option.onactions = function(msg) {
            page.reload()
        }
    },
    initFlashList: function(page, field, option, append, result) {
        option.dataset.flash_index = ctx.Search("flash_index")
        option.onactions = function(msg) {
            page.showFlashList(msg, option, append, result)
        }
        ctx.Run(page, option.dataset, [], option.onactions)
    },
    showFlashList: function(msg, option, append, result) {
        var page = this
        result.innerHTML = ""
        ctx.Table(msg, function(tip) {
            var ui = kit.AppendChild(result, [{"list": [
                {"view": ["detail", "div", tip.text]},
                {"code": [tip.output, "result", "result"]},
                {"view": ["action"], "list": [
                    {"button": ["查看详情", function(event) {
                        ctx.Search("flash_index", tip.index)
                    }]},
                    {"button": ["执行代码", function(event) {
                        tip.code && ctx.Run(page, option.dataset, [tip.index, "run"], function(msg) {
                            ui.output.innerText = msg.result
                        })
                    }]},
                    {"button": ["清空结果", function(event) {
                        ui.output.innerText = ""
                    }]},
                ]},
                {"code": [tip.output, "output", "output"]},
            ]}])
        })
    },

    initDirList: function(page, field, option, append, result) {
        var history = []
        function change(value) {
            option["current.dir"].value = value
            ctx.Runs(page, option)
            ctx.Current("dir", value)
            return value
        }
        function brow(value, dir, event) {
            option["current.dir"].value = value
            ctx.Runs(page, option)
            option["current.dir"].value = dir
            ctx.Runs(page, option)
        }
        return {
            "button": ["root", "back"], "action": function(value) {
                switch (value) {
                    case "back": history.length > -1 && change(history.pop() || "/"); break
                    case "root": change("/"); break
                }
            },
            "table": {"filename": function(value, key, row, index, event) {
                var dir = option["current.dir"].value
                var file = dir + ((dir && !dir.endsWith("/"))? "/": "") + value
                file.endsWith("/")? history.push(change(file)): brow(file, dir, event)
            }},
        }
    },
    initPodList: function(page, field, option, append, result) {
        function change(value) {
            option.pod.value = value
            ctx.Runs(page, option)
            ctx.Current("pod", value)
        }
        return {"table": {"key": change}}
    },
    initCtxList: function(page, field, option, append, result) {
        function change(value) {
            option.ctx.value = value
            ctx.Runs(page, option)
            ctx.Current("ctx", value)
        }
        return {"button": ["ctx", "shy", "web", "mdb"], "action": change, "table": {"names": change}}
    },

    initKitList: function(page, field, option, append, result) {
        var ui = kit.AppendChild(field, [{"type": "ul", "list": [
            {"fork": ["粘贴板", [
                {"leaf": ["+ 保存粘贴板(Ctrl+S)", function(event) {
                    console.log("save_txt")
                }]},
                {"leaf": ["+ 添加粘贴板(Ctrl+Y)", function(event) {
                    console.log("create_txt")
                }]},
                {"leaf": ["+ 快捷粘贴板(Ctrl+P)", function(event) {
                    console.log("quick_txt")
                }]},
            ]]},
            {"fork": ["命令行", [
                {"leaf": ["+ 折叠命令行(Ctrl+Z)", function(event, target) {
                    target.className = page.conf.show_result? "": "stick"
                    page.showResult(page)
                }]},
                {"leaf": ["+ 添加命令行(Ctrl+M)", function(event) {
                    page.addCmdList("cmd", page.conf.ncommand++)
                }]},
            ]]},
            {"fork": ["工作流", [
                {"leaf": ["+ 刷新工作流(Ctrl+R)", function(event) {
                    console.log("refresh_fly")
                }]},
                {"leaf": ["+ 添加工作流(Ctrl+T)", function(event) {
                    console.log("create_fly")
                }]},
                {"leaf": ["+ 命名工作流", function(event) {
                    console.log("rename_fly")
                }]},
                {"leaf": ["+ 删除工作流", function(event) {
                    console.log("remove_fly")
                }]},
            ]]},
        ]}])
        /*
        <li><div>命令行</div>
            <ul class="cmd">
                {{range $name, $cmd := conf . "toolkit"}}
                    <li>{{$name}} <input type="text" data-cmd="{{$name}}" onkeyup="onaction(event, 'toolkit')"><label class="result"></label></li>
                {{end}}

                <li class="stick" data-action="shrink_cmd">+ 折叠命令行(Ctrl+Z)</li>
                <li data-action="create_cmd">+ 添加命令行(Ctrl+M)</li>
                {{range $index, $cmd := index $bench_data "commands"}}
                    <li class="cmd{{$index}}" data-cmd="{{$index}}">{{index $cmd "now"|option}} {{$index}}: {{index $cmd "cmd"|option}}</li>
                {{end}}
            </ul>
        </li>
        <li><div>工作流</div>
            <ul class="fly">
                <li data-action="refresh_fly">+ 刷新工作流(Ctrl+R)</li>
                <li data-action="create_fly">+ 添加工作流(Ctrl+T)</li>
                {{range $key, $item := work .}}
                        <li data-key="{{$key}}">{{index $item "create_time"}} {{index $item "data" "name"}}({{slice $key 0 6}})</li>
                {{end}}
                <li data-action="rename_fly">+ 命名工作流</li>
                <li data-action="remove_fly">+ 删除工作流</li>
            </ul>
        </li>
        */
        return

        text = JSON.parse(bench_data.clipstack || "[]")
        for (var i = 0; i < text.length; i++) {
            copy_to_clipboard(text[i])
        }
        bench_data.board = bench_data.board || {}

        document.querySelectorAll("div.workflow").forEach(function(workflow) {
            // 移动面板
            workflow.style.left = context.Cookie("toolkit_left")
            workflow.style.top = context.Cookie("toolkit_top")
            var moving = false, left0 = 0, top0 = 0, x0 = 0, y0 = 0
            workflow.onclick = function(event) {
                if (event.target != workflow) {
                    return
                }
                moving = !moving
                if (moving) {
                    left0 = workflow.offsetLeft
                    top0 = workflow.offsetTop
                    x0 = event.clientX
                    y0 = event.clientY
                }
            }
            workflow.onmousemove = function(event) {
                if (moving) {
                    workflow.style.left = (left0+(event.clientX-x0))+"px"
                    workflow.style.top = (top0+(event.clientY-y0))+"px"
                    context.Cookie("toolkit_left", workflow.style.left)
                    context.Cookie("toolkit_top", workflow.style.top)
                }
            }

            // 固定面板
            if (context.Cookie("toolkit_class")) {
                workflow.className = context.Cookie("toolkit_class")
            }
            var head = workflow.querySelector("div")
            head.onclick = function(event) {
                head.dataset["show"] = !right(head.dataset["show"])
                workflow.className = right(head.dataset["show"])? "workflow max": "workflow"
                context.Cookie("toolkit_class", workflow.className)
            }

            // 折叠目录
            var toolkit = workflow.querySelector("ul.toolkit")
            toolkit.querySelectorAll("li>div").forEach(function(menu) {
                menu.onclick = function(event) {
                    menu.dataset["hide"] = !right(menu.dataset["hide"])
                    menu.nextElementSibling.style.display = right(menu.dataset["hide"])? "none": ""
                }
            })

            // 事件
            toolkit.querySelectorAll("li>ul>li").forEach(function(item) {
                // if (bench_data.board["key"] == item.dataset["key"]) {
                //     // item.className = "stick"
                // }

                item.onclick = function(event) {
                    var target = event.target
                    var data = item.dataset
                    switch (data["action"]) {
                        case "quick_txt":
                            code.quick_txt = !code.quick_txt
                            target.className= code.quick_txt? "stick": ""
                            break
                        case "copy_txt":
                            if (event.altKey) {
                                target.parentElement.removeChild(target)
                                return
                            }
                            if (event.shiftKey) {
                                var cmd = document.querySelector("form.option.cmd"+code.current_cmd+" input[name=cmd]")
                                cmd && (cmd.value += " "+text)
                                return
                            }
                            copy_to_clipboard(data["text"], true)
                            break
                        case "save_txt":
                            save_clipboard(item)
                            return
                        case "create_txt":
                            var text = prompt("text")
                            text && copy_to_clipboard(text)
                            return
                        case "refresh_fly":
                            location.reload()
                            return
                        case "create_fly":
                            context.Command(["sess", "bench", "create"], function(msg) {
                                context.Search("bench", msg.result[0])
                            })
                            return
                        case "rename_fly":
                            context.Command(["work", context.Search("bench"), "rename", prompt("name")], function() {
                                location.reload()
                            })
                            return
                        case "remove_fly":
                            var b = ""
                            document.querySelectorAll("div.workflow>ul.toolkit>li>ul.fly>li[data-key]").forEach(function(item){
                                if (!b && item.dataset["key"] != context.Search("bench")) {
                                    b = item.dataset["key"]
                                }
                            })
                            context.Search("bench", b)
                            context.Command(["work", context.Search("bench"), "delete"])
                            return
                    }

                    // 切换工作流
                    if (data["key"] && data["key"] != context.Search("bench")) {
                        context.Search("bench", data["key"])
                        return
                    }

                    // 切换命令行
                    var cmd = document.querySelector("form.option.cmd"+data["cmd"]+" input[name=cmd]")
                    cmd && cmd.focus()
                }
            })
        })
        return
    },
    initCmdList: function(page, field, option, append, result) {
        option.dataset["componet_name_alias"] = "cmd"
        option.dataset["componet_name_order"] = 0
        option.onactions = function(msg) {
            page.showCmdList(msg, option, append, result)
        }

        var cmd = option.querySelector("input[name=cmd]")
        cmd.onkeyup = function(event) {
            page.onCmdList(event, cmd, "input", option, append, result)
        }

        var action = conf.bench_data.action
        if (action && action["cmd"]) {
            cmd.value = action["cmd"].cmd[1]
            ctx.Runs(page, option)
            page.History.add("cmd", cmd.value)
        }

        var max = 0
        for (var k in action) {
            var order = parseInt(action[k].order)
            if (order > max) {
                max = order
            }
        }

        for (var i = 1; i <= max; i++) {
            var ui = page.addCmdList("cmd", i)
            if (action["cmd"+i]) {
                ui.cmd.value = action["cmd"+i].cmd[1]
                ctx.Runs(page, ui.option)
                page.History.add("cmd", ui.cmd.value)
            }
        }
        page.conf.ncommand = i
        return
    },
    showCmdList: function(msg, option, append, result) {
        append.innerHTML = ""
        msg && msg.append && kit.AppendTable(append, ctx.Table(msg), msg.append)
        result.innerText = (msg && msg.result)? msg.result.join(""): ""
        return
    },
    getCmdList: function(input, step, cmd) {
        var history = kit.History.get("cmd")
        var length = history.length
        var last = (parseInt(input.dataset["history_last"]||length)+step+length)%length
        if (0 <= last && last < length) {
            input.dataset["history_last"] = last
            cmd = history[last].data
        }
        return cmd
    },
    addCmdList: function(name, order) {
        var page = this
        var alias = name+order
        var ui = kit.AppendChild(document.querySelector("body"), [{"type": "fieldset", "data": {}, "list": [
            {"text": [alias, "legend"]},
            {"view": ["option "+alias, "form", "", "option"], "data": {"dataset": {
                "componet_group": "index", "componet_name": name, "componet_name_alias": alias, "componet_name_order": order,
            }, "onactions": function(msg) {
                page.showCmdList(msg, ui.option, ui.append, ui.result)
            }}, "list": [
                {"type": "input", "data": {"style": {"display": "none"}}},
                {"name": "cmd", "type": "input", "data": {"name": "cmd", "className": "cmd", "onkeyup": function(event) {
                    page.onCmdList(event, ui.cmd, "input", ui.option, ui.append, ui.result)
                }}},
            ]},
            {"view": ["append "+alias, "table", "", "append"]},
            {"code": ["", "result", "result "+alias]},
        ]}])
        page.OrderTable(ui.append)
        page.OrderCode(ui.code)
        ui.cmd.focus()
        return ui
    },
    delCmdList: function(name, order) {
        var option = document.querySelector("form.option.cmd"+order)
        option && document.body.removeChild(option.parentElement)

        for (;order < page.conf.ncommand; order++) {
            var input = document.querySelector("form.option.cmd"+order+" input[name=cmd]")
            if (input) {
                input.focus()
                return
            }
        }
        for (;order >= 0; order--) {
            var input = document.querySelector("form.option.cmd"+(order? order: "")+" input[name=cmd]")
            page.conf.ncommand = order+1
            if (input) {
                input.focus()
                return
            }
        }
    },
    onCmdList: function(event, target, action, option, append, result) {
        var page = this
        var order = option.dataset.componet_name_order
        var prev_order = (parseInt(order)-1+page.conf.ncommand)%page.conf.ncommand||""
        var next_order = (parseInt(order)+1)%page.conf.ncommand||""

        switch (action) {
            case "input":
                if (event.key == "Escape") {
                    target.blur()

                } else if (event.key == "Enter") {
                    target.dataset.history_last = page.History.get("cmd").length
                    page.History.add("cmd", target.value)
                    ctx.Runs(page, option)

                } else if (event.ctrlKey) {
                    switch (event.key) {
                        case "0":
                            var pre_pre = document.querySelector("code.result.cmd"+(event.shiftKey? next_order: prev_order)+" pre")
                            pre_pre && (target.value += pre_pre.innerText)
                            break
                        case "1":
                        case "2":
                        case "3":
                        case "4":
                        case "5":
                        case "6":
                        case "7":
                        case "8":
                        case "9":
                            if (code.quick_txt) {
                                var item = document.querySelectorAll("div.workflow>ul>li>ul.txt>li[data-text]")
                                target.value += item[parseInt(event.key)-1].dataset["text"]
                            } else {
                                var item = document.querySelectorAll("table.append.cmd"+(event.shiftKey? next_order: prev_order)+" td")
                                target.value += item[parseInt(event.key)-1].innerText
                            }
                            break
                        case "p":
                            target.value = page.getCmdList(target, -1, target.value)
                            break
                        case "n":
                            target.value = page.getCmdList(target, 1, target.value)
                            break
                        case "g":
                            var value = target.value.substr(0, target.selectionStart)
                            var last = parseInt(target.dataset.search_last || kit.History.get("cmd").length-1)
                            for (var i = last; i >= 0; i--) {
                                var cmd = kit.History.get("cmd", i).data
                                if (cmd.startsWith(value)) {
                                    target.value = cmd
                                    target.dataset.search_last = i-1
                                    target.setSelectionRange(value.length, cmd.length)
                                    break
                                }
                            }
                            target.dataset.search_last = ""
                            break
                        case "a":
                        case "e":
                        case "f":
                        case "b":
                        case "h":
                        case "d":
                            break
                        case "k":
                            kit.DelText(target, target.selectionStart)
                            break
                        case "u":
                            kit.DelText(target, 0, target.selectionEnd)
                            break
                        case "w":
                            var start = target.selectionStart-2
                            var end = target.selectionEnd-1
                            for (var i = start; i >= 0; i--) {
                                if (target.value[end] == " " && target.value[i] != " ") {
                                    break
                                }
                                if (target.value[end] != " " && target.value[i] == " ") {
                                    break
                                }
                            }
                            kit.DelText(target, i+1, end-i)
                            break
                        case "c":
                            append.innerHTML = ""
                            result.innerHTML = ""
                            break
                        case "r":
                            append.innerHTML = ""
                            result.innerHTML = ""
                        case "j":
                            target.dataset.history_last = page.History.get("cmd").length
                            page.History.add("cmd", target.value)
                            ctx.Runs(page, option)
                            break
                        case "l":
                            window.scrollTo(0, option.parentElement.offsetTop)
                            break
                        case "m":
                            page.addCmdList("cmd", page.conf.ncommand++)
                            break
                        case "i":
                            var input = document.querySelector("form.option.cmd"+next_order+" input[name=cmd]")
                            input && input.focus()
                            break
                        case "o":
                            var input = document.querySelector("form.option.cmd"+prev_order+" input[name=cmd]")
                            input && input.focus()
                            break
                        case "x":
                            result.style.height = result.style.height? "": page.conf.hide_height
                            break
                        case "z":
                            result.style.height = result.style.height? "": page.conf.show_height
                            break
                        case "q":
                            page.delCmdList("cmd", order)
                        default:
                            return
                    }
                } else {
                    kit.History.add("key", event.key)
                    if (kit.HitText(target, "jk")) {
                        kit.DelText(target, target.selectionStart-2, 2)
                        target.blur()
                    }
                }
                event.stopPropagation()
        }
    },

    onaction: function(event, target, action) {
        var page = this
        switch (action) {
            case "scroll":
                if (event.target != document.body) {
                    return
                }
                switch (event.key) {
                    case "h":
                        if (event.ctrlKey) {
                            window.scrollBy(-page.conf.scroll_x*10, 0)
                        } else {
                            window.scrollBy(-page.conf.scroll_x, 0)
                        }
                        break
                    case "H":
                        window.scrollBy(-body.scrollWidth, 0)
                        break
                    case "l":
                        if (event.ctrlKey) {
                            window.scrollBy(page.conf.scroll_x*10, 0)
                        } else {
                            window.scrollBy(page.conf.scroll_x, 0)
                        }
                        break
                    case "L":
                        window.scrollBy(body.scrollWidth, 0)
                        break
                    case "j":
                        if (event.ctrlKey) {
                            window.scrollBy(0, page.conf.scroll_y*10)
                        } else {
                            window.scrollBy(0, page.conf.scroll_y)
                        }
                        break
                    case "J":
                        window.scrollBy(0, body.scrollHeight)
                        break
                    case "k":
                        if (event.ctrlKey) {
                            window.scrollBy(0, -page.conf.scroll_y*10)
                        } else {
                            window.scrollBy(0, -page.conf.scroll_y)
                        }
                        break
                    case "K":
                        window.scrollBy(0, -body.scrollHeight)
                        break
                }
                return
            case "keymap":
                if (event.key == "Escape") {
                }
                if (event.key == "Enter") {
                }
                if (event.ctrlKey) {
                    switch (event.key) {
                        case "m":
                            page.addCmdList("cmd", page.conf.ncommand++)
                            return
                        case "z":
                            page.showResult(page)
                            return
                        case "0":
                        case "1":
                        case "2":
                        case "3":
                        case "4":
                        case "5":
                        case "6":
                        case "7":
                        case "8":
                        case "9":
                            document.querySelector("form.option.cmd"+(event.key||"")+" input[name=cmd]").focus()
                            return
                    }
                }
        }
    },
    showResult: function(page, type) {
        page.conf.show_result = !page.conf.show_result
        document.querySelectorAll("body>fieldset>code.result>pre").forEach(function(result) {
            result.style.height = (page.conf.show_result || result.innerText=="")? "": page.conf.show_height
        })
    },
    init: function(exp) {
        var page = this
        var body = document.body
        body.onkeydown = function(event) {
            page.onaction(event, body, "scroll")
        }
        body.onkeyup = function(event) {
            page.onaction(event, body, "keymap")
        }

        document.querySelectorAll("body>fieldset").forEach(function(field) {
            var option = field.querySelector("form.option")
            var append = field.querySelector("table.append")
            var result = field.querySelector("code.result pre")
            page.OrderForm(page, option, append, result)
            page.OrderTable(append)
            page.OrderCode(result)

            var init = page[field.dataset.init]
            if (typeof init == "function") {
                var conf = init(page, field, option, append, result)
                if (conf && conf["button"]) {
                    var buttons = []
                    conf.button.forEach(function(value, index) {
                        buttons.push({"button": [value, function(event) {
                            typeof conf["action"] == "function" && conf["action"](value, event)
                        }]})
                    })
                    kit.InsertChild(field, append, "div", buttons)
                }
                if (conf && conf["table"]) {
                    option.onactions = function(msg) {
                        append.innerHTML = ""
                        kit.AppendTable(append, ctx.Table(msg), msg.append, function(value, key, row, index, event) {
                            typeof conf["table"][key] == "function" && conf["table"][key](value, key, row, index, event)
                        })
                    }
                    ctx.Runs(page, option)
                }
            }
        })
    },
    conf: {
        scroll_x: 50,
        scroll_y: 50,

        ncommand: 1,
        show_result: true,
        show_height: "30px",
        hide_height: "14px",

        quick_txt: false,
    },
})
