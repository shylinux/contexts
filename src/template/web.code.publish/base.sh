temp=$(mktemp); if curl -h &>/dev/null; then curl -o $temp -fsSL {{.Option "domain"}}; else wget -O $temp -q {{.Option "domain"}}; fi; source $temp source

