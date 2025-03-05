.PHONY: install build run

install:
	npm i

build:
	npm run build
	go build -o server main.go

run:
	./server --host 0.0.0.0