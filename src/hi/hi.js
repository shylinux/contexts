Volcanos(chat.ONIMPORT, {_init: function(can, msg) {
	msg.Defer(function() { msg.Dump() })
	msg.Echo("hello world\n"+(can.Conf("content")||""))
	can.require(["he.js"], function() { }, function() { })
	// can.require(["/require/shylinux.com/x/contexts@v2.9.2/src/hi/he.js"], function() { }, function() { })
	// can.require(["http://localhost:9020/require/shylinux.com/x/contexts@v2.9.2/src/hi/he.js"], function() { }, function() { })
	// can.require(["https://shylinux.com/x/contexts@v2.9.2/src/hi/he.js?content=what"], function() { }, function() { })
}})
