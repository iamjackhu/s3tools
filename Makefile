TARGET=./build
BINARY_NAME=bucketPerf
VERSION=$(version)

.PNONY: clean build

clean:
	rm -rf $(TARGET)/$(BINARY_NAME)

build: clean
	go mod tidy
	go build -ldflags "-X main._version='${VERSION}'" -o $(TARGET)/$(BINARY_NAME) *.go

run: build
	$(TARGET)/$(BINARY_NAME)