TARGET=itbt

build:
	@go build -o $(TARGET)

clean:
	@rm itbt

test:
	@go test ./... -v

.PHONY: build clean test
