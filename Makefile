.PHONY: run

run:
	go build cmd/app/main.go && ./main

build: 
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

start: build
	docker-compose up --remove-orphans

rebuild: build
	docker-compose up --remove-orphans --build

down:
	docker-compose down --remove-orphans