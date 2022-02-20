.phony: windows linux

noop:
	:

windows:
	env GOOS=windows GOARCH=amd64 go build -o renlfn_win.exe

linux:
	env GOOS=linux GOARCH=amd64 go build -o renlfn_linux.exe