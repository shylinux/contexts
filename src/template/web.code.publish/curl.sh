temp=$(mktemp); curl -o $temp -fsSL {{.Option "domain"}}; source $temp binary
