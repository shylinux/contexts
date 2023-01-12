Volcanos(chat.ONIMPORT, {help: "导入数据", _init: function(can, msg) {
	can.page.Append(can, can._output, [{view: "img", style: {
		"background": "url(/publish/icon/full.jpg)", "background-size": "320px", width: 16, height: 16,
		"background-repeat": "no-repeat", "background-position-x": "-82px", "background-position-y": "-158px",
	}}])
	can.page.Append(can, can._output, [{view: "img", style: {
		"background": "url(/publish/icon/full.jpg)", "background-size": "360px", width: 18, height: 18,
		"background-repeat": "no-repeat", "background-position-x": "-93px", "background-position-y": "-177px",
	}}])
	can.page.Append(can, can._output, [{view: "img", style: {
		"background": "url(/publish/icon/full.jpg)", "background-size": "480px", width: 24, height:24,
		"background-repeat": "no-repeat", "background-position-x": "-123px", "background-position-y": "-236px",
	}}])
	can.page.Append(can, can._output, [{view: "img", style: {
		"background": "url(/publish/icon/full.jpg)", "background-size": "480px", width: 480, height:1290,
	}}])
	can.sup.onimport.size(can.sup, can.ConfHeight()+html.ACTION_HEIGHT, can.sup.ConfWidth(480)), can.sup.onaction._output(can.sup, msg)
}})
