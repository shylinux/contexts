Volcanos(chat.ONIMPORT, {_init: function(can, msg) {
	function degToRad(n) { return n * 2 * Math.PI / 360 }
	function drawFill(ctx, cb) { ctx.beginPath(), cb(), ctx.fill() }
	var ctx = can.page.Append(can, can._output, [{type: html.CANVAS, height: can.ConfHeight(), width: can.ConfWidth()}])._target.getContext("2d")
	drawFill(ctx, function() {
		ctx.strokeStyle = "red"
		ctx.fillStyle = "green"
		ctx.lineWidth = 2
		ctx.fillRect(0, 0, 200, 200)
		ctx.moveTo(200, 200), ctx.arc(200, 200, 100, degToRad(-45), degToRad(90), false), ctx.lineTo(200, 200)
		can.misc.Debug("what", 123)
		
		ctx.strokeStyle = "white"
		ctx.lineWidth = 2
		ctx.font = "36px arial"
		ctx.strokeText("Canvas text", 50, 50)
		
		ctx.fillStyle = "red"
		ctx.font = "48px georgia"
		ctx.fillText("Canvas text", 50, 150)
	})
}})
