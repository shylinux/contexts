Volcanos(chat.ONIMPORT, {_init: function(can, msg) {
	can.onappend.layout(can, can._output, html.FLEX, [
		{view: ["hi", html.DIV, "project"]}, [
			[
				{view: ["hi", html.DIV, "content"]},
				{view: ["he", html.DIV, "profile"]},
			],
			{view: ["he", html.DIV, "display"]},
		],
	]).layout(can.ConfWidth(), can.ConfHeight())
	can.onmotion.clear(can, can._output)
	var list = [
		{name: "h1", list: [{view: ["h1", html.DIV, "h1111"]}]},
		{name: "h2", list: [{view: ["h1", html.DIV, "h222"]}]},
		{name: "h3", list: [{view: ["h1", html.DIV, "h3333"]}]},
	]
	can.onappend.layout(can, can._output, "tabs-bottom", list)
}})
