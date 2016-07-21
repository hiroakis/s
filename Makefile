.PHONY: all clean build

all: clean build

build:
	GOOS=darwin GOARCH=amd64 go build -o s_darwin_amd64
	GOOS=linux GOARCH=amd64 go build -o s_linux_amd64

clean:
	rm -f s_darwin_amd64
	rm -f s_linux_amd64

