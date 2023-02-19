Volcanos(chat.ONIMPORT, {_init: function(can, msg) {
	msg.Defer(function() { msg.Dump() })
	msg.Echo("hello world\n"+(can.Conf("content")||""))
}})