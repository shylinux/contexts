Volcanos(chat.ONIMPORT, {_init: function(can, msg) {
	can.onappend.layout(can, can._output, "flex", [
		{view: ["hi", html.DIV, "he"]},
		{view: ["he", html.DIV, "he"]},
		[
			{view: ["hi", html.DIV, "he"]},
			{view: ["he", html.DIV, "he"]},
		],
	]).layout(can.ConfWidth(), can.ConfHeight())
}})
