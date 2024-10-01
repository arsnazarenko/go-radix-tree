TARGET_DIR = build
TARGET = $(TARGET_DIR)/main
BENCH_DIR = $(shell pwd)/radix
all: run

$(TARGET): ./cmd/main.go
	go build -o $@ $^ 

build: $(TARGET)

run: build
	./$(TARGET)

test:
	go test -v ./...

bench:
	go test -bench=. $(BENCH_DIR) -benchmem -benchtime=100000x
clean:
	@rm -rf $(TARGET_DIR) 

.PHONY: all build clean test run bench
