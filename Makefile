.PHONY: all clean build

all: clean build

build:
	GOOS=darwin GOARCH=386 go build -o s_darwin_386
	GOOS=linux GOARCH=386 go build -o s_linux_386
	GOOS=windows GOARCH=386 go build -o s_windows_386.exe

clean:
	rm -f s_darwin_386
	rm -f s_linux_386
	rm -f s_windows_386.exe

