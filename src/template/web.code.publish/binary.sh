mkdir contexts; cd contexts
ctx_dev={{.Option "domain"}}{{.Option "ctx_env"}} temp=$(mktemp); if curl -h &>/dev/null; then curl -o $temp -fsSL $ctx_dev; else wget -O $temp -q $ctx_dev; fi; source $temp binary
