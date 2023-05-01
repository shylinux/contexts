Volcanos(chat.ONIMPORT, {
	_init: function(can, msg, cb) { if (can.isCmdMode()) { can.onappend.style(can, html.OUTPUT), can.ConfHeight(can.page.height()) }
		can.ui = {}, can.base.isFunc(cb) && cb(msg), can.onmotion.clear(can), can.page.styleHeight(can, can._output, can.ConfHeight())
		can.onimport._menu(can), can.onimport._dock(can), can.onimport._desktop(can, msg), can.onlayout.background(can, can.user.info.background, can._fields)
	},
	_menu: function(can) { can.onappend.plugin(can, {index: "web.chat.macos.menu", style: html.OUTPUT}, function(sub) { can.ui.menu = sub
		sub.onexport.record = function(_, value, key, item) {
			switch (value) {
				case "system":
					var carte = can.user.carte(event, can, {
						"About This Mac": function(event) { can.user.toast(can, "about this mac") },
						"System Preferences...": function(event) { can.user.toast(can, "about this mac") },
						"App Store...": function(event) { can.user.toast(can, "about this mac") },
						"Rencent iterms >": function(event) { can.user.toast(can, "about this mac") },
						"Force Quit Chrome": function(event) { can.user.toast(can, "about this mac") },
						"Sleep": function(event) { can.user.toast(can, "about this mac") },
						"Restart...": function(event) { can.user.toast(can, "about this mac") },
						"Shut Down...": function(event) { can.user.toast(can, "about this mac") },
						"Lock Screen": function(event) { can.user.toast(can, "about this mac") },
						"Log Out shy...": function(event) { can.user.toast(can, "about this mac") },
						full: function(event) { can.onaction.full(event, can) },
						scale: function(event) { can.page.ClassList.neg(can, can.ui.desktop, "scale") },
					}); var list = can.page.Select(can, carte._target, html.DIV_ITEM)
					can.core.List([1, 3, 4, 5, 8], function(i) { can.page.insertBefore(can, [{type: html.HR}], list[i]) })
					break
			}
		}
	}) },
	_dock: function(can) { can.onappend.plugin(can, {index: "web.chat.macos.dock", style: html.OUTPUT}, function(sub) { can.ui.dock = sub
		sub.onexport.record = function(sub, value, key, item) { can.onimport._window(can, item) }
		sub.onexport.output = function(sub, msg) { can.page.style(can, sub._target, html.LEFT, (can.ConfWidth()-msg.Length()*80)/2) }
	}) },
	_desktop: function(can, msg) { var target = can.page.Append(can, can._output, [{view: "desktop"}])._target
		msg = msg||can._msg, msg.Table(function(item) { can.page.Append(can, target, [{view: html.ITEM, list: [{view: html.ICON, list: [{img: can.misc.PathJoin(item.icon)}]}, {view: [mdb.NAME, "", item.name]}]}]) })
		can.onimport.tabs(can, [{name: "desktop"}], function() { can.onmotion.select(can, can._output, "div.desktop", target), can.ui.desktop = target }, function() { can.page.Remove(can, target) })
	},
	_window: function(can, item) { item.height = can.ConfHeight()-(item.index == web.CODE_VIMER? 400: 400), item.width = can.base.Max(can.ConfWidth()-200, item.index == web.CODE_VIMER? 1000: 1000)
		can.onappend.plugin(can, item, function(sub) {
			var index = 0; can.core.Item({
				"#f95f57": function(event) { sub.onaction.close(event, sub) },
				"#fcbc2f": function(event) { sub.onmotion.hidden(sub, sub._target) },
				"#32c840": function(event) { sub.onaction.full(event, sub) },
			}, function(color, cb) { can.page.insertBefore(can, [{view: [[html.ITEM, html.BUTTON]], style: {"background-color": color, right: 10+20*index++}, onclick: cb}], sub._output) })
			sub.onimport.size(sub, item.height, item.width, true), can.onmotion.move(can, sub._target, {"z-index": 10, top: 25, left: 100})
			sub.onexport.record = function(sub, value, key, item) { can.onimport._window(can, item) }
		}, can.ui.desktop)
	},
	layout: function(can) { can.page.styleHeight(can, can._output, can.ConfHeight()) },
}, [""])
Volcanos(chat.ONACTION, {list: ["full"],
	create: function(event, can, button) { can.onimport._desktop(can) },
	full: function(event, can) { document.body.requestFullscreen() },
})
