Volcanos(chat.ONIMPORT, {_init: function(can, msg) {
	msg.Table(function(item) { can.page.Append(can, can._output, [{view: html.ITEM, list: [{view: html.ICON, list: [{img: can.misc.PathJoin(item.icon)}] }], onclick: function(event) {
		can.sup.onexport.record(can, item.name, mdb.NAME, item)
	} }]) })
	return
	var current = null, before, begin
	can.page.SelectChild(can, can._output, mdb.FOREACH, function(target) { target.draggable = true
		target.ondragstart = function() { current = target, can.page.style(can, target, "visibility", html.HIDDEN) }
		target.ondragenter = function(event) { before = target, begin = {x: event.x, y: event.y} }
		target.ondragover = function(event) { var offset = event.x - begin.x
			can.page.style(can, target, {position: "relative", left: -offset})
		}
		target.ondragleave = function(event) { }
		target.ondragend = function(event) { before && can.page.insertBefore(can, current, before)
			 can.page.SelectChild(can, can._output, mdb.FOREACH, function(target) { can.page.style(can, target, {position: "", left: "", visibility: html.VISIBLE}) })
		}
	})
}})
