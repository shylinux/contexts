publish = usr/publish
binarys = bin/ice.bin
version = src/version.go
binpack = src/binpack.go

all: def
	@date +"%Y-%m-%d %H:%M:%S"
	go build -v -o ${binarys} src/main.go ${version} ${binpack} && ./${binarys} forever restart &>/dev/null
	# go build -v -o ${binarys} src/main.go ${version} ${binpack}

app: def
	CGO_ENABLED=1 go build -v -o ${publish}/contexts.app/Contents/MacOS/contexts src/webview.go ${version} ${binpack} && ./${binarys} forever restart &>/dev/null
	# hdiutil create ${publish}/tmp.dmg -ov -volname "ContextsInstall" -fs HFS+ -srcfolder "${publish}/contexts.app"
	# rm -f ${publish}/ContextsInstall.dmg
	# hdiutil convert ${publish}/tmp.dmg -format UDZO -o ${publish}/ContextsInstall.dmg

%: src/%.go def
	@date +"%Y-%m-%d %H:%M:%S"
	go build -v -o ${publish}/$@ src/$@.go ${version} && chmod u+x ${publish}/$@

def:
	@[ -f ${version} ] || echo "package main">${version}
	@[ -f ${binpack} ] || echo "package main">${binpack}
