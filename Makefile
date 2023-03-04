export CGO_ENABLED=0

binarys = bin/ice.bin

all: def
	@date +"%Y-%m-%d %H:%M:%S"
	go build -v -o ${binarys} src/main.go src/version.go src/binpack.go && ./${binarys} forever restart &>/dev/null

app: def
	CGO_ENABLED=1 go build -v -o usr/publish/contexts.app/Contents/MacOS/contexts src/webview.go src/version.go src/binpack.go && ./${binarys} forever restart &>/dev/null
	# hdiutil create usr/publish/tmp.dmg -ov -volname "ContextsInstall" -fs HFS+ -srcfolder "usr/publish/contexts.app"
	# rm -f usr/publish/ContextsInstall.dmg
	# hdiutil convert usr/publish/tmp.dmg -format UDZO -o usr/publish/ContextsInstall.dmg

%: src/%.go def
	@date +"%Y-%m-%d %H:%M:%S"
	go build -v -o usr/publish/$@ src/$@.go src/version.go && chmod u+x usr/publish/$@

def:
	@ [ -f src/version.go ] || echo "package main" > src/version.go
	@ [ -f src/binpack.go ] || echo "package main" > src/binpack.go

