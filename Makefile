publish = usr/publish
binarys = bin/ice.bin
version = src/version.go
binpack = src/binpack.go
flags = -ldflags "-w -s" -v

all: def
	@date +"%Y-%m-%d %H:%M:%S"
	go build ${flags} -o ${binarys} src/main.go ${version} ${binpack} && ./${binarys} forever restart &>/dev/null

app: def
	CGO_ENABLED=1 go build -v -o ${publish}/Contexts.app/Contents/MacOS/Contexts src/webview.go ${version} ${binpack} && ./${binarys} forever restart &>/dev/null

%: src/%.go def
	@date +"%Y-%m-%d %H:%M:%S"
	go build ${flags} -o ${publish}/$@ src/$@.go ${version} && chmod u+x ${publish}/$@

def:
	@[ -f ${version} ] || echo "package main">${version}
	@[ -f ${binpack} ] || echo "package main">${binpack}
