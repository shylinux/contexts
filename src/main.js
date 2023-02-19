Volcanos({river: {
	service: {name: "运营群", storm: {
		wx: {name: "公众号 wx", list: [
			{name: "context", help: "微信公众号", index: "web.wiki.word", args: ["usr/icebergs/misc/wx/wx.shy"]},
		]},
		mp: {name: "小程序 mp", list: [
			{name: "context", help: "微信小程序", index: "web.wiki.word", args: ["usr/icebergs/misc/mp/mp.shy"]},
		]},
		lark: {name: "机器人 lark", list: [
			{name: "context", help: "飞书机器人", index: "web.wiki.word", args: ["usr/icebergs/misc/lark/lark.shy"]},
		]},
	}},
	product: {name: "产品群", storm: {
		office: {name: "办公 office",  list: [
			{name: "feel", help: "影音媒体", index: "web.wiki.feel"},
			{name: "draw", help: "思维导图", index: "web.wiki.draw"},
			{name: "data", help: "数据表格", index: "web.wiki.data"},
			{name: "location", help: "地图导航", index: "web.chat.location"},
			{name: "context", help: "编程", index: "web.wiki.word", args: ["src/main.shy"]},
		]},
		website: {name: "定制 website", index: [
			"web.chat.website",
			"web.chat.div",
			"web.code.vimer",
			"web.dream",
		]},
	}},
	project: {name: "研发群", storm: {
		studio: {name: "研发 studio", list: [
			{name: "vimer", help: "编辑器", index: "web.code.vimer"},
			{name: "repos", help: "代码库", index: "web.code.git.status"},
			{name: "favor", help: "收藏夹", index: "web.chat.favor"},
			{name: "plan", help: "任务表", index: "web.team.plan"},
			{name: "ctx", help: "上下文", index: "web.wiki.word"},
		]},
		linux: {name:"系统 linux", list: [
			{name: "ctx", help: "平台", index: "web.wiki.word", args:["usr/linux-story/idc/idc.shy"]},
			{name: "ctx", help: "镜像", index: "web.wiki.word", args:["usr/linux-story/iso/iso.shy"]},
			{name: "ctx", help: "系统", index: "web.wiki.word", args:["usr/linux-story/src/main.shy"]},
		]},
		nginx: {name: "网关 nginx", list: [
			{name: "ctx", help: "代理", index: "web.wiki.word", args:["usr/nginx-story/src/main.shy"]},
		]},
		context: {name: "编程 context", list: [
			{name: "ctx", help: "编程", index: "web.wiki.word", args:["usr/golang-story/src/main.shy"]},
		]},
		redis: {name: "缓存 redis", list: [
			{name: "ctx", help: "数据缓存", index: "web.wiki.word", args:["usr/redis-story/src/main.shy"]},
			{name: "ctx", help: "消息队列", index: "web.wiki.word", args:["usr/redis-story/src/kafka/kafka.shy"]},
			{name: "ctx", help: "消息队列", index: "web.wiki.word", args:["usr/redis-story/src/pulsar/pulsar.shy"]},
		]},
		mysql: {name: "存储 mysql", list: [
			{name: "ctx", help: "数据存储", index: "web.wiki.word", args:["usr/mysql-story/src/main.shy"]},
			{name: "ctx", help: "搜索引擎", index: "web.wiki.word", args:["usr/mysql-story/src/elasticsearch/elasticsearch.shy"]},
			{name: "ctx", help: "搜索引擎", index: "web.wiki.word", args:["usr/mysql-story/src/clickhouse/clickhouse.shy"]},
		]},
	}},
	profile: {name: "测试群", storm: {
		release: {name: "发布 release", index: [
			"web.code.compile",
			"web.code.publish",
			"web.code.docker.client",
			"web.space",
			"web.dream",
			"web.code.git.server",
			"web.code.git.status",
		]},
		toolkit: {name: "工具 toolkit", index: [
			"web.code.favor",
			"web.code.xterm",
			"web.code.inner",
			"web.code.vimer",
			"web.code.bench",
			"web.code.pprof",
			"web.code.oauth",
		]},
		language: {name: "语言 language", index: [
			"web.code.c",
			"web.code.sh",
			"web.code.py",
			"web.code.shy",
			"web.code.js",
			"web.code.go",
		]},
	}},
	operate: {name: "运维群", storm: {
		aaa: {name: "权限 aaa", index: ["offer", "email", "user", "totp", "sess", "role"]},
		web: {name: "应用 web", index: ["broad", "serve", "space", "dream", "share", "cache", "spide"]},
		cli: {name: "系统 cli", index: ["qrcode", "daemon", "system", "runtime", "mirrors", "forever", "host", "port"]},
		nfs: {name: "文件 nfs", index: ["dir", "cat", "pack", "tail", "trash"]},
	}},
}})
