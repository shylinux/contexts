export ctx_dev=%s ctx_pod=%s ctx_mod=%s
temp=$(mktemp); if curl -V &>/dev/null; then curl -o $temp -fsSL $ctx_dev; else wget -O $temp -q $ctx_dev; fi && source $temp $ctx_mod
