export GOPROXY=https://goproxy.cn
export GOPRIVATE=github.com
export CGO_ENABLED=0

all:
	@echo && date
	[ -f src/version.go ] || echo "package main" > src/version.go
	go build -v -o bin/ice.bin src/main.go src/version.go && chmod u+x bin/ice.bin && ./bin/ice.sh restart

%: src/%.go
	@echo && date
	go build -v -o usr/publish/$@ src/$@.go && chmod u+x usr/publish/$@
