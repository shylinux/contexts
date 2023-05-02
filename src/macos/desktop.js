Volcanos(chat.ONIMPORT, {
	_init: function(can, msg, cb) { if (can.isCmdMode()) { can.onappend.style(can, html.OUTPUT), can.ConfHeight(can.page.height()) }
		can.ui = {}, can.base.isFunc(cb) && cb(msg), can.onmotion.clear(can), can.page.styleHeight(can, can._output, can.ConfHeight())
		can.onimport._menu(can), can.onimport._dock(can), can.onimport._desktop(can, msg), can.onlayout.background(can, can.user.info.background, can._fields)
	},
	_menu: function(can) { can.onappend.plugin(can, {index: "web.chat.macos.menu", style: html.OUTPUT}, function(sub) { can.ui.menu = sub
		sub.onexport.record = function(_, value, key, item) {
			switch (value) {
				case "system":
					var carte = can.user.carte(event, can, {}, can.core.Item(can.onfigure), function(event, button, meta, carte) {
						can.onfigure[button](event, can, carte)
					}); break
			}
		}
	}) },
	_dock: function(can) { can.onappend.plugin(can, {index: "web.chat.macos.dock", style: html.OUTPUT}, function(sub) { can.ui.dock = sub
		sub.onexport.output = function(sub, msg) { can.page.style(can, sub._target, html.LEFT, (can.ConfWidth()-msg.Length()*80)/2) }
		sub.onexport.record = function(sub, value, key, item) { can.onimport._window(can, item) }
	}) },
	_desktop: function(can, msg) { var target = can.page.Append(can, can._output, [{view: "desktop"}])._target
		msg = msg||can._msg, msg.Table(function(item) { can.page.Append(can, target, [{view: html.ITEM, list: [{view: html.ICON, list: [{img: can.misc.PathJoin(item.icon)}]}, {view: [mdb.NAME, "", item.name]}]}]) })
		can.onimport.tabs(can, [{name: "desktop"+(can.page.Select(can, can._output, "div.desktop").length-1)}], function() { can.onmotion.select(can, can._output, "div.desktop", target), can.ui.desktop = target }, function() { can.page.Remove(can, target) })._desktop = target
	},
	_window: function(can, item) {
		item.height = can.base.Min(can.ConfHeight()-400, 320, 800), item.width = can.base.Min(can.ConfWidth()-400, 480, 1000)
		can.onappend.plugin(can, item, function(sub) { can.ondetail.select(can, sub._target)
			var index = 0; can.core.Item({
				"#f95f57": function(event) { sub.onaction.close(event, sub) },
				"#fcbc2f": function(event) { sub.onmotion.hidden(sub, sub._target) },
				"#32c840": function(event) { sub.onaction.full(event, sub) },
			}, function(color, cb) { can.page.insertBefore(can, [{view: [[html.ITEM, html.BUTTON]], style: {"background-color": color, right: 10+20*index++}, onclick: cb}], sub._output) })
			sub.onimport.size(sub, item.height, item.width, true), can.onmotion.move(can, sub._target, {"z-index": 10, top: 125, left: 100})
			sub.onmotion.resize(can, sub._target, function(height, width) { sub.onimport.size(sub, height, width) }, 25)
			sub.onexport.actionHeight = function(sub) { return can.page.ClassList.has(can, sub._target, html.OUTPUT)? 0: html.ACTION_HEIGHT+20 },
			sub.onexport.record = function(sub, value, key, item) { can.onimport._window(can, item) }
		}, can.ui.desktop)
	},
	layout: function(can) { can.page.styleHeight(can, can._output, can.ConfHeight()) },
}, [""])
Volcanos(chat.ONACTION, {list: ["full"],
	create: function(event, can, button) { can.onimport._desktop(can) },
	full: function(event, can) { document.body.requestFullscreen() },
})
Volcanos(chat.ONDETAIL, {
	select: function(can, target) {
		can.page.SelectChild(can, can.ui.desktop, "fieldset", function(fieldset) {
			can.page.style(can, fieldset, "z-index", fieldset == target? "10": "9")
			fieldset == target && can.onmotion.toggle(can, fieldset, true)
		})
	},
})
Volcanos(chat.ONFIGURE, {
	"desktop\t>": function(event, can, carte) {
		can.user.carteRight(event, can, {}, [{view: [html.ITEM, "", mdb.CREATE], onclick: function(event) {
			can.onaction.create(event, can)
		}}].concat(can.page.Select(can, can._action, "div.tabs>span.name", function(target) {
			return {view: [html.ITEM, "", target.innerText+(can.page.ClassList.has(can, target.parentNode, html.SELECT)? " *": "")],
				onclick: function(event) { target.click() },
				oncontextmenu: function(event) { can.user.carteRight(event, can, {
					remove: function() { target.parentNode._close() },
				}, [], function() {}, carte) },
			}
		})), function(event) {}, carte)
	},
	"window\t>": function(event, can, carte) {
		can.user.carteRight(event, can, {}, can.page.Select(can, can.ui.desktop, "fieldset>legend", function(legend) {
			return {view: [html.ITEM, "", legend.innerText+(legend.parentNode.style["z-index"] == "10"? " *": "")], onclick: function(event) {
				can.ondetail.select(can, legend.parentNode)
			}}
		}), function(event) {}, carte)
	},
	"layout\t>": function(event, can, carte) { var list = can.page.SelectChild(can, can.ui.desktop, html.FIELDSET)
		can.user.carteRight(event, can, {
			grid: function(event) { for (var i = 0; i*i < list.length; i++) {} for (var j = 0; j*i < list.length; j++) {}
				var height = (can.ConfHeight()-25)/j, width = can.ConfWidth()/i; can.core.List(list, function(target, index) {
					can.page.style(can, target, html.TOP, parseInt(index/i)*height+25, html.LEFT, index%i*width)
					target._can.onimport.size(target._can, height, width)
				})
			},
			free: function(event) { can.core.List(list, function(target, index) {
				can.page.style(can, target, html.TOP, can.ConfHeight()/2/list.length*index+25, html.LEFT, can.ConfWidth()/2/list.length*index)
			}) },
		}, [], function(event) {}, carte)
	},
	full: function(event, can, carte) { can.onaction.full(event, can) },
})
