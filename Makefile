run:
	go run ./cmd/timesplit/
debug: 
#TODO: add the debug build stuff so when you have access to a half decent debugger use it 
build-linux : 
	fyne build -os linux -o ./bin/linux/dev/timesplit ./cmd/timesplit/
build-windows: 
	CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 fyne build -os windows -o ./bin/windows/timesplit.exe ./cmd/timesplit/
package: 
# TODO: add the packaging for linux and windows and get an icon to be used with it 
	fyne package --os linux ./cmd/timesplit/
test: 

