.PHONY: build
build:
	go build -o server cmd/main.go

.PHONY: run
run:
	./server

.PHONY: clean
clean:
	rm ./server

.DEFAULT_GOAL = build