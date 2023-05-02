Volcanos(chat.ONIMPORT, { _init: function(can, msg) { can.onmotion.clear(can), can.ui = can.onappend.layout(can), msg.Table(function(value, index) {
		var item = can.onimport.item(can, value, function(event) { if (can.onmotion.cache(can, function() { return value.name }, can.ui.content)) { return }
			can.runActionCommand(event, value.index, [], function(msg) {
				switch (value.name) {
					case "Applications": can.onimport.icons(can, msg, can.ui.content); break
					default: can.onappend.table(can, msg, null, can.ui.content)
				} can.onimport.layout(can)
			})
		}); index == 0 && item.click()
	}), can.onmotion.hidden(can, can.ui.profile), can.onmotion.hidden(can, can.ui.display) },
	icons: function(can, msg, target) { msg.Table(function(value) { value.icon = can.misc.PathJoin(value.icon||can.page.drawText(can, value.name, 80))
		can.page.Append(can, target, [{view: html.ITEM, list: [{view: html.ICON, list: [{img: value.icon}]}, {view: [mdb.NAME, "", value.name]}], onclick: function(event) {
			can.sup.onexport.record(can.sup, value.name, mdb.NAME, value)
		}}])
	}) },
	layout: function(can) { can.ui.layout(can.ConfHeight(), can.ConfWidth()) },
})
