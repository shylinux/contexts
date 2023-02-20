#! /bin/sh
export ctx_dev=http://192.168.10.8:9020 ctx_pod= ctx_mod=web.code.hi.hi
temp=$(mktemp); if curl -h &>/dev/null; then
	curl -o $temp -fsSL $ctx_dev
else
   	wget -O $temp -q $ctx_dev
fi && source $temp $ctx_mod
