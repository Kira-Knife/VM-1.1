.PHONY: build run test

build:
 docker build -t vmalertservice-db .

run:
 docker run --name vmalertservice-db -p 5432:5432 -d vmalertservice-db

test:
 go test ./...