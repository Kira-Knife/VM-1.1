.PHONY: build run test

build:
	docker build -t vmalertservice-db .

run:
	docker run --name vmalertservice-db -p 5432:5432 -d vmalertservice-db

test:
	go test ./...

run_database:
	docker build -t pg_alertservice:2.0.0 .\app\alertservice\internal\database\Dockerfile
	docker run -d --name pg_alertservice -p 5432:5432 -v /opt/data/pg_alertservice:/var/lib/postgresql/data pg_alertservice:2.0.0
.PHONY: run_database
