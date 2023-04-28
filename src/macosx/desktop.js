Volcanos(chat.ONIMPORT, {
	_init: function(can, msg, cb) { can.ui = {}, cb(msg), can.onmotion.clear(can)
		can.page.styleHeight(can, can._output, can.ConfHeight())
		can.onimport._menu(can), can.onimport._dock(can), can.onimport._desktop(can, msg), can.onlayout.background(can, can.user.info.background, can._fields)
	},
	_desktop: function(can, msg) { var target = can.page.Append(can, can._output, [{view: "desktop"}])._target
		msg = msg||can._msg, msg.Table(function(item) { can.page.Append(can, target, [{view: html.ITEM, list: [{view: html.ICON, list: [{img: "/require/"+item.icon}]}, {view: [mdb.NAME, "", item.name]}]}]) })
		can.onimport.tabs(can, [{name: "desktop", text: ""}], function() { can.onmotion.select(can, can._output, "div.desktop", target), can.ui.desktop = target }, function() { can.page.Remove(can, target) })
	},
	_window: function(can, item) { item.height = can.ConfHeight()-282, item.width = can.ConfWidth()-100
		can.onappend.plugin(can, item, function(sub) { sub.onimport.size(sub, item.height, item.width, true), can.onmotion.move(can, sub._target, {"z-index": 10, top: 100, left: 100}) }, can.ui.desktop)
	},
	_notifications: function(can) { can.onappend.plugin(can, {index: "web.chat.macosx.notifications", style: html.OUTPUT}, function(sub) { can.ui.notifications = sub, can.onmotion.hidden(can, sub._target) }) },
	_controls: function(can) { can.onappend.plugin(can, {index: "web.chat.macosx.controls", style: html.OUTPUT}, function(sub) { can.ui.controls = sub, can.onmotion.hidden(can, sub._target) }) },
	_menu: function(can) { can.onappend.plugin(can, {index: "web.chat.macosx.menu", style: html.OUTPUT}, function(sub) { can.ui.menu = sub
		can.onimport._notifications(can), can.onimport._controls(can)
		sub.onexport.record = function(_, value, key, item) {
			switch (value) {
				case "notifications": can.onmotion.toggle(can, can.ui.notifications._target); break
				case "controls": can.onmotion.toggle(can, can.ui.controls._target); break
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
						scale: function(event) {
							can.page.ClassList.neg(can, can.ui.desktop, "scale")
						},
					}); var list = can.page.Select(can, carte._target, html.DIV_ITEM)
					can.core.List([1, 3, 4, 5, 8], function(i) { can.page.insertBefore(can, [{type: html.HR}], list[i]) })
					can.page.style(can, carte._target, html.TOP, "25px")
					break
			}
		}
	}) },
	_dock: function(can) { can.onappend.plugin(can, {index: "web.chat.macosx.dock", style: html.OUTPUT}, function(sub) { can.ui.dock = sub
		sub.onexport.record = function(sub, value, key, item) { can.onimport._window(can, item) }
	}) },
	layout: function(can) {
		can.page.styleHeight(can, can._output, can.ConfHeight())
		return
		if (can._fields.offsetHeight > 32) {
			can.ConfHeight(can._fields.offsetHeight-html.ACTION_HEIGHT)
		}
		can.ConfWidth(can._output.offsetWidth)
		can.ui && can.ui.dock && can.page.style(can, can.ui.dock._target, "bottom", "10px")
	},
}, [""])
Volcanos(chat.ONACTION, {list: ["full"],
	create: function(event, can, button) { can.onimport._desktop(can) },
	// full: function(event, can, button) { can._fields.requestFullscreen() },
	full: function(event, can) { document.body.requestFullscreen() },
})
