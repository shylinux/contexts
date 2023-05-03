Volcanos(chat.ONIMPORT, {_init: function(can, msg) { can.page.Append(can, can._output, msg.Table(function(item) {
	return {view: html.ITEM, _init: function(target) {
		target.onclick = function(event) { can.sup.onexport.record(can.sup, item.index, ctx.INDEX, item)
			can.runAction(event, mdb.REMOVE, [item.hash], function() { can.page.Remove(can, target) })
		}
		var ui = can.onappend.layout(can, [html.ICON, [[wiki.TITLE, mdb.TIME], wiki.CONTENT]], "", target)
		can.page.Append(can, ui.icon, [{img: can.misc.PathJoin(item.icon||can.page.drawText(can, item.index, 60))}])
		ui.time.innerHTML = item.time.split(lex.SP).pop().split(nfs.DF).slice(0, 2).join(nfs.DF)
		ui.title.innerHTML = item.name||"", ui.content.innerHTML = item.text||""
	}}
})) }})