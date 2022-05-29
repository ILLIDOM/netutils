EXECUTABLE=netutils
WINDOWS=$(EXECUTABLE)_windows.exe
LINUX=$(EXECUTABLE)_linux
DARWIN=$(EXECUTABLE)_darwin

windows: $(WINDOWS)
linux: $(LINUX)
darwin: $(DARWIN)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -o $(WINDOWS)

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -o $(LINUX)

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -o $(DARWIN)

build: windows linux darwin

all: build

test:
	go test ./...

clean:
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)

.PHONY: all build clean test