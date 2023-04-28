Volcanos(chat.ONIMPORT, {_init: function(can, msg, cb) {
	msg.Table(function(item) { can.page.Append(can, can._output, [{view: html.ITEM, list: [{view: html.ICON, list: [{img: "/require/"+item.icon}] }], onclick: function(event) {
		can.sup.onexport.record(can, item.name, mdb.NAME, item)
	} }]) })
	can._output.oncontextmenu = function(event) { var carte = can.user.carte(event, can, {
		create: function() { can.Update(event, [ctx.ACTION, mdb.CREATE]) },
	}); can.page.style(can, carte._target, html.LEFT, event.x) }
	return
	var current = null, before, begin
	can.page.SelectChild(can, can._output, "*", function(target) { target.draggable = true
		target.ondragstart = function() { current = target, can.page.style(can, target, "visibility", "hidden") }
		target.ondragenter = function(event) { before = target, begin = {x: event.x, y: event.y} }
		target.ondragover = function(event) { var offset = event.x - begin.x
			can.page.style(can, target, {position: "relative", left: -offset})
		}
		target.ondragleave = function(event) { }
		target.ondragend = function(event) { before && can.page.insertBefore(can, current, before)
			 can.page.SelectChild(can, can._output, "*", function(target) { can.page.style(can, target, {position: "", left: "", visibility: "visible"}) })
		}
	})
}})
