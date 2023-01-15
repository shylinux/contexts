Volcanos(chat.ONIMPORT, {_init: function(can, msg) {
	function appends(list) { can.core.Item(list, function(key, list) { can.core.Item(list, function(size, value) {
		can.page.Append(can, can._output, [{view: html.IMG, style: {
			"background": "url(/publish/icon/full.jpg)", "background-size": size*20+"px", height:size, width: size,
			"background-repeat": "no-repeat", "background-position-x": -value[0]+"px", "background-position-y": -value[1]+"px",
			border: "lightgray solid 1px", margin: "2px", float: "left",
		}}])
	}), can.page.Append(can, can._output, [{type: "br", style: {"clear": "both"}}]) }) }

	appends({
		close: {16: [82, 158], 18: [93, 177], 20: [103, 197], 24: [123, 236]},
		create: {16: [27, 158], 18: [30, 177], 20: [33, 197], 24: [40, 236]},
		refresh: {16: [194, 241], 18: [218, 271], 20: [242, 301], 24: [291, 362]},
		back: {16: [27, 185], 18: [30, 208], 20: [33, 232], 24: [40, 278]},
		"goto": {16: [82, 185], 18: [92, 208], 20: [102, 232], 24: [123, 278]},
	})
	can.page.Append(can, can._output, [{view: html.IMG, style: {"background": "url(/publish/icon/full.jpg)", "background-size": "480px", height: 1290, width: 480}}])
	can.sup.onimport.size(can.sup, can.ConfHeight()+html.ACTION_HEIGHT, can.sup.ConfWidth(480)), can.sup.onaction._output(can.sup, msg)
}})
