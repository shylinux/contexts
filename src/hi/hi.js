
Volcanos(chat.ONIMPORT, {_init: function(can, msg) {
	msg.Defer(function() { msg.Dump() })
	can.misc.Info("what", can.base.MergeURL(location.href, "hi", 123, "debug", undefined))
	can.misc.Info("what", can.base.ParseURL(location.href))
	can.misc.Info("what", can.base.ParseJSON(location.href))
	can.misc.Info("what", can.base.ParseJSON('{"hi":1}'))
	can.misc.Info("what", can.base.ParseSize('1k'))
	can.misc.Info("what", can.base.ParseSize('1.2k'))
	can.misc.Info("what", can.base.Size('1224'))
	msg.Echo("hello world\n")
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
}})
